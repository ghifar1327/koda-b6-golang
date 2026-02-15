package order

import "time"

type OrderItem struct {
	Id       int
	Name     string
	Price    int
	Qty      int
	Subtotal int
}


type Transaction struct {
	Id        int
	Items     []OrderItem
	GrandTotal     int
	CreatedAt time.Time
}

func (o *OrderItem) Calc() int {
	o.Subtotal = o.Price * o.Qty
	return o.Subtotal
}

func (t *Transaction) Calc() int {
	var total int
	for i := range t.Items {
		total += t.Items[i].Calc()
	}
	t.GrandTotal = total
	return t.GrandTotal
}


type Calculate interface{
	Calc() int
}