package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scanf("%f", &num)

	truncated := math.Trunc(num)
	fmt.Println(truncated)
}
