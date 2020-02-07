package main

import "fmt"

func main() {

	fmt.Println("function")
	xs := []float64{2, 3, 4, 5, 6, 8, 9, 10}
	fmt.Println("The average of xs is ", averange(xs))

	//Returning multiple parameters
	a, b, c := multipleparams()
	fmt.Println(a, b, c)
	//variadic function

	fmt.Println(variadicFunction(1, 2, 3, 4, 5))
	variadic := []int{1, 2, 3, 4, 5}
	fmt.Println(variadicFunction(variadic...))

	//Clousure scopes

	x := 0

	increment := func() int {
		x++
		return x
	}
	fmt.Println("Print increment from 1 to 4")
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println("Print evens from 0 to 8")
	nextEven := makeGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("Factorial a number 10")
	fmt.Println(factorial(10))
	fmt.Println("Factorial a number 5")
	fmt.Println(factorial(5))

	//Defer
	defer third()
	defer second()
	first()

	//panic("PANIC")

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC HADLER")
}

func averange(xs []float64) float64 {
	total := 0.0

	for _, value := range xs {
		total += value
	}

	return total / float64(len(xs))
}

func multipleparams() (int, int, int) {
	return 1, 2, 3
}

func variadicFunction(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func makeGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}
func third() {
	fmt.Println("Third")
}
