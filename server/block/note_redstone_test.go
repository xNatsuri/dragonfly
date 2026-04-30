package block

import (
	"fmt"
	"testing"
	_ "unsafe"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

//go:linkname finaliseBlockRegistryForNoteRedstone github.com/df-mc/dragonfly/server/world.finaliseBlockRegistry
func finaliseBlockRegistryForNoteRedstone()

func init() {
	finaliseBlockRegistryForNoteRedstone()
}

func testNoteRedstoneTx(t *testing.T, f func(tx *world.Tx) error) {
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

func TestPoweredNoteBlockUpdatesAdjacentNoteBlocks(t *testing.T) {
	testNoteRedstoneTx(t, func(tx *world.Tx) error {
		leverPos := cube.Pos{0, 1, 0}
		firstNotePos := leverPos.Side(cube.FaceEast)
		secondNotePos := firstNotePos.Side(cube.FaceEast)

		tx.SetBlock(leverPos, Lever{Powered: true, Facing: cube.FaceWest}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(firstNotePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(secondNotePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

		updateRedstone(firstNotePos, tx)
		if note := tx.Block(firstNotePos).(Note); !note.Powered {
			return fmt.Errorf("first note block was not powered")
		}
		if note := tx.Block(secondNotePos).(Note); !note.Powered {
			return fmt.Errorf("second note block was not updated through powered note block")
		}
		return nil
	})
}
