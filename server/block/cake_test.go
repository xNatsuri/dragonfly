package block

import (
	"testing"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

func TestCakeWithCandleDropsCandleWhenBroken(t *testing.T) {
	drops := Cake{Candle: true, CandleColour: item.NewOptionalColour(item.ColourRed())}.BreakInfo().Drops(item.ToolNone{}, nil)
	if len(drops) != 1 {
		t.Fatalf("drop count = %v, expected 1", len(drops))
	}
	if _, ok := drops[0].Item().(Candle); !ok {
		t.Fatalf("drop = %#v, expected candle", drops[0].Item())
	}
}

func TestCakeWithLitCandleExtinguishesWhenWaterlogged(t *testing.T) {
	w := world.New()
	defer w.Close()

	pos := cube.Pos{0, 64, 0}
	<-w.Exec(func(tx *world.Tx) {
		cake := Cake{Candle: true, CandleLit: true}
		tx.SetBlock(pos.Side(cube.FaceDown), Stone{}, nil)
		tx.SetBlock(pos, cake, nil)
		tx.SetLiquid(pos, Water{Still: true, Depth: 8})

		cake.NeighbourUpdateTick(pos, pos.Side(cube.FaceNorth), tx)
		got, ok := tx.Block(pos).(Cake)
		if !ok {
			t.Fatalf("block = %#v, expected cake", tx.Block(pos))
		}
		if got.CandleLit {
			t.Fatal("candle cake is still lit after waterlogging")
		}
	})
}

func TestCakeWithCandleCannotIgniteWhenWaterlogged(t *testing.T) {
	w := world.New()
	defer w.Close()

	pos := cube.Pos{0, 64, 0}
	<-w.Exec(func(tx *world.Tx) {
		cake := Cake{Candle: true}
		tx.SetBlock(pos.Side(cube.FaceDown), Stone{}, nil)
		tx.SetBlock(pos, cake, nil)
		tx.SetLiquid(pos, Water{Still: true, Depth: 8})

		if cake.Ignite(pos, tx, nil) {
			t.Fatal("Ignite() unexpectedly succeeded for a waterlogged candle cake")
		}
		got, ok := tx.Block(pos).(Cake)
		if !ok {
			t.Fatalf("block = %#v, expected cake", tx.Block(pos))
		}
		if got.CandleLit {
			t.Fatal("candle cake is lit after Ignite() returned false")
		}
	})
}
