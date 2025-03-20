package main

import "fmt"

func main() {

	revenue, expenses, taxRate := takeInput() // taking input from user

	EBT, PROFIT, RATIO := logic(revenue, expenses, taxRate) // calculation ebt,profit, ration

	ebt, profit, ration := presicion(EBT, PROFIT, RATIO) // getting precised value

	fmt.Print(ebt, profit, ration) //printing
}

func outputText(text string) {

	fmt.Print(text)
}

func takeInput() (float64, float64, float64) {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("taxRate: ")
	fmt.Scan(&taxRate)

	return revenue, expenses, taxRate
}

func logic(revenue, expenses, taxRate float64) (float64, float64, float64) {

	var ebt = revenue - expenses

	var profit = revenue - expenses - ((revenue - expenses) * taxRate / 100)

	var ratio = ebt / profit

	return ebt, profit, ratio

}

func presicion(EBT, PROFIT, RATIO float64) (string, string, string) {

	ebt := fmt.Sprintf("ebt: %.2f\n", EBT)
	profit := fmt.Sprintf("profit: %.2f\n", PROFIT)
	ratio := fmt.Sprintf("ratio: %.2f\n", RATIO)

	return ebt, profit, ratio

}
