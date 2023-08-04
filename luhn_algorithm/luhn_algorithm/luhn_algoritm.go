package luhnalgorithm

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

	var numbers_arr []int = transformNumberIntoList(number)
	var sum uint

	for i := len(numbers_arr) - 2; i >= 0; i -= 2 {
		numbers_arr[i] = numbers_arr[i] * 2
	}

	for i := 0; i < len(numbers_arr); i++ {

		if numbers_arr[i] > 9 {
			numbers_arr[i] = numbers_arr[i]/10 + numbers_arr[i]%10
		}

		sum += uint(numbers_arr[i])

	}

	if len(numbers_arr) < 10 || sum%10 != 0 {
		return false
	}

	return true

}
