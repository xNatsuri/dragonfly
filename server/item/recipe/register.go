package recipe

import (
	"slices"
)

// recipes is a list of each recipe.
var recipes []Recipe

// Recipes returns each recipe in a slice.
func Recipes() []Recipe {
	return slices.Clone(recipes)
}

func Register(r Recipe) {
	recipes = append(recipes, r)
}
