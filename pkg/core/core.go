package core

type (
	Item interface {
		ID() interface{}
		Name() string
		Price() int64
	}

	Catalogue interface {
		SetItems(items []Item)
		AddItem(item []Item)
		RemoveItem(id interface{})
	}

	Offer interface {
		ID() interface{}
		Items() []Item
		DiscountType() DiscountType
		Description() string
		SetQuantityDiscount(buy, pay int64)
		SetPercentageDiscount(percentage float64)
	}

	Basket interface {
		Items() []Item
	}

	OfferSet interface {
		AddOffer(item Item, offer Offer)
	}

	Pricer interface {
		SetCatalogue(Catalogue)
		SetOffers(OfferSet)
		DumpBasket(basket Basket)
		AddItems(id interface{}, quantity int64)
		AddItem(id interface{}, quantity int64)
		RemoveItem(id interface{}, quantity int64)
		Discount()
		Total()
	}
)
