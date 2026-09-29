// Covert celsius to fahrenheit and km to miles
package main

import (
	"fmt"
)

func main() {
	var choice string

	fmt.Println("Choose a conversion:")
	fmt.Println("1. Celcius to Fahrenheit")
	fmt.Println("2. Kilometers to miles")

	fmt.Scan(&choice)

	switch choice {
	default:
		fmt.Println("Invalid Choice.")

	case "1":
		var celcius float32

		fmt.Print("Enter celcius: ")
		_, err := fmt.Scan(&celcius)

		if err != nil {
			fmt.Println("Input must be an number.")
			return
		}

		fmt.Printf("Coverted to Fahrenheit: %.2f°F\n", temperature(celcius))

	case "2":
		var kilometers float32

		fmt.Print("Enter kilometers: ")
		_, err := fmt.Scan(&kilometers)
		if err != nil {
			fmt.Println("Input must be an number.")
			return
		}
		fmt.Printf("Coverted to miles: %.2f miles\n", distance(kilometers))
	}
}

func temperature(c float32) float32 {
	return (c * 9 / 5) + 32
}

func distance(km float32) float32 {
	return km / 1.609
}
