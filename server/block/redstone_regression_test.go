package block

import (
	"fmt"
	"os"
	"testing"
	_ "unsafe"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

//go:linkname finaliseBlockRegistry github.com/df-mc/dragonfly/server/world.finaliseBlockRegistry
func finaliseBlockRegistry()

func TestMain(m *testing.M) {
	finaliseBlockRegistry()
	os.Exit(m.Run())
}

func testRedstoneTx(t *testing.T, f func(tx *world.Tx) error) {
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

func TestUnconnectedRedstoneWirePowersHorizontalMechanisms(t *testing.T) {
	testRedstoneTx(t, func(tx *world.Tx) error {
		wirePos := cube.Pos{0, 1, 0}
		notePos := wirePos.Side(cube.FaceEast)

		tx.SetBlock(wirePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, nil)

		Note{}.RedstoneUpdate(notePos, tx)
		if note := tx.Block(notePos).(Note); !note.Powered {
			return fmt.Errorf("note block next to unconnected powered dust was not powered")
		}
		return nil
	})
}

func TestPoweredSolidBlockPropagatesRedstoneUpdates(t *testing.T) {
	testRedstoneTx(t, func(tx *world.Tx) error {
		torchPos := cube.Pos{0, 1, 0}
		stonePos := torchPos.Side(cube.FaceUp)
		notePos := stonePos.Side(cube.FaceEast)

		tx.SetBlock(stonePos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(torchPos, RedstoneTorch{Facing: cube.FaceDown, Lit: true}, &world.SetOpts{DisableBlockUpdates: true})

		updateAroundRedstone(torchPos, tx)
		if note := tx.Block(notePos).(Note); !note.Powered {
			return fmt.Errorf("note block beside indirectly powered stone was not updated")
		}
		return nil
	})
}

func TestWeaklyPoweredSolidBlockPowersAdjacentMechanism(t *testing.T) {
	testRedstoneTx(t, func(tx *world.Tx) error {
		leverPos := cube.Pos{0, 1, 0}
		stonePos := leverPos.Side(cube.FaceEast)
		notePos := stonePos.Side(cube.FaceEast)

		tx.SetBlock(leverPos, Lever{Powered: true, Facing: cube.FaceWest}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(stonePos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

		Note{}.RedstoneUpdate(notePos, tx)
		if note := tx.Block(notePos).(Note); !note.Powered {
			return fmt.Errorf("note block was not powered through a weakly powered solid block")
		}
		return nil
	})
}

func TestRedstoneDustPowersMechanismThroughSolidBlock(t *testing.T) {
	testRedstoneTx(t, func(tx *world.Tx) error {
		wirePos := cube.Pos{0, 1, 0}
		stonePos := wirePos.Side(cube.FaceEast)
		notePos := stonePos.Side(cube.FaceEast)

		tx.SetBlock(wirePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(stonePos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

		Note{}.RedstoneUpdate(notePos, tx)
		if note := tx.Block(notePos).(Note); !note.Powered {
			return fmt.Errorf("note block was not powered through a dust-powered solid block")
		}
		return nil
	})
}

func TestRedstoneBlockDoesNotPowerMechanismThroughSolidBlock(t *testing.T) {
	testRedstoneTx(t, func(tx *world.Tx) error {
		blockPos := cube.Pos{0, 1, 0}
		stonePos := blockPos.Side(cube.FaceEast)
		notePos := stonePos.Side(cube.FaceEast)

		tx.SetBlock(blockPos, RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(stonePos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

		Note{}.RedstoneUpdate(notePos, tx)
		if note := tx.Block(notePos).(Note); note.Powered {
			return fmt.Errorf("note block was incorrectly powered through a redstone block and solid block")
		}
		return nil
	})
}

func TestGlowstoneDoesNotConductRedstonePower(t *testing.T) {
	testRedstoneTx(t, func(tx *world.Tx) error {
		leverPos := cube.Pos{0, 1, 0}
		glowstonePos := leverPos.Side(cube.FaceEast)
		notePos := glowstonePos.Side(cube.FaceEast)

		tx.SetBlock(leverPos, Lever{Powered: true, Facing: cube.FaceWest}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(glowstonePos, Glowstone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

		Note{}.RedstoneUpdate(notePos, tx)
		if note := tx.Block(notePos).(Note); note.Powered {
			return fmt.Errorf("note block was powered through glowstone")
		}
		return nil
	})
}

func TestGlowstoneBlocksSkylight(t *testing.T) {
	testRedstoneTx(t, func(tx *world.Tx) error {
		pos := cube.Pos{0, 10, 0}

		tx.SetBlock(pos, Glowstone{}, nil)
		if y := tx.HighestLightBlocker(pos.X(), pos.Z()); y != pos.Y() {
			return fmt.Errorf("highest light blocker after placing glowstone = %d, want %d", y, pos.Y())
		}
		return nil
	})
}

func TestUnlitRedstoneTorchStillShapesDustConnection(t *testing.T) {
	testRedstoneTx(t, func(tx *world.Tx) error {
		wirePos := cube.Pos{0, 1, 0}
		torchPos := wirePos.Side(cube.FaceEast)
		notePos := wirePos.Side(cube.FaceNorth)

		tx.SetBlock(torchPos, RedstoneTorch{Facing: cube.FaceDown}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(wirePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

		Note{}.RedstoneUpdate(notePos, tx)
		if note := tx.Block(notePos).(Note); note.Powered {
			return fmt.Errorf("dust powered perpendicular receiver despite being shaped by unlit torch")
		}
		return nil
	})
}

func TestRedstoneWireShapeOnlyNeighbourUpdatePropagates(t *testing.T) {
	testRedstoneTx(t, func(tx *world.Tx) error {
		wirePos := cube.Pos{0, 1, 0}
		supportPos := wirePos.Side(cube.FaceDown)
		sourcePos := wirePos.Side(cube.FaceSouth)
		torchPos := wirePos.Side(cube.FaceEast)
		notePos := wirePos.Side(cube.FaceNorth)

		tx.SetBlock(supportPos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(sourcePos, RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(wirePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

		Note{}.RedstoneUpdate(notePos, tx)
		if note := tx.Block(notePos).(Note); !note.Powered {
			return fmt.Errorf("note block was not powered before dust shape changed")
		}

		tx.SetBlock(torchPos, RedstoneTorch{Facing: cube.FaceDown}, &world.SetOpts{DisableBlockUpdates: true})
		RedstoneWire{Power: 15}.NeighbourUpdateTick(wirePos, torchPos, tx)
		if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 15 {
			return fmt.Errorf("redstone wire power = %d, want 15", wire.Power)
		}
		if note := tx.Block(notePos).(Note); note.Powered {
			return fmt.Errorf("note block stayed powered after dust shape changed")
		}
		return nil
	})
}

func TestRedstoneWireStepsDownGlassButNotSlabs(t *testing.T) {
	t.Run("stone", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			wirePos := cube.Pos{0, 1, 0}
			sidePos := wirePos.Side(cube.FaceEast)
			sourcePos := sidePos.Side(cube.FaceUp)

			tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sidePos, Stone{}, nil)
			tx.SetBlock(sourcePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sourcePos.Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

			RedstoneWire{}.RedstoneUpdate(wirePos, tx)
			if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 14 {
				return fmt.Errorf("redstone wire below stone step-down = %d, want 14", wire.Power)
			}
			return nil
		})
	})

	t.Run("glass", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			wirePos := cube.Pos{0, 1, 0}
			sidePos := wirePos.Side(cube.FaceEast)
			sourcePos := sidePos.Side(cube.FaceUp)

			tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sidePos, Glass{}, nil)
			tx.SetBlock(sourcePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sourcePos.Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

			RedstoneWire{}.RedstoneUpdate(wirePos, tx)
			if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 14 {
				return fmt.Errorf("redstone wire below glass step-down = %d, want 14", wire.Power)
			}
			return nil
		})
	})

	t.Run("glass spiral", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			a := func(y int) cube.Pos { return cube.Pos{0, y, 0} }
			b := func(y int) cube.Pos { return cube.Pos{1, y, 0} }
			wires := []cube.Pos{b(4), a(3), b(2), a(1)}

			for _, pos := range []cube.Pos{a(4), a(2), b(3), b(1)} {
				tx.SetBlock(pos, Glass{}, nil)
			}
			tx.SetBlock(a(0), Stone{}, nil)
			tx.SetBlock(wires[0], RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(wires[0].Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})
			for _, pos := range wires[1:] {
				tx.SetBlock(pos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			}

			for i, pos := range wires[1:] {
				RedstoneWire{}.RedstoneUpdate(pos, tx)
				want := 14 - i
				if wire := tx.Block(pos).(RedstoneWire); wire.Power != want {
					return fmt.Errorf("wire %d in glass spiral = %d, want %d", i+1, wire.Power, want)
				}
			}
			return nil
		})
	})

	t.Run("glass spiral upward", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			a := func(y int) cube.Pos { return cube.Pos{0, y, 0} }
			b := func(y int) cube.Pos { return cube.Pos{1, y, 0} }
			wires := []cube.Pos{a(1), b(2), a(3), b(4)}

			for _, pos := range []cube.Pos{a(0), a(2), b(1), b(3)} {
				tx.SetBlock(pos, Glass{}, nil)
			}
			tx.SetBlock(wires[0], RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(wires[0].Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})
			for _, pos := range wires[1:] {
				tx.SetBlock(pos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			}

			updateStrongRedstone(wires[0], tx)
			for i, pos := range wires {
				want := 15 - i
				if wire := tx.Block(pos).(RedstoneWire); wire.Power != want {
					return fmt.Errorf("wire %d in upward glass spiral = %d, want %d", i, wire.Power, want)
				}
			}
			return nil
		})
	})

	t.Run("double slab", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			wirePos := cube.Pos{0, 1, 0}
			sidePos := wirePos.Side(cube.FaceEast)
			sourcePos := sidePos.Side(cube.FaceUp)

			tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sidePos, Slab{Block: Stone{}, Double: true}, nil)
			tx.SetBlock(sourcePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sourcePos.Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

			RedstoneWire{}.RedstoneUpdate(wirePos, tx)
			if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 14 {
				return fmt.Errorf("redstone wire below double slab step-down = %d, want 14", wire.Power)
			}
			return nil
		})
	})

	t.Run("slab", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			wirePos := cube.Pos{0, 1, 0}
			sidePos := wirePos.Side(cube.FaceEast)
			sourcePos := sidePos.Side(cube.FaceUp)

			tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sidePos, Slab{Block: Stone{}, Top: true}, nil)
			tx.SetBlock(sourcePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sourcePos.Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

			RedstoneWire{}.RedstoneUpdate(wirePos, tx)
			if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 0 {
				return fmt.Errorf("redstone wire below slab step-down = %d, want 0", wire.Power)
			}
			return nil
		})
	})

	t.Run("stairs", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			wirePos := cube.Pos{0, 1, 0}
			sidePos := wirePos.Side(cube.FaceEast)
			sourcePos := sidePos.Side(cube.FaceUp)

			tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sidePos, Stairs{Block: Stone{}, Facing: cube.North, UpsideDown: true}, nil)
			tx.SetBlock(sourcePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sourcePos.Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

			RedstoneWire{}.RedstoneUpdate(wirePos, tx)
			if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 0 {
				return fmt.Errorf("redstone wire below stairs step-down = %d, want 0", wire.Power)
			}
			return nil
		})
	})

	t.Run("hopper", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			wirePos := cube.Pos{0, 1, 0}
			sidePos := wirePos.Side(cube.FaceEast)
			sourcePos := sidePos.Side(cube.FaceUp)

			tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sidePos, Hopper{Facing: cube.FaceDown}, nil)
			tx.SetBlock(sourcePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sourcePos.Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

			RedstoneWire{}.RedstoneUpdate(wirePos, tx)
			if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 0 {
				return fmt.Errorf("redstone wire below hopper step-down = %d, want 0", wire.Power)
			}
			return nil
		})
	})

	t.Run("glowstone", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			wirePos := cube.Pos{0, 1, 0}
			sidePos := wirePos.Side(cube.FaceEast)
			sourcePos := sidePos.Side(cube.FaceUp)

			tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sidePos, Glowstone{}, nil)
			tx.SetBlock(sourcePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sourcePos.Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

			RedstoneWire{}.RedstoneUpdate(wirePos, tx)
			if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 0 {
				return fmt.Errorf("redstone wire below glowstone step-down = %d, want 0", wire.Power)
			}
			return nil
		})
	})

	t.Run("sea lantern", func(t *testing.T) {
		testRedstoneTx(t, func(tx *world.Tx) error {
			wirePos := cube.Pos{0, 1, 0}
			sidePos := wirePos.Side(cube.FaceEast)
			sourcePos := sidePos.Side(cube.FaceUp)

			tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sidePos, SeaLantern{}, nil)
			tx.SetBlock(sourcePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(sourcePos.Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

			RedstoneWire{}.RedstoneUpdate(wirePos, tx)
			if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 14 {
				return fmt.Errorf("redstone wire below sea lantern step-down = %d, want 14", wire.Power)
			}
			return nil
		})
	})
}

