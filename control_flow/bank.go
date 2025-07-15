package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	for {
		fmt.Print("Welcome to Go Bank")
		fmt.Println("Choose an option: ")
		fmt.Println("1. Check your account balance")
		fmt.Println("2. Deposite Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			feedbackNotification(accountBalance, choice)
		case 2:
			fmt.Print("Enter amount you want to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount")
			} else {
				accountBalance += depositAmount

				feedbackNotification(accountBalance, choice)
			}
			continue
		case 3:
			fmt.Print("How much do you want to withdraw: ")
			var withdrawAmmount float64
			fmt.Scan(&withdrawAmmount)

			if withdrawAmmount > accountBalance {
				fmt.Println("Insufficient Balance")
			} else {
				accountBalance -= withdrawAmmount

				feedbackNotification(accountBalance, choice)
			}
			continue
		default:
			fmt.Println("Goodbye!!")
			break
		}
	}

}

func feedbackNotification(x float64, choice int) {
	if choice == 1 {
		fmt.Printf("Your account balance is: %.2f\n", x)
	} else if choice == 2 {
		fmt.Printf("Deposite success!!, Your new balance is: %.2f\n", x)
	} else if choice == 3 {
		fmt.Printf("Withdrawal successfull!!, Your new balance is: %.2f\n", x)
	}
}
