package core

import "errors"

type (
	DiscountType struct {
		Name        string
		Description string
	}

	discountTypeSet struct {
		Percentage      DiscountType
		Quantity        DiscountType
		CheapestFromSet DiscountType
		Types           []DiscountType
	}
)

type (
	QuantityDiscount struct {
		BuyQty int64
		PayQty int64
	}

	PercentageDiscount struct {
		Percentage float64
	}

	CheapestFromSetDiscount struct {
		Items []Item
	}
)

var Discounts = newDiscountTypeSet()

var percentage = DiscountType{
	Name:        "percentage",
	Description: "Percentage discount",
}

var quantity = DiscountType{
	Name:        "quantity",
	Description: "Buy N items and pay M",
}

var cheapestFromSet = DiscountType{
	Name:        "cheapest-from-set",
	Description: "Buy N items of a set of products and get the cheapest for free",
}

func newDiscountTypeSet() *discountTypeSet {
	return &discountTypeSet{
		Percentage:      percentage,
		Quantity:        quantity,
		CheapestFromSet: cheapestFromSet,
		Types:           []DiscountType{percentage, quantity, cheapestFromSet},
	}
}

func (ots *discountTypeSet) DiscountByName(name string) (DiscountType, error) {
	for _, dt := range ots.Types {
		if dt.Name == name {
			return dt, nil
		}
	}
	return DiscountType{}, errors.New("not a valid discount type")
}
