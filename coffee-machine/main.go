package main

import (
	"fmt"
)

func main() {
	// write your code here
	var water, milk, beans, cups int
	fmt.Scan(&water)
	fmt.Scan(&milk)
	fmt.Scan(&beans)
	fmt.Scan(&cups)

	water /= 200
	milk /= 50
	beans /= 15

	n := findMin(water, milk, beans)

	if cups < n {
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)", n-cups)
	} else if n == cups {
		fmt.Println("Yes, I can make that amount of coffee")
	} else {
		fmt.Printf("No, I can make only %d cup(s) of coffee", n)
	}
}

func findMin(a ...int) int {
	min := a[0]
	for _, val := range a {
		if val < min {
			min = val
		}
	}
	return min
}
