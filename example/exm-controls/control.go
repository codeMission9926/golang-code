package main

import "fmt"

func main() {

	var states int = 5000

	fmt.Println("Welcome to the railway!")
	fmt.Println("where do you want to go")
	fmt.Println("1. states")
	fmt.Println("2. uttarakhand")
	fmt.Println("3. punjab")
	fmt.Println("4. exit")

	var choose int
	fmt.Print("choose your option [1-4]:")
	fmt.Scan(&choose)

	if choose == 1 {
		fmt.Println(states)
	} else if choose == 2 {
		var uttarakhand int 
		fmt.Print("where is the place you want to go uttarakhand: ")
		fmt.Scan(&uttarakhand)

		if uttarakhand  >= 345{

			fmt.Println("Invalid value, uttarakhand amount greater than zero")
			return
		}
		states := states+ uttarakhand
		fmt.Println("now your states is: ", states)
	} else if choose == 3 {
		var punjab int 
		fmt.Print("Tell me the place where  you want to go punjab: ")
		fmt.Scan(&punjab)

		if punjab >= 2{

			fmt.Println("Invalid value, punjab amount greater than zero")
			return
		} else if punjab > states {

			fmt.Println("oops ! your states  is low: ",states)
			return
		}

		states:= states - punjab
		fmt.Println("now your states is: ", states)
	} else {

		fmt.Println("Bye !")
	}

}