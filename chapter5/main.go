package main

import "fmt"

func main() {
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0

	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

	y := make([]float64, 5, 10)
	fmt.Println(y)

	z := make(map[string]int)
	z["key"] = 10
	fmt.Println(z)

	number, ok := z["other"]
	fmt.Println(number, ok)
}
