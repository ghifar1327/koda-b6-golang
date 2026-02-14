package menu

import (
	"fmt"
)

func ShowMenu() {
	menus, err := FecthMenu()
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i, m := range menus {
		fmt.Printf("%d. %s, harga: Rp.%d\n", i+1, m.Name, m.Price)
	}
}
