package main

import (
	"fmt"
)

func main() {

	n := 20
	fizzbuzz(n)

}

func fizzbuzz(N int) {

	n := N
	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {

			fmt.Println("fizzbuzz")

		} else if i%3 == 0 {

			fmt.Println("fizz")
		} else if i%5 == 0 {

			fmt.Println("buzz")
		} else {

			fmt.Println(i)
		}
	}

}
