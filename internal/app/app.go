package app

import (
	"fmt"
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

		choose := utils.ReadInt("masukan input: ") 

		switch choose {
		case 1 :
			fmt.Println("menu")
		case 2 :
			fmt.Println("Keranjang")
		case 3 :
			fmt.Println("history")
		case 99 :
			fmt.Println("Selamat datang kembali")
		default:
			panic("404")
		}
	}
}