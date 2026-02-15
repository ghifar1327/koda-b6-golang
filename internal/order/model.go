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
	Total     int
	CreatedAt time.Time
}