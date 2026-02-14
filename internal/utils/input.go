package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadString(prompt string) string {
	for {
		fmt.Print(prompt)

		var input string
		fmt.Scanln(&input)

		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input tidak boleh kosong")
			continue
		}

		return input
	}
}

func ReadInt(prompt string) int {
	for {
		fmt.Print(prompt)

		var input string
		fmt.Scanln(&input)

		number, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Masukkan angka yang valid")
			continue
		}

		return number
	}
}

func Confirm(prompt string) bool {
	for {
		fmt.Print(prompt + " (y/n): ")

		var input string
		fmt.Scanln(&input)

		input = strings.ToLower(strings.TrimSpace(input))

		if input == "y" {
			return true
		}
		if input == "n" {
			return false
		}

		fmt.Println("Masukkan y atau n")
	}
}

func ReadEnter(prompt string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}