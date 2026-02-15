package app

import (
	"fmt"
	"koda-b6-golang/internal/order"
	"koda-b6-golang/internal/utils"
)

func EmadosApp(){
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Print(err, "not found page")
		} 
	}()
	for{
		fmt.Printf("======= Selamat Datang di Emados =========\n\n")
		fmt.Println("1. Lihat Menu")
		fmt.Println("2. Lihat Keranjang")
		fmt.Println("3. Lihat History")
		fmt.Printf("\n\n99. Keluar\n")

		choose := utils.ReadInt("\nmasukan input: ") 

		switch choose {
		case 1 :
			order.AddMenu()
		case 2 :
			order.ShowCart()
		case 3 :
			order.ShowHistory()
		case 99 :
			fmt.Println("Selamat datang kembali")
		default:
			panic("404")
		}
	}
}