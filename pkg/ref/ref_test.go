package ref_test

import (
	"os"
	"testing"

	"github.com/adrianpk/gl-shopping/pkg/core"
	"github.com/adrianpk/gl-shopping/pkg/ref"
)

var (
	bakedBeans    = ref.NewItem("BakedBeans", 990)
	biscuits      = ref.NewItem("Biscuits", 1200)
	sardines      = ref.NewItem("Sardines", 1890)
	shampooSmall  = ref.NewItem("Shampoo Small", 2000)
	shampooMedium = ref.NewItem("Shampoo Medium", 2500)
	shampooLarge  = ref.NewItem("Shampoo Large", 3500)
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

	if st != 5160 {
		t.Errorf("subtotal should be 5160 (%d)", st)
	}

	d, err := pricer.Discount()
	if err != nil {
		t.Errorf("cannot calculate discount (%e)", err)

	}

	if d != 1980 {
		t.Errorf("discount should be 1980 (%d)", d)
	}

	total, err := pricer.Total()
	if err != nil {
		t.Errorf("cannot calculate total (%e)", err)

	}

	if total != 3180 {
		t.Errorf("total should should be 3180 (%d)", d)
	}
}

func setup() {
}
