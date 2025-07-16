package main

import (
	"fmt"
	"os"
)
func getName(num float64) {
	numText := fmt.Sprint(num)
	os.WriteFile("name.txt", []byte(numText), 0644)
}

func main() {
	num := 12.6

	getName(num)

	fmt.Print("file created")
}