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
	/*
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


		// for loop
		i := 1
		for i <= 3 {
			println(i)
			i++
		}

		for i := 0; i < 10; i++ {
			print(i, " ")
		}

		// infinity loop
		for {
			print("I'm infinite")
		}

		// break and continue are same as C

	*/

	/*
		// Arrays

		var numbers [3]int
		fmt.Printf("%T \n", numbers)
		for index, number := range numbers {
			print("Before : ", index, ":", number, "\n")
			numbers[index] = index * 10
			print("After : ", index, ":", numbers[index], "\n")
		}
		count, error := fmt.Println(numbers) // Print return number of bites written and an error if any
		println((count))

		fmt.Println(error) // <nil>
		print(error, "\n") // (0x0,0x0)>

		another := []int{77, 88, 99}
		fmt.Println(another)

		var matrix [2][3]int
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				matrix[i][j] = i
			}
		}
		fmt.Println("Matrix: ", matrix)

		matrix2 := [2][3]int{
			{1, 5, 9},
			{1, 2, 3},
		}
		fmt.Println("Matrix: ", matrix2)
		var temp [2][3]int = addMatrix(matrix, matrix2)
		fmt.Println(temp)
	*/
	/*
		// Slices

		my_name := make([]string, 5)
		some_extras := make([]string, 3)
		some_extras[0] = "C"
		some_extras[1] = "C++"
		some_extras[2] = "Java"

		fmt.Println(my_name)
		fmt.Println("The length of my_slice is : ", len(my_name))
		my_name[4] = "n"
		my_name[3] = "a"
		my_name[2] = "m"
		my_name[1] = "u"
		my_name[0] = "S"

		my_name = append(my_name, " ", "D", "e")

		fmt.Println(my_name)

		my_name = append(my_name, some_extras...)
		fmt.Println(my_name)

		fmt.Println("The length of my_slice is : ", len(my_name))
		fmt.Println(my_name[:])
		fmt.Println(my_name[4:8])
		fmt.Println(my_name[4:])
		fmt.Println(my_name[:8])
		fmt.Println()

		// Just explore some other functions like copy() etc
	*/
	// Map
	/*
		// To create an empty map, use the builtin `make`:
		// `make(map[key-type]val-type)`.
		m := make(map[string]int)

		// Set key/value pairs using typical `name[key] = val`
		// syntax.
		m["k1"] = 7
		m["k2"] = 13

		// Printing a map with e.g. `fmt.Println` will show all of
		// its key/value pairs.
		fmt.Println("map:", m)
	*/
	laptop_prices := map[string]int{
		"Hp":   20000,
		"Dell": 21000,
	}
	// add another couple of key-value pairs to laptop_prices
	laptop_prices["Acer"] = 19000
	laptop_prices["Apple"] = 50000
	laptop_prices["Lenovo"] = 25000

	fmt.Println(laptop_prices)

	numbers := []int{1, 2, 3, 4}

	for index, item := range numbers {
		fmt.Println(index, ": ", item)
	}

	for key, value := range laptop_prices {
		fmt.Println(key, ": ", value)
	}

}

/*

// this function will take two two matrix and return another matrix (their sum)
func addMatrix(mat1 [2][3]int, mat2 [2][3]int) [2][3]int {
	var temp [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			temp[i][j] = mat1[i][j] + mat2[i][j]
		}
	}
	return temp

}
*/
