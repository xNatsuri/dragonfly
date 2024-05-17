package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

type SnowLayer struct {
	gravityAffected
	solid

	// Layers the number of layers.
	Layers int
	// Covered is if the snow is covering a plant
	Covered bool
}

// Activate ...
func (s SnowLayer) Activate(pos cube.Pos, _ cube.Face, w *world.World, u item.User, _ *item.UseContext) bool {
	h, _ := u.HeldItems()
	if _, ok := h.Item().(SnowLayer); ok {
		if s.Layers < 6 {
			s.Layers++
			w.SetBlock(pos, s, nil)
		} else {
			w.SetBlock(pos, Snow{}, nil)
		}

		return true
	}

	if b, ok := h.Item().(world.Block); ok && !s.Covered {
		w.SetBlock(pos, b, nil)
		return true
	}

	return false
}

// NeighbourUpdateTick ...
func (s SnowLayer) NeighbourUpdateTick(pos, _ cube.Pos, w *world.World) {
	s.fall(s, pos, w)
}

// BreakInfo ...
func (s SnowLayer) BreakInfo() BreakInfo {
	return newBreakInfo(0.1, alwaysHarvestable, shovelEffective, func(t item.Tool, enchantments []item.Enchantment) []item.Stack {
		switch s.Layers {
		case 3, 4:
			return []item.Stack{item.NewStack(item.Snowball{}, 2)}
		case 5, 6:
			return []item.Stack{item.NewStack(item.Snowball{}, 3)}
		default:
			return []item.Stack{item.NewStack(item.Snowball{}, 1)}
		}
	})
}

// EncodeItem ...
func (SnowLayer) EncodeItem() (name string, meta int16) {
	return "minecraft:snow_layer", 0
}

// EncodeBlock ...
func (s SnowLayer) EncodeBlock() (string, map[string]any) {
	return "minecraft:snow_layer", map[string]any{
		"height":      int32(s.Layers),
		"covered_bit": s.Covered,
	}
}

// allSnowLayers ...
func allSnowLayers() (snowLayers []world.Block) {
	for l := 0; l <= 6; l++ {
		snowLayers = append(snowLayers, SnowLayer{Layers: l})
	}
	return
}
