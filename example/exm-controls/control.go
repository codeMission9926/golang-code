package main

import "fmt"

var states int = 5000

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


	
	







// 	if choose == 1 {
// 		fmt.Println(states)
// 	} else if choose == 2 {
// 		var uttarakhand int 
// 		fmt.Print("where is the place you want to go uttarakhand: ")
// 		fmt.Scan(&uttarakhand)

// 		if uttarakhand  >= 345{

// 			fmt.Println("Invalid value, uttarakhand amount greater than zero")
// 			return
// 		}
// 		states := states+ uttarakhand
// 		fmt.Println("now your states is: ", states)
// 	} else if choose == 3 {
// 		var punjab int 
// 		fmt.Print("Tell me the place where  you want to go punjab: ")
// 		fmt.Scan(&punjab)

// 		if punjab >= 2{

// 			fmt.Println("Invalid value, punjab amount greater than zero")
// 			return
// 		} else if punjab > states {

// 			fmt.Println("oops ! your states  is low: ",states)
// 			return
// 		}

// 		states:= states - punjab
// 		fmt.Println("now your states is: ", states)
// 	} else {

// 		fmt.Println("Bye !")
// 	}

// }