package main

import (
	"fmt"
)

func main() {
	for r := range "각나다" {
		fmt.Println(r)
	}
	fmt.Println(len("가나다"))
}
