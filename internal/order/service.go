package order

import (
	"fmt"
	"koda-b6-golang/internal/menu"
	"koda-b6-golang/internal/utils"
	"time"
)

var Cart []OrderItem

var Service struct {
	Items   []OrderItem
	History []Transaction
}

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

			item := OrderItem{
				Id:       len(Cart) + 1,
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
		fmt.Printf("%d. %s, harga: %d, %d pcs, total: Rp.%d \n", Cart[i].Id, Cart[i].Name, Cart[i].Price, Cart[i].Qty, Cart[i].Subtotal)
	}
	fmt.Printf("\n=========== =========== =============\n\n")

	id := utils.ReadInt("\n\nMasukan input yang ingin dihapus: ")
	var newCart []OrderItem
	for _, c := range Cart {
		if c.Id != id {
			item := OrderItem{
				Id:    len(newCart) + 1,
				Name:  c.Name,
				Price: c.Price, Qty: c.Qty, Subtotal: c.Subtotal}
			newCart = append(newCart, item)
		}
	}
	Cart = newCart
	fmt.Printf("\n\nItem berhasil di hapus...\n")
}

func Checkout() {
	var total int
	for i := range len(Cart) {
		total += Cart[i].Subtotal
	}
	fmt.Printf("\nTotal Pesanan Anda: Rp. %d", total)
	chose := utils.Confirm("\n\nLanjutkan Pemabayaran")

	if chose {
		transaction := Transaction{
			Id:        len(Service.History) + 1,
			Items:     Cart,
			Total:     total,
			CreatedAt: time.Now(),
		}
		Service.History = append(Service.History, transaction)
		Cart = []OrderItem{}
	} else {
		return
	}
	fmt.Printf("\nYeay... Pesananan Berhasil!\n")	
	utils.ReadEnter("\nTekan Enter Untuk Keluar...")
}


func GetHistory(){
	for _ , h := range Service.History {
		fmt.Printf("\nTransaksi #%d | Tanggal: %s\n\n",h.Id, h.CreatedAt.Format("2006-01-02 15:04:05"))
		for _, i := range h.Items{
			fmt.Printf("- %s x%d = Rp.%d\n", i.Name, i.Qty, i.Subtotal)
		}
		fmt.Printf("\nTotal: Rp.%d\n", h.Total)
		fmt.Println("--------------------------------")
	}

}