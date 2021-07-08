package main

import (
	"fmt"

	"github.com/guitarlnw/go-calculator/calculator"
)

func main() {
	fmt.Println("Pi : ", calculator.GetPI())
	data := []int{3, 4}
	fmt.Println("Sum Slice : ", calculator.SumFromSlice(data))
}
