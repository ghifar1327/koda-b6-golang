package order

import (
	"koda-b6-golang/internal/menu"
	"koda-b6-golang/internal/utils"
)


func AddMenu(){
	for{
		menu.ShowMenu()
		chose := utils.ReadInt("\nMasukan input menu (1/10): ")
		if chose > 0 && chose <= 10{
			qty := utils.ReadInt("\nMau berapa pcs? ")
			AddCart(chose, qty)
			close := utils.Confirm("\nMau menu tambah lagi?")
			if (!close){
				return
			}
		}
	}
}