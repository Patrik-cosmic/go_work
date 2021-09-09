package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

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
}
