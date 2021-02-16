package core

import (
	"fmt"
	"sort"
)

type (
	Item interface {
		ID() interface{}
		Name() string
		Price() float64
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
		CheapestFromSetDiscount() (requiredQty int64)
	}

	Basket interface {
		Items() map[interface{}]int64
	}

	Pricer struct {
		catalogue Catalogue
		offers    []Offer
		basket    Basket
		subtotal  float64
		discount  float64
		total     float64
		singles   []*singleItemDiscountUnit
		multis    []*multiItemsDiscountUnit
	}
)

type (
	singleItemDiscountUnit struct {
		item     Item
		qty      int64
		offer    Offer
		discount float64
	}

	multiItemsDiscountUnit struct {
		items    []itemQty
		offer    Offer
		discount float64
	}

	itemQty struct {
		item Item
		qty  int64
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

func (p *Pricer) Result() (subtotal, discount, total float64, err error) {
	err = p.calcSubtotal()
	if err != nil {
		return subtotal, discount, total, err
	}

	err = p.calcDiscount()
	if err != nil {
		return subtotal, discount, total, err
	}

	p.calcTotal()

	return p.subtotal, p.discount, p.total, nil
}

func (p *Pricer) calcSubtotal() (err error) {
	p.subtotal = 0.0

	for itemID, qty := range p.basket.Items() {

		catItem, ok := p.findItemInCatalogue(itemID)
		if !ok {
			return fmt.Errorf("item '%s' not found in catalogue", itemID)
		}

		p.subtotal = p.subtotal + catItem.Price()*float64(qty)
	}

	return nil
}

func (p *Pricer) calcDiscount() (err error) {
	err = p.collectOffers()
	if err != nil {
		return err
	}

	err = p.processOffers()
	if err != nil {
		return err
	}

	return nil
}

// TODO: Decompose this logic into smaller functions
func (p *Pricer) collectOffers() error {
	p.singles = []*singleItemDiscountUnit{}
	p.multis = []*multiItemsDiscountUnit{}

	for itemID, qty := range p.basket.Items() {

		item, ok := p.findItemInCatalogue(itemID)
		if !ok {
			continue
		}

		for _, offer := range p.offers {
			if p.isItemIncludedInOffer(itemID, offer) {

				switch dt := offer.DiscountType(); dt {

				case Discounts.Percentage:
					stu := &singleItemDiscountUnit{
						item:     item,
						qty:      qty,
						offer:    offer,
						discount: 0.0,
					}

					p.singles = append(p.singles, stu)

				case Discounts.Quantity:
					stu := &singleItemDiscountUnit{
						item:     item,
						qty:      qty,
						offer:    offer,
						discount: 0.0,
					}

					p.singles = append(p.singles, stu)

				case Discounts.CheapestFromSet:

					existent := false
					for _, m := range p.multis {
						if m.offer.ID() == offer.ID() {
							m.items = append(m.items, itemQty{item, qty})
							existent = true
						}
					}

					if existent {
						continue
					}

					mtu := &multiItemsDiscountUnit{
						items:    []itemQty{},
						offer:    offer,
						discount: 0.0,
					}

					mtu.items = append(mtu.items, itemQty{item, qty})
					p.multis = append(p.multis, mtu)

				default:
					// Do nothing
				}
			}
		}
	}
	return nil
}

func (p *Pricer) processOffers() (err error) {
	err = p.processSingleItemDiscounts()
	if err != nil {
		return err
	}

	err = p.processMultipleItemsDiscounts()
	if err != nil {
		return err
	}

	err = p.accumulateDiscounts()
	if err != nil {
		return err
	}

	return nil
}

func (p *Pricer) processSingleItemDiscounts() (err error) {
	for _, stu := range p.singles {

		switch dt := stu.offer.DiscountType(); dt {
		case Discounts.Percentage:
			stu.discount = stu.item.Price() * float64(stu.qty) * stu.offer.PercentageDiscount() / 100

		case Discounts.Quantity:
			buyN, freeN := stu.offer.QuantityDiscount()
			applicableFor := stu.qty / buyN

			stu.discount = stu.item.Price() * float64(applicableFor*freeN)

		default:
			// Do nothing for now
		}

	}

	return nil
}

func (p *Pricer) processMultipleItemsDiscounts() (err error) {
	for _, mtu := range p.multis {

		switch dt := mtu.offer.DiscountType(); dt {
		case Discounts.CheapestFromSet:
			requiredQty := int(mtu.offer.CheapestFromSetDiscount())

			sort.Slice(mtu.items, func(i, j int) bool {
				return mtu.items[i].item.Price() > mtu.items[j].item.Price()
			})

			items := []Item{}
			for _, item := range mtu.items {
				for i := 0; i < int(item.qty); i++ {
					items = append(items, item.item)
				}
			}

			for i := requiredQty - 1; i < len(items); i = i + requiredQty {
				mtu.discount = mtu.discount + items[i].Price()
			}

		default:
			// Do nothing for now
		}

	}

	return nil
}

func (p *Pricer) accumulateDiscounts() (err error) {
	p.discount = 0.0

	for _, stu := range p.singles {
		p.discount = p.discount + stu.discount
	}

	for _, mtu := range p.multis {
		p.discount = p.discount + mtu.discount
	}

	return nil
}

func (p *Pricer) calcTotal() {
	p.total = p.subtotal - p.discount
}

func (p *Pricer) findItemInCatalogue(itemID interface{}) (item Item, ok bool) {
	for _, catItem := range p.catalogue.Items() {
		if catItem.ID() == itemID {
			return catItem, true
		}
	}

	return nil, false
}

func (p *Pricer) isItemIncludedInOffer(itemID interface{}, offer Offer) (ok bool) {
	for _, offerItem := range offer.Items() {
		if offerItem == itemID {
			return true
		}

	}

	return false
}
