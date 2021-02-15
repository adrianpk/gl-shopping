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
	offers          = []core.Offer{bakedBeansOffer, sardinesOffer}
)

var (
	basket1 = ref.NewBasket()
	basket2 = ref.NewBasket()
)

var ()

func TestMain(m *testing.M) {
	setup()
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
	st, err := pricer.SubTotal()
	if err != nil {
		t.Errorf("cannot calculate subtotal (%e)", err)
	}

	if !eq(st, 5.16) {
		t.Errorf("subtotal should be 5.16 (%.2f)", st)
	}

	d, err := pricer.Discount()
	if err != nil {
		t.Errorf("cannot calculate discount (%e)", err)

	}

	if !eq(d, 1.98) {
		t.Errorf("discount should be 1.98 (%.2f)", d)
	}

	total, err := pricer.Total()
	if err != nil {
		t.Errorf("cannot calculate total (%e)", err)

	}

	if !eq(total, 3.18) {
		t.Errorf("total should should be 3.18 (%.2f)", d)
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
	st, err := pricer.SubTotal()
	if err != nil {
		t.Errorf("cannot calculate subtotal (%e)", err)
	}

	if !eq(st, 6.96) {
		t.Errorf("subtotal should be 6.96 (%.2f)", st)
	}

	d, err := pricer.Discount()
	if err != nil {
		t.Errorf("cannot calculate discount (%e)", err)

	}

	if !eq(d, 1.94) {
		t.Errorf("discount should be 1.94 (%.2f)", d)
	}

	total, err := pricer.Total()
	if err != nil {
		t.Errorf("cannot calculate total (%e)", err)

	}

	if !eq(total, 5.02) {
		t.Errorf("total should should be 5.02 (%.2f)", d)
	}
}

func setup() {
}

// Helpers
func eq(a, b float64) bool {
	return math.Abs(a-b) <= eqThreshold
}
