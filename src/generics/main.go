package generics

import "fmt"

func Test() {
	slice := []int{5, 1, 2, 4}
	fmt.Println(reverse(slice))
}

type contraintReverse interface {
	int | string
}

func reverse[T contraintReverse](slice []T) []T {
	newSlice := make([]T, len(slice))

	initialPosition := len(newSlice) - 1

	for i := 0; i < len(slice); i++ {
		newSlice[initialPosition] = slice[i]
		initialPosition--
	}

	return newSlice
}
