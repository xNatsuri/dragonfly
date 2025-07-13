package enchantment

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// Piercing is an enchantment applied to a crossbow that causes arrows
// to pierce through entities.
var Piercing piercing

type piercing struct{}

func (piercing) Name() string {
	return "Piercing"
}

func (piercing) MaxLevel() int {
	return 4
}

// Cost ...
func (piercing) Cost(level int) (int, int) {
	minCost := level * 25
	return minCost, minCost + 50
}

// Rarity ...
func (piercing) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// SurviveEntity ...
func (piercing) SurviveEntity() bool {
	return true
}

// CompatibleWithEnchantment ...
func (piercing) CompatibleWithEnchantment(t item.EnchantmentType) bool {
	//TODO: multi shot does not work with this enchantment.
	return true
}

// CompatibleWithItem ...
func (piercing) CompatibleWithItem(i world.Item) bool {
	_, ok := i.(item.Crossbow)
	return ok
}
