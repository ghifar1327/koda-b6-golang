package order

import (
	"fmt"
	"koda-b6-golang/internal/menu"
)

var Cart []Item

func AddCart(id int, qty int) {
	menus, err := menu.FecthMenu()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, m := range menus {
		if m.Id == id {

			for i, c := range Cart {
				if c.Name == m.Name {
					Cart[i].Qty += qty
					Cart[i].Subtotal = Cart[i].Qty * Cart[i].Price
					return
				}

			}

			item := Item{
				Name:     m.Name,
				Price:    m.Price,
				Qty:      qty,
				Subtotal: m.Price * qty,
			}

			Cart = append(Cart, item)

			fmt.Println("\nBerhasil ditambahkan ke keranjang")
			return
		}
	}
}