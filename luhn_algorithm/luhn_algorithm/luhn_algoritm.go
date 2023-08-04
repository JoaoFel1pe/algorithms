package luhnalgorithm

import "fmt"

// Function that transformes a intenger in a list of all his algarisms
func transformNumberIntoList(number int) []int {
	var list []int

	for number > 0 {
		digit := number % 10
		list = append([]int{digit}, list...)
		number /= 10
	}

	return list
}

// Function that returns True if the given number is a valid credit card number
func LuhnAlgorithm(number int) bool {

	// Creates an array to store the numbers and an uint variable that will receive the final sum
	var numbers_arr []int = transformNumberIntoList(number)
	var sum uint

	// A loop to iterate through the array declared before. It starts on the second-last number to the first and skips one number at a time. Ex.: [1 2 3 4 5] -> 4, 2
	for i := len(numbers_arr) - 2; i >= 0; i -= 2 {
		numbers_arr[i] = numbers_arr[i] * 2 // multiplies the numbers by 2
	}

	// A loop to iterate through the array
	for i := 0; i < len(numbers_arr); i++ {

		if numbers_arr[i] > 9 { // if the numbers is greater than 9 (it has 2 digits), sums the algarisms of the number
			numbers_arr[i] = numbers_arr[i]/10 + numbers_arr[i]%10
		}

		sum += uint(numbers_arr[i])

	}

	// if the lenght of the given number is less than 10 or the module of the sum is diferent from 0, the credit card number is invalid
	if len(numbers_arr) < 10 || sum%10 != 0 {
		fmt.Println("The credit card number is invalid")
		return false
	}

	fmt.Println("The credit card number is valid")
	return true // else, the credit card number is valid

}
