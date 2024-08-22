package main

func EvenOrOdd(num int) string {
	// Return "Even" if the number is even, "Odd" otherwise
	if num%2 == 0 {
		return "Even"
	}
	return "Odd"
}