package main

import (
	"fmt"
	"strconv"
)

func main() {
	FizzBuzz(100)
}

func FizzBuzz(num int) string {
	for i := 0; i < num; i++ {
		fizz := i % 3
		buzz := i % 5
		switch {
		case fizz == 0 && buzz == 0:
			fmt.Println("fizz buzz")
		case fizz == 0:
			fmt.Println("fizz")
		case buzz == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(strconv.Itoa(i))
		}
	}

	return "\nEnd"
}
