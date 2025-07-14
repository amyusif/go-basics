// profit calculator 

package main

import "fmt"

func main() {
	var  revenue, expenses, tax_rate float64

	getInput("Enter your revenue: ", &revenue)
	getInput("Enter your expense: ", &expenses)
	getInput("Enter your tax rate: ", &tax_rate)
	

	EBT, profit, ratio := calcP(revenue, expenses, tax_rate)

	fmt.Println("Your EBT is:", EBT)


	fmt.Printf("Your profit is: %.2f\nYour ratio is: %.2f", profit,  ratio)

}

func getInput(intro string, val *float64) {
	fmt.Print(intro)
	fmt.Scan(val)
}

func calcP(rev, expense, tax_rate float64) (float64, float64, float64) {
	bt := rev - expense

	p := bt * (1 - tax_rate/100)

	r := bt / p

	return bt, p, r
}

