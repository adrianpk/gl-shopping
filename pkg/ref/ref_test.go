package ref_test

import (
	"github.com/adrianpk/gl-shopping/pkg/ref"
)

var (
	backedBeans = ref.NewItem("BackedBeans", 990)

	biscuits = ref.NewItem("Biscuits", 1200)

	sardines = ref.NewItem("Sardines", 189)

	shampooSmall = ref.NewItem("Shampoo Small", 2000)

	shampooMedium = ref.NewItem("Shampoo Medium", 2500)

	shampooLarge = ref.NewItem("Shampoo Large", 3500)
)

var (
	catalogue = ref.NewCatalogue("Default", []ref.Item{backedBeans, sardines, shampooSmall, shampooMedium, shampooLarge})
)

var (
	basket1 = ref.NewBasket()
	
	basket2 = ref.NewBasket()
)
