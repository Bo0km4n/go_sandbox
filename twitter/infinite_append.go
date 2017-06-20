package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	x := []int{1, 2, 3}
	for i := range x {
		fmt.Println("==============>")
		fmt.Println("i :=", i)
		fmt.Println("###############")
		pp.Print(x)
		x = append(x, i)
	}
	fmt.Println(x)
}
