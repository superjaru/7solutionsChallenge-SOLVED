package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Test cases
	// fmt.Println(decode("LLRR=")) // Output: 210122
	// fmt.Println(decode("==RLL")) // Output: 000210
	// fmt.Println(decode("=LLRR")) // Output: 221012
	// fmt.Println(decode("RRL=R")) // Output: 012001
	fmt.Println("Enter your encoded input:")
	var input string
	for {
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)
		input = strings.ToUpper(input)
		if isValidInput(input) {
			break
		} else {
			fmt.Println("Invalid input. Please enter only 'L', 'R', or '='.")
		}
	}
	fmt.Println("Your decoded output is ", decode(input))
}

func isValidInput(input string) bool {
	validChars := "LR="
	for _, char := range input {
		if !strings.ContainsRune(validChars, char) {
			return false
		}
	}
	return true
}

func decode(encoded string) string {
	minDigitSum := 9 * len(encoded) //possible max sum
	result := ""

	for i := 0; i < int(math.Pow10(len(encoded)+1)); i++ {
		integerStr := fmt.Sprintf("%0*d", len(encoded)+1, i) // create integer with leading 0 ex.002550
		// fmt.Println(minDigitSum)
		digitSum := calculateDigitSum(integerStr)
		encodedInteger := encode(integerStr)

		//compare input and new encoded
		if encodedInteger == encoded {
			if digitSum < minDigitSum {
				minDigitSum = digitSum
				result = integerStr
			}
		}
	}

	return result
}

func calculateDigitSum(s string) int {
	sum := 0
	for _, digit := range s {
		sum += int(digit - '0')
	}
	return sum
}

func encode(integer string) string {
	encoded := ""
	for i := 0; i < len(integer)-1; i++ {
		if integer[i] > integer[i+1] {
			encoded += "L"
		} else if integer[i] < integer[i+1] {
			encoded += "R"
		} else {
			encoded += "="
		}
	}
	return encoded
}
