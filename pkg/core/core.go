package core

type (
	Item interface {
		ID() string
		Name() string
		Price() int64
	}

	Catalogue interface {
		SetItems(items []Item)
		AddItem(item []Item)
		RemoveItem(id interface{}) error
	}

	Offer interface {
		ID() interface{}
		Items() []Item
		DiscountType() DiscountType
		Description() string
		SetQuantityDiscount(buyQty, payQty int64)
		SetPercentageDiscount(percentage float64)
		SetCheapestFromSetDiscount(items []Item, minQty int64)
	}

	Basket interface {
		Items() []Item
	}

	OfferSet interface {
		AddOffer(offer Offer)
		RemoveOffer(offerID interface{}) error
		Offers() []Offer
	}

	Pricer struct {
		catalogue *Catalogue
		offers    *OfferSet
		basket    *Basket
	}
)

func (p *Pricer) SetCatalogue(c *Catalogue) {
	p.catalogue = c
}

func (p *Pricer) SetOffers(os *OfferSet) {
	p.offers = os
}

func (p *Pricer) SetBasket(b *Basket) {
	p.basket = b
}

func (p *Pricer) Discount() int64 {
	return 0
}

func (p *Pricer) Total() int64 {
	return 0
}
