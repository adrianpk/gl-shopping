package core

import "errors"

type (
	DiscountType struct {
		Name        string
		Description string
	}

	discountTypeSet struct {
		Percentage  *DiscountType
		BuyNPayM    *DiscountType
		BuyNFromSet *DiscountType
		Types       []*DiscountType
	}
)

var Discounts = newDiscountTypeSet()

var percentage = &DiscountType{
	Name:        "percentage",
	Description: "Percentage discount",
}

var buyNPayM = &DiscountType{
	Name:        "buy-n-pay-m",
	Description: "Buy N items and pay M",
}

var buyNFromSet = &DiscountType{
	Name:        "buy-n-from-set",
	Description: "Buy N items of a set of products and get the cheapest for free",
}

func newDiscountTypeSet() *discountTypeSet {
	return &discountTypeSet{
		Percentage:  percentage,
		BuyNPayM:    buyNPayM,
		BuyNFromSet: buyNFromSet,
		Types:       []*DiscountType{percentage, buyNPayM, buyNFromSet},
	}
}

func (ots *discountTypeSet) DiscountByName(name string) (*DiscountType, error) {
	for _, dt := range ots.Types {
		if dt.Name == name {
			return dt, nil
		}
	}
	return nil, errors.New("not a valid offer type")
}
