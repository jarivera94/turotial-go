package main

import "fmt"

var nameUser string = "Jael Rivera Oviedo"

const ITEMS string = "constanst values "

func main() {
	var (
		age  int     = 10
		size float32 = 1.76
	)
	fmt.Println("hello world")
	fmt.Println("2 * 2 =", 2*2)
	fmt.Println(len("Hello world"))
	fmt.Println("Hello world"[1])
	fmt.Println("Hello world")
	fmt.Println((true && false) || (false && true) || !(false && false))
	var hello string = "Hello World from string. "
	fmt.Println(hello)
	//using large syntax
	var addition string = "In addition"
	//using short syntax
	ten := 14
	hello = hello + addition
	fmt.Println(hello)
	fmt.Println(ten)
	fmt.Println(nameUser)
	fmt.Println(age)
	fmt.Println(size)
	test()
	for_test()
}

func test() {
	fmt.Println("From test function", nameUser)
}

func for_test() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("best way to write for")
	for index := 0; index < 20; index++ {
		fmt.Println(index)
	}

	fmt.Println("sentence IF")

	for j := 0; j <= 20; j++ {
		if j%2 == 0 {
			fmt.Println(j, " is even")
		} else {
			fmt.Println(j, "is odd")
		}
		switch_test(j)
	}
}

func switch_test(number int) {

	switch number {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("The number is greater than 5")
	}
}
