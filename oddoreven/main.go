package main

import "fmt"

func main() {
	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range numSlice {
		if val%2 == 0 {
			fmt.Printf("%v is even\n", val)
		} else {
			fmt.Printf("%v is odd\n", val)
		}
	}
}
