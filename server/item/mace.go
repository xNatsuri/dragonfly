package item

import "time"

type Mace struct{}

// MaxCount always returns 1.
func (m Mace) MaxCount() int {
	return 1
}

// DurabilityInfo ...
func (m Mace) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability:    501,
		BrokenItem:       simpleItem(Stack{}),
		AttackDurability: 1,
		BreakDurability:  1,
	}
}

// AttackDamage ...
func (m Mace) AttackDamage() float64 {
	return 5
}

// Cooldown ...
func (Mace) Cooldown() time.Duration {
	return time.Second / 2
}

// EncodeItem ...
func (m Mace) EncodeItem() (name string, meta int16) {
	return "minecraft:mace", 0
}
