package core

import "fmt"

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
		SetQuantityDiscount(buyQty, payQty int64)
		SetPercentageDiscount(percentage float64)
		SetCheapestFromSetDiscount(items []Item, minQty int64)
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

func (p *Pricer) Discount() int64 {
	return 0
}

func (p *Pricer) SubTotal() (subtotal int64, err error) {

	for id, qty := range p.basket.Items() {

		catItem, err := p.findItemInCatalogue(id)
		if err != nil {
			return subtotal, err
		}

		subtotal = subtotal + catItem.Price()*qty
	}

	return subtotal, nil
}

func (p *Pricer) DiscountTotal() int64 {
	return 0
}

func (p *Pricer) Total() int64 {
	return 0
}

func (p *Pricer) findItemInCatalogue(itemID interface{}) (item Item, err error) {
	for _, catItem := range p.catalogue.Items() {
		if catItem.ID() == itemID {
			return catItem, nil
		}

	}

	return nil, fmt.Errorf("item '%s' not found in catalogue", itemID)
}
