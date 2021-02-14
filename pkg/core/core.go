package core

import (
	"fmt"
)

type (
	Item interface {
		ID() interface{}
		Name() string
		Price() int64
	}

	Catalogue interface {
		Items() []Item
	}

	Offer interface {
		ID() interface{}
		Items() []interface{}
		DiscountType() DiscountType
		Description() string
		QuantityDiscount() (buyQty, payQty int64)
		PercentageDiscount() (percentage float64)
		CheapestFromSetDiscount() (itemIDs []Item)
	}

	Basket interface {
		Items() map[interface{}]int64
	}

	Pricer struct {
		catalogue Catalogue
		offers    []Offer
		basket    Basket
	}
)

func NewPricer(c Catalogue, o []Offer) *Pricer {
	return &Pricer{
		catalogue: c,
		offers:    o,
	}
}

func (p *Pricer) SetCatalogue(c Catalogue) {
	p.catalogue = c
}

func (p *Pricer) SetOffers(offers []Offer) {
	p.offers = offers
}

func (p *Pricer) SetBasket(b Basket) {
	p.basket = b
}

func (p *Pricer) SubTotal() (subtotal int64, err error) {
	for itemID, qty := range p.basket.Items() {

		catItem, err := p.findItemInCatalogue(itemID)
		if err != nil {
			return subtotal, err
		}

		subtotal = subtotal + catItem.Price()*qty
	}

	return subtotal, nil
}

func (p *Pricer) Discount() (discount int64, err error) {
	for itemID, qty := range p.basket.Items() {

		item, err := p.findItemInCatalogue(itemID)
		if err != nil {
			return discount, err
		}

		offer, ok := p.findOffer(itemID)
		if !ok {
			continue
		}

		switch dt := offer.DiscountType(); dt {
		case Discounts.Percentage:
			// TODO: Implement some helpers to make this inline conversions
			discount = discount + int64(float64(discount)+float64((item.Price()*qty))*offer.PercentageDiscount())

		case Discounts.Quantity:
			buyN, freeN := offer.QuantityDiscount()
			applicableFor := qty / buyN

			// TODO: Implement some helpers to make this inline conversions
			discount = discount + int64(float64(item.Price())*float64(applicableFor*freeN))

		case Discounts.CheapestFromSet:
			// TODO: implement it
			discount = discount

		default:
			// Do nothing for now
			discount = discount
		}
	}

	return discount, nil
}

func (p *Pricer) Total() (total int64, err error) {
	subtotal, err := p.SubTotal()
	if err != nil {
		return total, err
	}

	discount, err := p.Discount()
	if err != nil {
		return total, err
	}

	return subtotal - discount, err
}

func (p *Pricer) findItemInCatalogue(itemID interface{}) (item Item, ok error) {
	for _, catItem := range p.catalogue.Items() {
		if catItem.ID() == itemID {
			return catItem, nil
		}

	}

	return nil, fmt.Errorf("item '%s' not found in catalogue", itemID)
}

func (p *Pricer) findOffer(itemID interface{}) (offer Offer, ok bool) {
	for _, offer := range p.offers {
		if findItem(itemID, offer.Items()) {
			return offer, true
		}
	}

	return nil, false
}

func findItem(itemID interface{}, items []interface{}) (ok bool) {
	for _, i := range items {
		if i == itemID {
			return true
		}
	}

	return false
}
