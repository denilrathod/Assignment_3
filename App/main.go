package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	}
	fmt.Println("Program exited.")
}