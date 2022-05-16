package warehouse

//   - includes GUID, price, title, artist, label, stock level

type cdUID string

type CD struct {
	title, artist, label string
	price                float64
	stock                uint
}

type Warehouse struct {
	stock map[cdUID]*CD
}

func New() *Warehouse {
	return &Warehouse{
		stock: make(map[cdUID]*CD, 0),
	}
}

func (w *Warehouse) Add(cd CD) error {
	return nil
}

func (w *Warehouse) NewCD(title, artist, label string, price float64, stock uint) cdUID {
	newID := newCDUID()

	w.stock[newID] = &CD{
		title:  title,
		artist: artist,
		label:  label,
		price:  price,
		stock:  stock,
	}

	return newID
}

func (w *Warehouse) QueryStockLevel(id cdUID) uint {
	return w.stock[id].stock
}

func (w *Warehouse) SellCD(id cdUID) {
	w.stock[id].stock--
}

func newCDUID() cdUID {
	return "asdf"
}
