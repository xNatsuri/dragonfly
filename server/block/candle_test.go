package block

import (
	"testing"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type testUser struct {
	tx      *world.Tx
	held    item.Stack
	offHand item.Stack
}

func (u *testUser) Close() error                              { return nil }
func (u *testUser) H() *world.EntityHandle                    { return nil }
func (u *testUser) Position() mgl64.Vec3                      { return mgl64.Vec3{} }
func (u *testUser) Rotation() cube.Rotation                   { return cube.Rotation{} }
func (u *testUser) HeldItems() (item.Stack, item.Stack)       { return u.held, u.offHand }
func (u *testUser) SetHeldItems(mainHand, offHand item.Stack) { u.held, u.offHand = mainHand, offHand }
func (u *testUser) UsingItem() bool                           { return false }
func (u *testUser) ReleaseItem()                              {}
func (u *testUser) UseItem()                                  {}

func (u *testUser) PlaceBlock(pos cube.Pos, b world.Block, ctx *item.UseContext) {
	u.tx.SetBlock(pos, b, nil)
	ctx.SubtractFromCount(1)
}

func TestCandleLightEmissionLevel(t *testing.T) {
	tests := []struct {
		name     string
		candle   Candle
		expected uint8
	}{
		{
			name:     "unlit",
			candle:   Candle{},
			expected: 0,
		},
		{
			name:     "one lit candle",
			candle:   Candle{Lit: true},
			expected: 3,
		},
		{
			name:     "four lit candles",
			candle:   Candle{AdditionalCandles: 3, Lit: true},
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.candle.LightEmissionLevel(); got != tt.expected {
				t.Fatalf("LightEmissionLevel() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestCandleModelUsesActualCandleCount(t *testing.T) {
	got := Candle{AdditionalCandles: 3}.Model()
	expected := model.Candle{Count: 4}
	if got != expected {
		t.Fatalf("Model() = %#v, expected %#v", got, expected)
	}
}

func TestCandleUseOnBlockRequiresSolidSupport(t *testing.T) {
	w := world.New()
	defer w.Close()

	pos := cube.Pos{0, 64, 0}
	target := pos.Side(cube.FaceEast)
	<-w.Exec(func(tx *world.Tx) {
		tx.SetBlock(pos, Stone{}, nil)
		u := &testUser{tx: tx, held: item.NewStack(Candle{}, 1)}
		used := Candle{}.UseOnBlock(pos, cube.FaceEast, mgl64.Vec3{}, tx, u, &item.UseContext{})
		if used {
			t.Fatal("UseOnBlock() unexpectedly succeeded without support below the target position")
		}
		if _, ok := tx.Block(target).(Air); !ok {
			t.Fatalf("target block = %#v, expected air", tx.Block(target))
		}
	})
}

func TestCandleUseOnBlockOnlyAddsToCandleAboveWhenClickingTopFace(t *testing.T) {
	w := world.New()
	defer w.Close()

	pos := cube.Pos{0, 64, 0}
	above := pos.Side(cube.FaceUp)
	<-w.Exec(func(tx *world.Tx) {
		tx.SetBlock(pos, Stone{}, nil)
		tx.SetBlock(above, Candle{}, nil)

		u := &testUser{tx: tx, held: item.NewStack(Candle{}, 1)}
		used := Candle{}.UseOnBlock(pos, cube.FaceEast, mgl64.Vec3{}, tx, u, &item.UseContext{})
		if used {
			t.Fatal("UseOnBlock() unexpectedly succeeded when clicking the side of the supporting block")
		}

		candle, ok := tx.Block(above).(Candle)
		if !ok {
			t.Fatalf("block above = %#v, expected candle", tx.Block(above))
		}
		if candle.AdditionalCandles != 0 {
			t.Fatalf("AdditionalCandles = %v, expected 0", candle.AdditionalCandles)
		}
	})
}

func TestCandleUseOnBlockAddsToCandleAboveWhenClickingTopFace(t *testing.T) {
	w := world.New()
	defer w.Close()

	pos := cube.Pos{0, 64, 0}
	above := pos.Side(cube.FaceUp)
	<-w.Exec(func(tx *world.Tx) {
		tx.SetBlock(pos, Stone{}, nil)
		tx.SetBlock(above, Candle{}, nil)

		u := &testUser{tx: tx, held: item.NewStack(Candle{}, 1)}
		used := Candle{}.UseOnBlock(pos, cube.FaceUp, mgl64.Vec3{}, tx, u, &item.UseContext{})
		if !used {
			t.Fatal("UseOnBlock() did not add to the candle above the clicked block")
		}

		candle, ok := tx.Block(above).(Candle)
		if !ok {
			t.Fatalf("block above = %#v, expected candle", tx.Block(above))
		}
		if candle.AdditionalCandles != 1 {
			t.Fatalf("AdditionalCandles = %v, expected 1", candle.AdditionalCandles)
		}
	})
}
