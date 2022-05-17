package warehouse

import "errors"

//   - includes GUID, price, title, artist, label, stock level

type cdUID string

type CD struct {
	/*title, artist,*/ label string
	i                        StockedItem
}

type StockedItem struct {
	price float64
	stock uint
}

type Warehouse struct {
	p     paymentProvider
	stock map[cdUID]*CD
}

var ErrCDNotFound = errors.New("CD does not exist")

func New() *Warehouse {
	return &Warehouse{
		p:     &dummyPayments{},
		stock: make(map[cdUID]*CD, 0),
	}
}

func (w *Warehouse) Add(cd CD) error {
	return nil
}

func (w *Warehouse) NewCD( /*title, artist,*/ label string, price float64, stock uint) cdUID {
	newID := newCDUID()

	w.stock[newID] = &CD{
		// title:  title,
		// artist: artist,
		label: label,
		i: StockedItem{
			price: price,
			stock: stock,
		},
	}

	return newID
}

func (w *Warehouse) NewStockedItem(price float64, amount uint) StockedItem {
	return StockedItem{price: price, stock: amount}
}

func (w *Warehouse) BatchReceipt(label string, si ...StockedItem) error {
	for _, item := range si {
		w.NewCD(label, item.price, item.stock)
	}

	return nil
}

func (w *Warehouse) QueryStockLevel(id cdUID) uint {
	return w.stock[id].i.stock
}

func (w *Warehouse) SellCD(id cdUID, qty uint) error {
	selling, exists := w.stock[id]
	if !exists {
		return ErrCDNotFound
	}

	if err := w.p.Sell(selling.i.price * float64(qty)); err != nil {
		return err
	}

	selling.i.stock -= qty

	return nil
}

func (w *Warehouse) SellSingleCD(id cdUID) error {
	return w.SellCD(id, 1)
}

func newCDUID() cdUID {
	return "asdf"
}