func TestRedstoneWireDirectAndNetworkUpdatesMatch(t *testing.T) {
	buildSpiral := func(tx *world.Tx) []cube.Pos {
		a := func(y int) cube.Pos { return cube.Pos{0, y, 0} }
		b := func(y int) cube.Pos { return cube.Pos{1, y, 0} }
		wires := []cube.Pos{a(1), b(2), a(3), b(4)}

		for _, pos := range []cube.Pos{a(0), a(2), b(1), b(3)} {
			tx.SetBlock(pos, Glass{}, nil)
		}
		tx.SetBlock(wires[0], RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(wires[0].Side(cube.FaceSouth), RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})
		for _, pos := range wires[1:] {
			tx.SetBlock(pos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
		}
		return wires
	}
	readPowers := func(tx *world.Tx, wires []cube.Pos) []int {
		powers := make([]int, len(wires))
		for i, pos := range wires {
			powers[i] = tx.Block(pos).(RedstoneWire).Power
		}
		return powers
	}

	var directPowers []int
	testRedstoneTx(t, func(tx *world.Tx) error {
		wires := buildSpiral(tx)
		for _, pos := range wires[1:] {
			RedstoneWire{}.RedstoneUpdate(pos, tx)
		}
		directPowers = readPowers(tx, wires)
		return nil
	})

	testRedstoneTx(t, func(tx *world.Tx) error {
		wires := buildSpiral(tx)
		updateStrongRedstone(wires[0], tx)
		networkPowers := readPowers(tx, wires)
		if len(networkPowers) != len(directPowers) {
			return fmt.Errorf("network wire count = %d, want %d", len(networkPowers), len(directPowers))
		}
		for i := range networkPowers {
			if networkPowers[i] != directPowers[i] {
				return fmt.Errorf("wire %d power mismatch: direct=%d network=%d", i, directPowers[i], networkPowers[i])
			}
		}
		return nil
	})
}
