/*
Package ref includes a basic implementation of core interfaces.
*/

package ref

import (
	"errors"

	"github.com/adrianpk/gl-shopping/pkg/core"
	"github.com/google/uuid"
)

type (
	Item struct {
		id    string
		name  string
		price float64
	}

	Catalogue struct {
		id    string
		name  string
		items []core.Item
	}

	Offer struct {
		id                      string
		items                   []interface{}
		discountType            core.DiscountType
		description             string
		quantityDiscount        core.QuantityDiscount
		percentageDiscount      core.PercentageDiscount
		cheapestFromSetDiscount core.CheapestFromSetDiscount
		priority                int8
		cumulative              bool
	}

	Basket struct {
		items map[interface{}]int64
	}
)

func NewItem(name string, price float64) *Item {
	return &Item{
		id:    uuid.NewString(),
		name:  name,
		price: price,
	}
}

func (i *Item) ID() interface{} {
	return i.id
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) Price() float64 {
	return i.price
}

func NewCatalogue(name string, items []core.Item) *Catalogue {
	return &Catalogue{
		id:    uuid.NewString(),
		name:  name,
		items: items,
	}
}

func (c *Catalogue) ID() string {
	return c.id
}

func (c *Catalogue) Name() string {
	return c.name
}

func (c *Catalogue) SetItems(items []core.Item) {
	c.items = items
}

func (c *Catalogue) AddItem(item core.Item) {
	c.items = append(c.items, item)
}

func (c *Catalogue) RemoveItem(id string) error {
	for i, item := range c.items {
		if item.ID() == id {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

func (c *Catalogue) Items() []core.Item {
	return c.items
}

func NewOffer(itemID interface{}, description string) *Offer {
	return &Offer{
		id:          uuid.NewString(),
		items:       []interface{}{itemID},
		description: description,
	}
}

func (o *Offer) ID() interface{} {
	return o.id
}

func (o *Offer) AddItem(itemID interface{}) {
	o.items = append(o.items, itemID)
}

func (o *Offer) Items() []interface{} {
	items := make([]interface{}, len(o.items))
	for i, v := range o.items {
		items[i] = v
	}

	return items
}

func (o *Offer) DiscountType() core.DiscountType {
	return o.discountType
}

func (o *Offer) Description() string {
	return o.description
}

func (o *Offer) QuantityDiscount() (buyQty, freeQty int64) {
	if o.discountType != core.Discounts.Quantity {
		return 0, 0
	}

	return o.quantityDiscount.BuyQty, o.quantityDiscount.FreeQty
}

func (o *Offer) PercentageDiscount() (percentage float64) {
	if o.discountType != core.Discounts.Percentage {
		return 0.0
	}

	return o.percentageDiscount.Percentage
}

func (o *Offer) CheapestFromSetDiscount() (items []core.Item, buyQty int64) {
	if o.discountType != core.Discounts.CheapestFromSet {
		return []core.Item{}, 0
	}

	return o.cheapestFromSetDiscount.Items, o.cheapestFromSetDiscount.BuyQty
}

func (o *Offer) SetQuantityDiscount(buyQty, freeQty int64) {
	o.quantityDiscount = core.QuantityDiscount{
		BuyQty:  buyQty,
		FreeQty: freeQty,
	}

	o.discountType = core.Discounts.Quantity
}

func (o *Offer) SetPercentageDiscount(percentage float64) {
	o.percentageDiscount = core.PercentageDiscount{
		Percentage: percentage,
	}

	o.discountType = core.Discounts.Percentage
}

func (o *Offer) SetCheapestFromSetDiscount(items []core.Item, buyQty int64) {
	var data []core.Item
	for _, v := range items {
		data = append(data, v)
	}

	o.cheapestFromSetDiscount = core.CheapestFromSetDiscount{
		Items:  data,
		BuyQty: buyQty,
	}

	o.discountType = core.Discounts.CheapestFromSet
}

func (b *Basket) Items() map[interface{}]int64 {
	return b.items
}

func NewBasket() *Basket {
	return &Basket{
		items: map[interface{}]int64{},
	}
}

func (b *Basket) AddItem(itemID interface{}, qty int64) {
	b.items[itemID] = b.items[itemID] + qty
}

func (b *Basket) RemoveItem(itemID interface{}, qty int64) {
	b.items[itemID] = b.items[itemID] + qty

	if b.items[itemID] <= 0 {
		delete(b.items, itemID)
	}
}
