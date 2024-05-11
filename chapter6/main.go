package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func f2() (r int) {
	return
}

func f() int {
	return 2
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	fmt.Println(f2())

	fmt.Println(add(1, 2, 3, 4))

	this_add := func(x, y int) int {
		return x + y
	}

	fmt.Println(this_add(1, 2))

	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	defer second() // called after the function completes
	first()

	y := 5
	zero(&y)
	fmt.Println("Y is", y)

	z := new(int)
	zero(z)
	fmt.Println(*z)
	test_recover()
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func test_recover() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func zero(x *int) {
	*x = 0
}
