package entity

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/cube/trace"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/particle"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

// NewWindCharge creates a new WindCharge entity with the given options.
func NewWindCharge(opts world.EntitySpawnOpts, owner world.Entity) *world.EntityHandle {
	conf := windChargeConf
	conf.Owner = owner.H()
	return opts.New(WindChargeType, conf)
}

var windChargeConf = ProjectileBehaviourConfig{
	Gravity:               0.00,
	Drag:                  0.01,
	Particle:              particle.WindCharge{},
	SurviveBlockCollision: false,
	Sound:                 sound.WindCharge{},
	Hit:                   windChargeHit,
}

// windChargeHit knocks nearby entities back with a force proportional to the distance from the charge.
func windChargeHit(e *Ent, tx *world.Tx, target trace.Result) {
	pos := target.Position()

	const (
		explosionRadius = 3.0
		knockbackForce  = 1.2
	)

	nearby := tx.EntitiesWithin(e.H().Type().BBox(e).Translate(pos).Grow(explosionRadius))
	for victim := range nearby {
		victimPos := victim.Position()
		delta := victimPos.Sub(pos)
		distance := delta.Len()

		if distance > explosionRadius || distance == 0 {
			continue
		}

		direction := delta.Mul(1 / distance)
		falloff := 1 - (distance / explosionRadius)
		strength := knockbackForce * falloff

		if living, ok := victim.(interface {
			SetVelocity(mgl64.Vec3)
			Velocity() mgl64.Vec3
		}); ok {
			currentVel := living.Velocity()
			knockback := direction.Mul(strength)

			knockback[1] += strength * 0.5
			living.SetVelocity(currentVel.Add(knockback))
		}
	}

	// TODO: interact with blocks (buttons, doors, etc.)
}

// WindChargeType is a world.EntityType implementation for WindCharge.
var WindChargeType windChargeType

type windChargeType struct{}

func (t windChargeType) Open(tx *world.Tx, handle *world.EntityHandle, data *world.EntityData) world.Entity {
	return &Ent{tx: tx, handle: handle, data: data}
}

func (windChargeType) EncodeEntity() string { return "minecraft:wind_charge_projectile" }
func (windChargeType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.3125, 0, -0.3125, 0.3125, 0.3125, 0.3125)
}
func (windChargeType) DecodeNBT(_ map[string]any, data *world.EntityData) {
	data.Data = windChargeConf.New()
}
func (windChargeType) EncodeNBT(*world.EntityData) map[string]any { return nil }
