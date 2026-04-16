package main

import (
	"fmt"
)

func main() {
	expr := "7*8+2"
	for _, ch := range expr {
		fmt.Println(string(ch))
	}
}
