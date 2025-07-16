// profit calculator

package main

import "fmt"

import "os"

func saveCalc(bt, p, r float64) {
	beforeTax, profit, ratio := fmt.Sprint(bt), fmt.Sprint(p), fmt.Sprint(r)
	os.WriteFile("profCalc.txt", []byte(beforeTax), 0644)
	os.WriteFile("profCalc.txt", []byte(profit), 0644)
	os.WriteFile("profCalc.txt", []byte(ratio), 0644)

	fmt.Println("file created")
}

func main() {
	var  revenue, expenses, tax_rate float64

	fmt.Println("Profit Calculator")
	fmt.Println("----------------------------")

	var choice string


	for {

		getInput("Enter your revenue: ", &revenue)
		getInput("Enter your expense: ", &expenses)
		getInput("Enter your tax rate: ", &tax_rate)
		
		EBT, profit, ratio := calcP(revenue, expenses, tax_rate)
		
		// fmt.Println("Your EBT is:", EBT)
		
		// fmt.Printf("Your profit is: %.2f\nYour ratio is: %.2f\n", profit,  ratio)

		saveCalc(EBT, profit, ratio)

		fmt.Print("Do you wanna perform a new calculation? Y or N : ")
		fmt.Scan(&choice)

		if choice == "y" || choice == "Y" {
			continue
		} else if  choice == "n" || choice == "N"{
			break
		} else {
			fmt.Println("Wrong choice input!!!!!")
			break
		}
	}



}

func getInput(intro string, val *float64) {

	fmt.Print(intro)
	fmt.Scan(val)

	if *val <= 0 {
		fmt.Println("Value should be greater than 0!!")

	}

	
}

func calcP(rev, expense, tax_rate float64) (float64, float64, float64) {
	
	bt := rev - expense

	p := bt * (1 - tax_rate/100)

	r := bt / p

	return bt, p, r
}

