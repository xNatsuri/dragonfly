package block

import (
	"fmt"
	"testing"
	_ "unsafe"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

//go:linkname finaliseBlockRegistryForRedstoneWireBreak github.com/df-mc/dragonfly/server/world.finaliseBlockRegistry
func finaliseBlockRegistryForRedstoneWireBreak()

func init() {
	finaliseBlockRegistryForRedstoneWireBreak()
}

func testRedstoneWireBreakTx(t *testing.T, f func(tx *world.Tx) error) {
	t.Helper()

	w := world.Config{Entities: testRedstoneWireBreakEntityRegistry}.New()
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

var testRedstoneWireBreakEntityRegistry = world.EntityRegistryConfig{
	Item: func(opts world.EntitySpawnOpts, _ any) *world.EntityHandle {
		return opts.New(testRedstoneWireBreakEntityType{}, testRedstoneWireBreakEntityConfig{})
	},
}.New(nil)

type testRedstoneWireBreakEntityConfig struct{}

func (testRedstoneWireBreakEntityConfig) Apply(*world.EntityData) {}

type testRedstoneWireBreakEntityType struct{}

func (testRedstoneWireBreakEntityType) Open(_ *world.Tx, handle *world.EntityHandle, data *world.EntityData) world.Entity {
	return testRedstoneWireBreakEntity{handle: handle, pos: data.Pos, rot: data.Rot}
}

func (testRedstoneWireBreakEntityType) EncodeEntity() string { return "test:item" }

func (testRedstoneWireBreakEntityType) BBox(world.Entity) cube.BBox { return cube.BBox{} }

func (testRedstoneWireBreakEntityType) DecodeNBT(map[string]any, *world.EntityData) {}

func (testRedstoneWireBreakEntityType) EncodeNBT(*world.EntityData) map[string]any { return nil }

type testRedstoneWireBreakEntity struct {
	handle *world.EntityHandle
	pos    mgl64.Vec3
	rot    cube.Rotation
}

func (e testRedstoneWireBreakEntity) Close() error { return nil }

func (e testRedstoneWireBreakEntity) H() *world.EntityHandle { return e.handle }

func (e testRedstoneWireBreakEntity) Position() mgl64.Vec3 { return e.pos }

func (e testRedstoneWireBreakEntity) Rotation() cube.Rotation { return e.rot }

func TestUnsupportedRedstoneWireUnpowersHorizontalMechanisms(t *testing.T) {
	testRedstoneWireBreakTx(t, func(tx *world.Tx) error {
		supportPos := cube.Pos{0, 0, 0}
		wirePos := supportPos.Side(cube.FaceUp)
		notePos := wirePos.Side(cube.FaceEast)

		tx.SetBlock(supportPos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(wirePos, RedstoneWire{Power: 15}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

		Note{}.RedstoneUpdate(notePos, tx)
		if note := tx.Block(notePos).(Note); !note.Powered {
			return fmt.Errorf("note block next to powered dust was not powered before support broke")
		}

		tx.SetBlock(supportPos, nil, &world.SetOpts{DisableBlockUpdates: true})
		RedstoneWire{Power: 15}.NeighbourUpdateTick(wirePos, supportPos, tx)
		if note := tx.Block(notePos).(Note); note.Powered {
			return fmt.Errorf("note block stayed powered after unsupported dust broke")
		}
		return nil
	})
}

func TestRedstoneWireNeighbourUpdateUsesCancellableHandler(t *testing.T) {
	testRedstoneWireBreakTx(t, func(tx *world.Tx) error {
		supportPos := cube.Pos{0, 0, 0}
		wirePos := supportPos.Side(cube.FaceUp)
		sourcePos := wirePos.Side(cube.FaceEast)
		handler := &testRedstoneUpdateHandler{cancel: map[cube.Pos]bool{wirePos: true}}
		tx.World().Handle(handler)

		tx.SetBlock(supportPos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(sourcePos, RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

		RedstoneWire{}.NeighbourUpdateTick(wirePos, sourcePos, tx)
		if !handler.called[wirePos] {
			return fmt.Errorf("redstone handler was not called for dust neighbour update")
		}
		if wire := tx.Block(wirePos).(RedstoneWire); wire.Power != 0 {
			return fmt.Errorf("cancelled dust neighbour update changed wire power to %d", wire.Power)
		}
		return nil
	})
}

func TestRedstoneWireNetworkUsesCancellableHandler(t *testing.T) {
	testRedstoneWireBreakTx(t, func(tx *world.Tx) error {
		supportPos := cube.Pos{0, 0, 0}
		wirePos := supportPos.Side(cube.FaceUp)
		nextWirePos := wirePos.Side(cube.FaceWest)
		sourcePos := wirePos.Side(cube.FaceEast)
		handler := &testRedstoneUpdateHandler{cancel: map[cube.Pos]bool{nextWirePos: true}}
		tx.World().Handle(handler)

		tx.SetBlock(supportPos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(supportPos.Side(cube.FaceWest), Stone{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(wirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(nextWirePos, RedstoneWire{}, &world.SetOpts{DisableBlockUpdates: true})
		tx.SetBlock(sourcePos, RedstoneBlock{}, &world.SetOpts{DisableBlockUpdates: true})

		updateRedstone(wirePos, tx)
		if !handler.called[nextWirePos] {
			return fmt.Errorf("redstone handler was not called for network dust update")
		}
		if wire := tx.Block(nextWirePos).(RedstoneWire); wire.Power != 0 {
			return fmt.Errorf("cancelled network dust update changed wire power to %d", wire.Power)
		}
		return nil
	})
}

func TestLeverDetachDoesNotDuplicateBreakRedstoneUpdates(t *testing.T) {
	noteUpdateCount := func(detach bool) (int, error) {
		leverPos := cube.Pos{0, 1, 0}
		supportPos := leverPos.Side(cube.FaceWest)
		notePos := leverPos.Side(cube.FaceEast)
		handler := &testRedstoneUpdateHandler{}

		err := runLeverBreakTx(func(tx *world.Tx) error {
			tx.World().Handle(handler)

			tx.SetBlock(supportPos, Stone{}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(leverPos, Lever{Powered: true, Facing: cube.FaceEast}, &world.SetOpts{DisableBlockUpdates: true})
			tx.SetBlock(notePos, Note{}, &world.SetOpts{DisableBlockUpdates: true})

			Note{}.RedstoneUpdate(notePos, tx)
			if note := tx.Block(notePos).(Note); !note.Powered {
				return fmt.Errorf("note block was not powered before lever broke")
			}

			lever := Lever{Powered: true, Facing: cube.FaceEast}
			tx.SetBlock(supportPos, nil, &world.SetOpts{DisableBlockUpdates: true})
			if detach {
				lever.NeighbourUpdateTick(leverPos, supportPos, tx)
			} else {
				breakBlock(lever, leverPos, tx)
			}
			if note := tx.Block(notePos).(Note); note.Powered {
				return fmt.Errorf("note block stayed powered after lever broke")
			}
			return nil
		})
		return handler.counts[notePos], err
	}

	directBreakCount, err := noteUpdateCount(false)
	if err != nil {
		t.Fatal(err)
	}
	detachCount, err := noteUpdateCount(true)
	if err != nil {
		t.Fatal(err)
	}
	if detachCount != directBreakCount {
		t.Fatalf("lever detach note block redstone update count = %d, want direct break count %d", detachCount, directBreakCount)
	}
}

func runLeverBreakTx(f func(tx *world.Tx) error) error {
	w := world.Config{Entities: testRedstoneWireBreakEntityRegistry}.New()
	var txErr error
	<-w.Exec(func(tx *world.Tx) {
		txErr = f(tx)
	})
	if err := w.Close(); err != nil {
		if txErr != nil {
			return fmt.Errorf("%v; close world: %w", txErr, err)
		}
		return fmt.Errorf("close world: %w", err)
	}
	return txErr
}

type testRedstoneUpdateHandler struct {
	world.NopHandler
	called map[cube.Pos]bool
	counts map[cube.Pos]int
	cancel map[cube.Pos]bool
}

func (h *testRedstoneUpdateHandler) HandleRedstoneUpdate(ctx *event.Context[*world.Tx], pos cube.Pos) {
	if h.called == nil {
		h.called = make(map[cube.Pos]bool)
	}
	if h.counts == nil {
		h.counts = make(map[cube.Pos]int)
	}
	h.called[pos] = true
	h.counts[pos]++
	if h.cancel[pos] {
		ctx.Cancel()
	}
}
