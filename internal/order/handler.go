package order

import (
	"fmt"
	"koda-b6-golang/internal/menu"
	"koda-b6-golang/internal/utils"
)

func AddMenu() {
	for {
		fmt.Printf("\n=========== Pilih Menu =============\n\n")
		menu.ShowMenu()
		fmt.Printf("\n====================================\n")
		chose := utils.ReadInt("Masukan input menu (1/10): ")
		if chose == 99 {
			return
		}
		if chose > 0 && chose <= 10 {
			qty := utils.ReadInt("\nMau berapa pcs? ")
			AddCart(chose, qty)
			close := utils.Confirm("\nMau menu tambah lagi?")
			if !close {
				return
			}
		} else {
			fmt.Printf("\nInput tidak sesuai...\n\n")
			utils.ReadEnter("\nTekan Enter untuk Input Ulang...\n")
		}
	}
}

func ShowCart() {
	for {
		fmt.Printf("\n=========== keranjang =============\n\n")
		if len(Cart) != 0 {
			var total int
			for i := range len(Cart) {
				fmt.Printf("%d. %s, harga: %d, %d pcs, total: Rp.%d \n", Cart[i].Id, Cart[i].Name, Cart[i].Price, Cart[i].Qty, Cart[i].Subtotal)
			}
			for i := range len(Cart){
				total+= Cart[i].Subtotal
			}
			fmt.Printf("\nTotal: Rp. %d" ,total)
			fmt.Printf("\n===================================\n")

			fmt.Printf("\n1. Checkout\n")
			fmt.Println("0. Hapus item")
			fmt.Printf("\n99. Exit\n")
			chose := utils.ReadInt("\nMasukan input: ")
			switch chose {
			case 1:
				Checkout()
				return
			case 0:
				DeleteItem()
			case 99:
				return
			}
		} else {
			fmt.Printf("Keranjang masih kosong!\n\n")
			utils.ReadEnter("Tekan Enter Untuk Keluar...")
			return
		}
	}
}

func ShowHistory(){
	fmt.Printf("\n=========== HISTORY PESANAN ===========\n")
	GetHistory()
	utils.ReadEnter("Tekan Enter untuk kembali...")

}