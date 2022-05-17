package warehouse

import "errors"

//   - includes GUID, price, title, artist, label, stock level

type cdUID string

type CD struct {
	title, label string
	i            stockedItem
}

type stockedItem struct {
	price float64
	stock uint
}

type WarehouseReceipt struct {
	Title string
	Price float64
	Stock uint
}

type Warehouse struct {
	p     paymentProvider
	stock map[cdUID]*CD
}

var (
	ErrCDNotFound           = errors.New("CD does not exist")
	ErrCDAlreadyInWarehouse = errors.New("a CD with this ID is already in the warehouse")
)

func New() *Warehouse {
	return &Warehouse{
		p:     &dummyPayments{},
		stock: make(map[cdUID]*CD, 0),
	}
}

func (w *Warehouse) Add(cd CD) (cdUID, error) {
	newID := newCDUID(cd)

	if _, exists := w.stock[newID]; exists {
		return "", ErrCDAlreadyInWarehouse
	}

	w.stock[newID] = &cd

	return newID, nil
}

func NewCD(title, label string, price float64, stock uint) CD {
	return CD{
		title: title,
		label: label,
		i: stockedItem{
			price: price,
			stock: stock,
		},
	}
}

func (w *Warehouse) NewStockedItem(price float64, amount uint) stockedItem {
	return stockedItem{price: price, stock: amount}
}

func (w *Warehouse) BatchReceipt(label string, wr ...WarehouseReceipt) error {
	for _, item := range wr {
		cd := NewCD(item.Title, label, item.Price, item.Stock)

		_, err := w.Add(cd)
		if err != nil {
			return err
		}
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

func newCDUID(cd CD) cdUID {
	return cdUID(cd.label + "|" + cd.title)
}

func (w *Warehouse) FindCDByTitle(title string) (CD, error) {
	return CD{}, nil
}
