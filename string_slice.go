package main

import "fmt"

func main() {
	var slice []*string
	strings := []string{"test", "sample"}
	fmt.Println(&strings[0])
	slice = append(slice, &strings[0])
	fmt.Println(slice)
}
