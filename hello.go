package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	/*
		// Initialization 101
			//type 1:
			var name string
			name = "Suman"
			fmt.Println(name)

			//type 2:
			var name = "Suman"
			fmt.Println(name)

			// another way
			name := "Suman"
			fmt.Println(name)

	*/

	/*

		// Check the type of the variable/identifier. Printf is same as C
		name := "Suman"
		fmt.Printf("The type of variable \"name\" is \"%T\" \n", name)

		// User Input and Output

		var number int
		fmt.Print("Enter a number: ")
		fmt.Scanf("%d", &number)
		fmt.Printf("The number is %d \n", number)
		// another way
		fmt.Printf("The number is %v \n", number)
		print("The number is %v \n", number)

	*/

	// control flow

	number1 := 10
	number2 := 80
	var big int

	// if else
	if number1 > number2 {
		big = number1
	} else {
		big = number2
	}

	fmt.Printf("The biggest number is : %d \n", big)
	// switch
	switch big {
	case 10:
		print("The value of big is 10\n")
	case 20:
		print("The value of big is 20\n")
	default:

		print("The value of big is Unknown!\n")
	}

	value := 4
	switch value {
	case 1, 3, 5, 7, 9:
		print("Odd\n")
	case 2, 4, 6, 8, 10:
		print("Even\n")
	default:
		print("Please Enter number between 1 and 10\n")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
