package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/internal/nbtconv"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"math/rand"
	"time"
)

// Target is a block that provides redstone power based on how close a projectile hits its center.
type Target struct {
	solid

	// Power is the redstone power level emitted by the target block, ranging from 0 to 15.
	Power int
}

// Source ...
func (t Target) Source() bool {
	return true
}

// WeakPower ...
func (t Target) WeakPower(cube.Pos, cube.Face, *world.World, bool) int {
	return t.Power
}

// StrongPower ...
func (t Target) StrongPower(_ cube.Pos, _ cube.Face, _ *world.World, _ bool) int {
	return t.Power
}

// HitByProjectile handles when a projectile hits the target block.
func (t Target) HitByProjectile(pos mgl64.Vec3, blockPos cube.Pos, w *world.World, delay time.Duration) {
	center := blockPos.Vec3Centre()
	distance := pos.Sub(center).Len()

	maxDistance := math.Sqrt(0.75)
	normalizedDistance := math.Min(distance/maxDistance, 1.0)

	var rawPower float64
	if normalizedDistance <= 0.58 {
		rawPower = 15
	} else if normalizedDistance > 0.9 {
		rawPower = 0
	} else {
		rawPower = 15 * (1 - math.Pow(normalizedDistance, 4.5))
	}

	newBlock := Target{Power: int(math.Max(math.Round(rawPower), 0))}
	w.SetBlock(blockPos, newBlock, &world.SetOpts{DisableBlockUpdates: false})
	w.PlaySound(blockPos.Vec3Centre(), sound.PowerOn{})

	w.ScheduleBlockUpdate(blockPos, delay)
}

// ScheduledTick ...
func (t Target) ScheduledTick(pos cube.Pos, w *world.World, _ *rand.Rand) {
	if t.Power > 0 {
		t.Power = 0
		w.SetBlock(pos, t, nil)
		w.PlaySound(pos.Vec3Centre(), sound.PowerOff{})
	}
}

// DecodeNBT ...
func (t Target) DecodeNBT(data map[string]any) any {
	t.Power = int(nbtconv.Int32(data, "Power"))
	return t
}

// EncodeNBT ...
func (t Target) EncodeNBT() map[string]any {
	m := map[string]any{
		"Power": int32(t.Power),
	}
	return m
}

// EncodeItem ...
func (t Target) EncodeItem() (name string, meta int16) {
	return "minecraft:target", 0
}

// EncodeBlock ...
func (t Target) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:target", nil
}