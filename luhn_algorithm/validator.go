package main

import (
	luhnalgorithm "credit-card-validator/luhn_algorithm"
)

func main() {
	// Valid number
	luhnalgorithm.LuhnAlgorithm(79927398713)

	// Invalid number (invalid lenght)
	luhnalgorithm.LuhnAlgorithm(123455)

	// Invalid number
	luhnalgorithm.LuhnAlgorithm(79994566312)
}
