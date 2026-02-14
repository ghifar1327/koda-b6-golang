package order

import (
	"fmt"
	"koda-b6-golang/internal/menu"
	"koda-b6-golang/internal/utils"
)

var Cart []Item

func AddCart(id int, qty int) {
	menus, err := menu.FecthMenu()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, m := range menus {
		if m.Id == id {

			for i, c := range Cart {
				if c.Name == m.Name {
					Cart[i].Qty += qty
					Cart[i].Subtotal = Cart[i].Qty * Cart[i].Price
					return
				}

			}

			item := Item{
				Id:       i + 1,
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

func DeleteItem() {
	fmt.Printf("\n\n=========== Hapus Item =============\n\n")

	for i := range len(Cart) {
		fmt.Printf("%d. %s, harga: %d, %d pcs, total: Rp.%d \n", i+1, Cart[i].Name, Cart[i].Price, Cart[i].Qty, Cart[i].Subtotal)
	}
	fmt.Printf("\n\n=========== ========== =============\n\n")

	var newCart []Item
	id := utils.ReadInt("\n\nMasukan input yang ingin dihapus: ")
	
	
	for i, c := range Cart {
		if c.Id != id {
			item := Item{
				Id:       i + 1,
				Name:     c.Name,
				Price:    c.Price,
				Qty:      c.Qty,
				Subtotal: c.Subtotal,
			}
			newCart = append(newCart, item)
			Cart = newCart
			fmt.Printf("\n\nItem berhasil di hapus...\n")
			return
		}
	}
	fmt.Println("Item tidak ditemukan")
}
