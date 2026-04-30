package block

import (
	"fmt"
	"testing"
	_ "unsafe"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

//go:linkname finaliseBlockRegistryForRedstoneTorch github.com/df-mc/dragonfly/server/world.finaliseBlockRegistry
func finaliseBlockRegistryForRedstoneTorch()

func init() {
	finaliseBlockRegistryForRedstoneTorch()
}

func testRedstoneTorchTx(t *testing.T, f func(tx *world.Tx) error) {
	t.Helper()

	w := world.New()
	var txErr error
	<-w.Exec(func(tx *world.Tx) {
		txErr = f(tx)
	})
	if err := w.Close(); err != nil {
		t.Fatalf("close world: %v", err)
	}
	if txErr != nil {
		t.Fatal(txErr)
	}
}

func TestRedstoneTorchUnpowersReceiversBehindStronglyPoweredBlock(t *testing.T) {
	testRedstoneTorchTx(t, func(tx *world.Tx) error {
		torchPos := cube.Pos{0, 1, 0}
		stonePos := torchPos.Side(cube.FaceUp)
		notePos := stonePos.Side(cube.FaceEast)

		tx.SetBlock(torchPos, RedstoneTorch{Facing: cube.FaceDown, Lit: true}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(stonePos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

		updateTorchRedstone(torchPos, tx)
		if note := tx.Block(notePos).(Note); !note.Powered {
			return fmt.Errorf("note block was not powered through strongly powered block")
		}

		tx.SetBlock(torchPos, RedstoneTorch{Facing: cube.FaceDown}, &world.SetOpts{DisableBlockUpdates: true})
		updateTorchRedstone(torchPos, tx)
		if note := tx.Block(notePos).(Note); note.Powered {
			return fmt.Errorf("note block stayed powered after torch stopped strongly powering block")
		}
		return nil
	})
}

func TestExternallyClockedRedstoneTorchDoesNotBurnOut(t *testing.T) {
	testRedstoneTorchTx(t, func(tx *world.Tx) error {
		torchPos := cube.Pos{1, 1, 0}
		inputPos := torchPos.Side(cube.FaceWest)

		tx.SetBlock(inputPos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(torchPos, RedstoneTorch{Facing: cube.FaceWest, Lit: true}, &world.SetOpts{DisableBlockUpdates: true})

		for i := 0; i < 10; i++ {
			tx.SetBlock(inputPos, RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.Block(torchPos).(RedstoneTorch).ScheduledTick(torchPos, tx, nil)
			if torch := tx.Block(torchPos).(RedstoneTorch); torch.Lit {
				return fmt.Errorf("torch stayed lit after powered external input on cycle %d", i)
			}

			tx.SetBlock(inputPos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.Block(torchPos).(RedstoneTorch).ScheduledTick(torchPos, tx, nil)
			if torch := tx.Block(torchPos).(RedstoneTorch); !torch.Lit {
				return fmt.Errorf("torch burned out after unpowered external input on cycle %d", i)
			}
		}
		return nil
	})
}

func TestBurnedOutRedstoneTorchRelightsOnNeighbourBlockUpdate(t *testing.T) {
	testRedstoneTorchTx(t, func(tx *world.Tx) error {
		supportPos := cube.Pos{0, 1, 0}
		torchPos := supportPos.Side(cube.FaceEast)
		changedPos := torchPos.Side(cube.FaceUp)

		tx.SetBlock(supportPos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(torchPos, RedstoneTorch{Facing: cube.FaceWest}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(changedPos, Air{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.Redstone().BurnOutTorch(torchPos)

		tx.SetBlock(changedPos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.Block(torchPos).(RedstoneTorch).NeighbourUpdateTick(torchPos, changedPos, tx)

		torch := tx.Block(torchPos).(RedstoneTorch)
		if !torch.Lit {
			return fmt.Errorf("burned out torch did not relight after neighbour block update")
		}
		if burnedOut, _ := tx.Redstone().TorchBurnoutStatus(torchPos, tx.CurrentTick()); burnedOut {
			return fmt.Errorf("burnout state was not cleared after neighbour block update recovery")
		}
		return nil
	})
}
