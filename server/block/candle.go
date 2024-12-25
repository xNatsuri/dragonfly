package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

// Candle is a dyeable block that emits light when lit with a flint and steel.
// Up to four of the same color of candle can be placed in one block space, which affects the amount of light produced.
type Candle struct {
	empty
	transparent
	sourceWaterDisplacer

	// Candles is the number of candles.
	Candles int
	// Colour is the colour of the candle.
	Colour item.Colour
	// Regular is the regular candles.
	Regular bool
	// Lit is whether the candles are lit.
	Lit bool
}

// BreakInfo ...
func (c Candle) BreakInfo() BreakInfo {
	return newBreakInfo(0.1, alwaysHarvestable, shearsEffective, oneOf(c))
}

// LightEmissionLevel ...
func (c Candle) LightEmissionLevel() uint8 {
	if c.Candles == 0 {
		return 3
	}
	return uint8(c.Candles + 1*3)
}

// SideClosed ...
func (Candle) SideClosed(cube.Pos, cube.Pos, *world.Tx) bool {
	return false
}

// UseOnBlock ...
func (c Candle) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, tx *world.Tx, user item.User, ctx *item.UseContext) (used bool) {
	if existing, ok := tx.Block(pos).(Candle); ok {
		if existing.Colour != c.Colour || existing.Regular != c.Regular {
			return false
		}

		if existing.Candles >= 3 {
			return false
		}

		existing.Candles++
		place(tx, pos, existing, user, ctx)
		return placed(ctx)
	}

	pos, face, used = firstReplaceable(tx, pos, face, c)
	if !used {
		return false
	}

	place(tx, pos, c, user, ctx)
	return placed(ctx)
}

// NeighbourUpdateTick ...
func (c Candle) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if liquid, ok := tx.Liquid(pos); ok {
		if liquid.LiquidType() == "water" {
			if !c.Lit {
				c.Lit = false
				tx.SetBlock(pos, c, nil)
				tx.PlaySound(pos.Vec3Centre(), sound.FireExtinguish{})
			}
		}
	}
}

// Ignite ...
func (c Candle) Ignite(pos cube.Pos, tx *world.Tx, _ world.Entity) bool {
	if c.Lit {
		return false
	}

	if _, ok := tx.Liquid(pos); ok {
		return false
	}

	c.Lit = true
	tx.SetBlock(pos, c, nil)
	tx.PlaySound(pos.Vec3(), sound.Ignite{})
	return true
}

// EncodeItem ...
func (c Candle) EncodeItem() (name string, meta int16) {
	if c.Regular {
		return "minecraft:candle", 0
	}
	return "minecraft:" + c.Colour.String() + "_candle", 0
}

// EncodeBlock ...
func (c Candle) EncodeBlock() (name string, properties map[string]any) {
	if c.Regular {
		return "minecraft:candle", map[string]any{"candles": int32(c.Candles), "lit": c.Lit}
	}
	return "minecraft:" + c.Colour.String() + "_candle", map[string]any{"candles": int32(c.Candles), "lit": c.Lit}
}

// allCandles returns candle blocks with all possible colours.
func allCandles() []world.Block {
	b := make([]world.Block, 0)
	for i := 0; i <= 3; i++ {
		for _, c := range item.Colours() {
			b = append(b, Candle{Colour: c, Candles: i})
			b = append(b, Candle{Colour: c, Candles: i, Lit: true})
		}
		b = append(b, Candle{Candles: i, Regular: true, Lit: true})
		b = append(b, Candle{Candles: i, Regular: true})
	}
	return b
}
