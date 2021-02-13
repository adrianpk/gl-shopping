package core_test

import (
	"errors"

	"github.com/adrianpk/gl-shopping/pkg/core"
)

type (
	Item struct {
		id    string
		name  string
		price int64
	}

	Catalogue struct {
		items []Item
	}

	Offer struct {
		id                      string
		items                   []Item
		discountType            core.DiscountType
		description             string
		quantityDiscount        core.QuantityDiscount
		percentageDiscount      core.PercentageDiscount
		cheapestFromSetDiscount core.CheapestFromSetDiscount
	}

	Basket struct {
		items []Item
	}

	OfferSet struct {
		offers []Offer
	}
)

func (i *Item) ID() string {
	return i.id
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) Price() int64 {
	return i.price
}

func (c *Catalogue) SetItems(items []Item) {
	c.items = items
}

func (c *Catalogue) AddItem(item Item) {
	c.items = append(c.items, item)
}

func (c *Catalogue) RemoveItem(id string) error {
	for i, item := range c.items {
		if item.id == id {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

func (o *Offer) ID() string {
	return o.id
}

func (o *Offer) Items() []Item {
	return o.items
}

func (o *Offer) DiscountType() core.DiscountType {
	return o.discountType
}

func (o *Offer) Description() string {
	return o.description
}

func (o *Offer) SetQuantityDiscount(buyQty, payQty int64) {
	o.quantityDiscount = core.QuantityDiscount{
		BuyQty: buyQty,
		PayQty: payQty,
	}

	o.discountType = core.Discounts.Quantity
}

func (o *Offer) SetPercentageDiscount(percentage float64) {
	o.percentageDiscount = core.PercentageDiscount{
		Percentage: percentage,
	}

	o.discountType = core.Discounts.Percentage
}

func (o *Offer) SetCheapestFromSetDiscount(items []*Item, minQty int64) {
	var data []core.Item
	for _, v := range items {
		data = append(data, v)
	}

	o.cheapestFromSetDiscount = core.CheapestFromSetDiscount{
		Items: data,
	}

	o.discountType = core.Discounts.CheapestFromSet
}

func (b *Basket) Items() []Item {
	return b.items
}

func (os *OfferSet) AddOffer(offer Offer) {
	os.offers = append(os.offers, offer)
}

func (os *OfferSet) RemoveOffer(offerID string) error {
	for i, offer := range os.offers {
		if offer.ID() == offerID {
			os.offers = append(os.offers[:i], os.offers[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}
