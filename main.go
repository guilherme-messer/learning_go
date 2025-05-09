package main

import (
	"fmt"
)

func hexadecimalToInt(c rune) int {
	switch {
	case '0' <= c && c <= '9':
		return int(c - '0')
	case 'a' <= c && c <= 'f':
		return int(c - 'a' + 10)
	case 'A' <= c && c <= 'F':
		return int(c - 'A' + 10)
	}
	return -1 // caractere inválido
}

func main() {
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		fmt.Printf("%d\t%s\n", k, v)
	}

	fmt.Println("-----")

	for pos, char := range "Gő!" {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}

	fmt.Println("-----")

	hexadecimalNum := rune('D')
	fmt.Print("Tradução do caractere D de hexadecimal para decimal: ")
	fmt.Print(hexadecimalToInt(hexadecimalNum))
}
