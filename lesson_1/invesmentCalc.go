package main  // defining main package

// importing built in packages
import (
	"fmt"
	"math"
)

// creating basic main func: To understand variables and types

// func main () {
	 
// 	// defining variables
//      var investmentAmount float64 = 2000  
// 	 var years float64 =5
// 	 var expectedReturnRate float64 = 8.5

// 	 // calc final Amount
// 	 var finalAmount = investmentAmount*math.Pow((1+(expectedReturnRate)/100), years)

//     // printing final Amount
// 	 fmt.Println(finalAmount)


// }


// shorthand
// func main () {
	 
//      investmentAmount,years, expectedReturnRate := 2000, 5, 8.5  

// 	 var finalAmount = investmentAmount*math.Pow((1+(expectedReturnRate)/100), years)

// 	 fmt.Println(finalAmount)


// }



// Taking values from user:

func main () {
	 
	
     var investmentAmount float64 
	 var years float64
	 const expectedReturnRate float64 = 8.5

	 fmt.Print ("Investment Amount: ")  // Asking user to put a value
	 fmt.Scan(&investmentAmount)        // taking value from user using scan func

     fmt.Print("Years: ")
	 fmt. Scan(&years)

	 var finalAmount = investmentAmount*math.Pow((1+(expectedReturnRate)/100), years)

    // printing final Amount
	 fmt.Println(finalAmount)


}


