package item

import (
	"time"

	"github.com/df-mc/dragonfly/server/world"
)

// WindCharge is a projectile that is crafted from breeze rod. Wind charges knock mobs back when thrown and deals 6 damage.
// When thrown at a player, it deals no damage and players won't receive fall damage when they land.
type WindCharge struct{}

// Use ...
func (w WindCharge) Use(tx *world.Tx, user User, ctx *UseContext) bool {
	create := tx.World().EntityRegistry().Config().WindCharge
	opts := world.EntitySpawnOpts{Position: eyePosition(user), Velocity: user.Rotation().Vec3().Mul(2)}
	tx.AddEntity(create(opts, user))

	ctx.SubtractFromCount(1)
	return true
}

// Cooldown ...
func (WindCharge) Cooldown() time.Duration {
	return time.Second / 2
}

// MaxCount ...
func (WindCharge) MaxCount() int {
	return 64
}

// EncodeItem ...
func (WindCharge) EncodeItem() (name string, meta int16) {
	return "minecraft:wind_charge", 0
}
