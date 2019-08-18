package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(convertToRoman(i))
	}
}

func convertToRoman(a int) (z string) {
	var result string

	for a > 0 {
		switch {
		case a == 100:
			a = 0
			result = result + "C"

		case a >= 90:
			a = a - 90
			result = result + "XC"

		case a >= 50:
			a = a - 50
			result = result + "L"

		case a >= 40:
			a = a - 40
			result = result + "XL"

		case a >= 10:
			a = a - 10
			result = result + "X"
		
		case a >= 9:
			a = a - 9
			result = result + "IX"

		case a >= 5:
			a = a - 5
			result = result + "V"

		case a >= 4:
			a = a - 4
			result = result + "IV"

		default:
			a = a - 1
			result = result + "I"
		}
	}
	return result
}
