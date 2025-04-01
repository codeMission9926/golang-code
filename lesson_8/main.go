package main

import "fmt"

var accountBalance float64 = 1000

func main() {

	for {

		fmt.Println("Welcome to the Bank !")
		fmt.Println("what do you want to do")
		fmt.Println("1. accountBalance")
		fmt.Println("2. deposit")
		fmt.Println("3. withdrawal")
		fmt.Println("4. exit")

		var choose int
		fmt.Print("choose your option [1-4]:")
		fmt.Scan(&choose)

		switch choose {

		case 1:
			fmt.Println(accountBalance)

		case 2:
			var deposit float64
			fmt.Print("Enter amount you want to deposit: ")
			fmt.Scan(&deposit)

			if deposit <= 0 {

				fmt.Println("Invalid value, deposit amount greater than zero")
				return
			}
			accountBalance := accountBalance + deposit
			fmt.Println("now your accountbalance is: ", accountBalance)

		case 3:
			var withdrawal float64
			fmt.Print("Enter amount you want to withdrawal: ")
			fmt.Scan(&withdrawal)

			if withdrawal <= 0 {

				fmt.Println("Invalid value, withdrawal amount greater than zero")
				return
			} else if withdrawal > accountBalance {

				fmt.Println("oops ! your account balance is low: ", accountBalance)
				return
			}

			accountBalance := accountBalance - withdrawal
			fmt.Println("now your accountbalance is: ", accountBalance)

		default:
			fmt.Println("Bye !")
			fmt.Println("Thankyou bank ane ke liye !")
			// break
			return

		}

		// if choose == 1 {
		// 	fmt.Println(accountBalance)
		// } else if choose == 2 {
		// 	var deposit float64
		// 	fmt.Print("Enter amount you want to deposit: ")
		// 	fmt.Scan(&deposit)

		// 	if deposit <= 0 {

		// 		fmt.Println("Invalid value, deposit amount greater than zero")
		// 		return
		// 	}
		// 	accountBalance := accountBalance + deposit
		// 	fmt.Println("now your accountbalance is: ", accountBalance)
		// } else if choose == 3 {
		// 	var withdrawal float64
		// 	fmt.Print("Enter amount you want to withdrawal: ")
		// 	fmt.Scan(&withdrawal)

		// 	if withdrawal <= 0 {

		// 		fmt.Println("Invalid value, withdrawal amount greater than zero")
		// 		return
		// 	} else if withdrawal > accountBalance {

		// 		fmt.Println("oops ! your account balance is low: ", accountBalance)
		// 		return
		// 	}

		// 	accountBalance := accountBalance - withdrawal
		// 	fmt.Println("now your accountbalance is: ", accountBalance)
		// } else {

		// 	fmt.Println("Bye !")
		// 	break
		// }
	}
}
