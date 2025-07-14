package main

import (
	"fmt"
	"math"
)

func main() {
	var amount float64
	var years float64
	rate := 0.0

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	fmt.Print("Enter rate: ")
	fmt.Scan(&rate)

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&amount)

	total := amount * math.Pow(1+rate/100, years)
	
	fmt.Print("The future value is: ")
	fmt.Print(total)

}
