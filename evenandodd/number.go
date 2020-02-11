package main

import "fmt"

type number []int

func newListNumbers() number {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func (n number) printType() {
	for _, value := range n {
		if value%2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}
}
