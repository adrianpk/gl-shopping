package ref_test

import (
	"math"
	"os"
	"testing"

	"github.com/adrianpk/gl-shopping/pkg/core"
	"github.com/adrianpk/gl-shopping/pkg/ref"
)

const (
	eqThreshold = 0.1
)

var (
	bakedBeans    = ref.NewItem("BakedBeans", 0.99)
	biscuits      = ref.NewItem("Biscuits", 1.2)
	sardines      = ref.NewItem("Sardines", 1.89)
	shampooSmall  = ref.NewItem("Shampoo Small", 2.0)
	shampooMedium = ref.NewItem("Shampoo Medium", 2.5)
	shampooLarge  = ref.NewItem("Shampoo Large", 3.5)
)

var (
	catalogue = ref.NewCatalogue("Default", []core.Item{bakedBeans, biscuits, sardines, shampooSmall, shampooMedium, shampooLarge})
)

var (
	bakedBeansOffer = ref.NewOffer(bakedBeans.ID(), "Backed Beans Qty")
	sardinesOffer   = ref.NewOffer(sardines.ID(), "Sardines Percentage")
	shampoosOffer   = ref.NewOffer(shampooLarge.ID(), "Shampoos Cheapest From Set")
	offers          = []core.Offer{bakedBeansOffer, sardinesOffer}
)

var (
	basket1 = ref.NewBasket()
	basket2 = ref.NewBasket()
	basket3 = ref.NewBasket()
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestDiscountBasic(t *testing.T) {
	// Setup offers
	bakedBeansOffer.SetQuantityDiscount(2, 1)
	sardinesOffer.SetPercentageDiscount(25.0)

	// Setup basket
	basket1.AddItem(bakedBeans.ID(), 4)
	basket1.AddItem(biscuits.ID(), 1)

	// Setup pricer
	pricer := core.NewPricer(catalogue, offers)
	pricer.SetBasket(basket1)

	// Verify
	subtotal, discount, total, err := pricer.Result()
	if err != nil {
		t.Errorf("cannot calculate result (%e)", err)
	}

	if !eq(subtotal, 5.16) {
		t.Errorf("subtotal should be 5.16 (%.2f)", subtotal)
	}

	if !eq(discount, 1.98) {
		t.Errorf("discount should be 1.98 (%.2f)", discount)
	}

	if !eq(total, 3.18) {
		t.Errorf("total should should be 3.18 (%.2f)", total)
	}
}

func TestDiscountBasic2(t *testing.T) {
	// Setup offers
	bakedBeansOffer.SetQuantityDiscount(2, 1)
	sardinesOffer.SetPercentageDiscount(25.0)

	// Setup basket
	basket2.AddItem(bakedBeans.ID(), 2)
	basket2.AddItem(biscuits.ID(), 1)
	basket2.AddItem(sardines.ID(), 2)

	// Setup pricer
	pricer := core.NewPricer(catalogue, offers)
	pricer.SetBasket(basket2)

	// Verify
	subtotal, discount, total, err := pricer.Result()
	if err != nil {
		t.Errorf("cannot calculate result (%e)", err)
	}

	if !eq(subtotal, 6.96) {
		t.Errorf("subtotal should be 6.96 (%.2f)", subtotal)
	}

	if !eq(discount, 1.94) {
		t.Errorf("discount should be 1.94 (%.2f)", discount)
	}

	if !eq(total, 5.02) {
		t.Errorf("total should should be 5.02 (%.2f)", total)
	}
}

func TestCheapestFromSetDiscount(t *testing.T) {
	// Setup offers
	shampoosOffer.AddItem(shampooMedium.ID())
	shampoosOffer.AddItem(shampooSmall.ID())

	offers2 := []core.Offer{shampoosOffer}

	shampoosOffer.SetCheapestFromSetDiscount(3)

	// Setup basket
	basket3.AddItem(shampooLarge.ID(), 3)
	basket3.AddItem(shampooMedium.ID(), 1)
	basket3.AddItem(shampooSmall.ID(), 2)

	// Setup pricer
	pricer := core.NewPricer(catalogue, offers2)
	pricer.SetBasket(basket3)

	// Verify
	subtotal, discount, total, err := pricer.Result()
	if err != nil {
		t.Errorf("cannot calculate result (%e)", err)
	}

	if !eq(subtotal, 6.96) {
		t.Errorf("subtotal should be 6.96 (%.2f)", subtotal)
	}

	if !eq(discount, 1.94) {
		t.Errorf("discount should be 1.94 (%.2f)", discount)
	}

	if !eq(total, 5.02) {
		t.Errorf("total should should be 5.02 (%.2f)", total)
	}
}

// Helpers
func eq(a, b float64) bool {
	return math.Abs(a-b) <= eqThreshold
}
