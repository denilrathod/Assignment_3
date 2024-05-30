package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to convert Fahrenheit to Celsius
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// Function to convert Celsius to Fahrenheit
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter temperature (e.g., 32 F or 100 C) or 'exit' to quit: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "exit" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input format. Please enter temperature and unit.")
			continue
		}

		value, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Invalid temperature value. Please enter a numeric value.")
			continue
		}

		unit := strings.ToUpper(parts[1])
		switch unit {
		case "F":
			celsius := FahrenheitToCelsius(value)
			fmt.Printf("%.2f F is %.2f C\n", value, celsius)
		case "C":
			fahrenheit := CelsiusToFahrenheit(value)
			fmt.Printf("%.2f C is %.2f F\n", value, fahrenheit)
		default:
			fmt.Println("Invalid unit. Please use 'C' for Celsius or 'F' for Fahrenheit.")
		}


	}
	fmt.Println("Program exited.")
}
