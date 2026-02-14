package menu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func FecthMenu() ([]Menu , error) {
	var result []Menu
	
	res , err := http.Get("https://raw.githubusercontent.com/ghifar1327/koda-b6-golang/refs/heads/master/data/menu.json")
	
	if err != nil {
		return nil , fmt.Errorf("failed to fetch")
		
	}else{
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil , fmt.Errorf("Error: cannot parse data body")
		}

		json.Unmarshal(data, &result)

	}
	return result , nil
}