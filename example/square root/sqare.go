
// package main
// import "fmt"

// func main() {

// 	fmt.Println(square(5))
// }


// func square (a int)  int {
// 	var sq int = a*a

// 	return sq



// }


// package main 

// import "fmt"

// func main() {

// 	fmt.Println(add(6,5))
// }


// func add (a int ,b int) int  {


// 	var ad int  = a * b 

// 	return ad
// }



// package main //it's the entry point of your program// 

// import "fmt" // inbuilt package called //

// func main () { // phela function yaha sha shru hota ha //

//   fmt.Println( multiply(9))

// }


// func multiply (a int ) int {

// 	var x int = a*a


// 	return x
// }


package main  

import "fmt" 

func main () { 

  fmt.Println( multiply(9))
  
  fmt.Println(Addition(4))

  fmt.Println(subtraction(5))

  fmt.Println(divided (8,8))
}
func divided (c int) int  {

	var did int = c/c

	return did
}


func subtraction (b int ) int {

	var sb int = b-b

	return sb

}


func Addition (a int) int {

	var z int = a+a

	return z
}
func multiply (a int ) int  {

	var x int = a*a


	return x
}



