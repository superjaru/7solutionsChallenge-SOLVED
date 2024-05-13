package main

import (
	"fmt"
	"math"
	"strconv"
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
	maxChar := countMaxChar(encoded)
	maxRange := int(math.Pow10(len(encoded) + 1))
	for i := 0; i < maxRange; i++ {
		if !isMoreThanMax(maxChar, i) {
			integerStr := fmt.Sprintf("%0*d", len(encoded)+1, i) // create integer with leading 0 ex.002550
			// fmt.Println(integerStr)
			digitSum := calculateDigitSum(integerStr)
			encodedInteger := encode(integerStr, encoded)

			//compare input and new encoded
			if encodedInteger == encoded {
				if digitSum < minDigitSum {
					minDigitSum = digitSum
					result = integerStr
				}
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

func encode(integer, encoded string) string {
	encoded2 := ""
	for i := 0; i < len(integer)-1; i++ {
		if integer[i] > integer[i+1] {
			encoded2 += "L"
		} else if integer[i] < integer[i+1] {
			encoded2 += "R"
		} else {
			encoded2 += "="
		}
		if encoded2[i] != encoded[i] {
			return ""
		}
	}
	return encoded2
}

func countMaxChar(str string) int {
	charCount := make(map[rune]int)
	maxCount := 0
	for _, char := range str {
		charCount[char]++
		if charCount[char] > maxCount {
			maxCount = charCount[char]
		}
	}
	return maxCount
}

func isMoreThanMax(number, target int) bool {
	targetStr := strconv.Itoa(target)
	for _, digit := range targetStr {
		digitInt, _ := strconv.Atoi(string(digit))
		if digitInt > number {
			return true
		}
	}
	return false
}
