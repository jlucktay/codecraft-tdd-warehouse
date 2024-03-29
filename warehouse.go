package warehouse

import "errors"

//   - includes GUID, price, title, artist, label, stock level

type cdUID string

type CD struct {
	artist, title, label string
	i                    stockedItem
}

type stockedItem struct {
	price float64
	stock uint
}

type WarehouseReceipt struct {
	Artist, Title string
	Price         float64
	Stock         uint
}

type Warehouse struct {
	p     PaymentProvider
	stock map[cdUID]*CD
	c     Chart
}

var (
	ErrCDNotFound           = errors.New("CD does not exist")
	ErrCDAlreadyInWarehouse = errors.New("a CD with this ID is already in the warehouse")
)

func NewCD(artist, title, label string, price float64, stock uint) CD {
	return CD{
		artist: artist,
		title:  title,
		label:  label,
		i: stockedItem{
			price: price,
			stock: stock,
		},
	}
}

func (c *CD) GetTitle() string {
	return c.title
}

func (c *CD) GetPrice() float64 {
	return c.i.price
}

func New(c Chart, p PaymentProvider) *Warehouse {
	return &Warehouse{
		c:     c,
		p:     p,
		stock: make(map[cdUID]*CD, 0),
	}
}

func generateCDUID(cd CD) cdUID {
	return cdUID(cd.label + "|" + cd.title)
}

func (w *Warehouse) Add(cd CD) (cdUID, error) {
	newID := generateCDUID(cd)

	if _, exists := w.stock[newID]; exists {
		return "", ErrCDAlreadyInWarehouse
	}

	w.stock[newID] = &cd

	return newID, nil
}

func (w *Warehouse) BatchReceipt(label string, wr ...WarehouseReceipt) error {
	for _, item := range wr {
		cd := NewCD(item.Artist, item.Title, label, item.Price, item.Stock)

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

	w.c.Notify(selling.artist, selling.title, qty)

	selling.i.stock -= qty

	return nil
}

func (w *Warehouse) SellSingleCD(id cdUID) error {
	return w.SellCD(id, 1)
}

func (w *Warehouse) FindCDByTitle(title string) (CD, error) {
	for _, cd := range w.stock {
		if cd.title == title {
			return *cd, nil
		}
	}

	return CD{}, ErrCDNotFound
}
