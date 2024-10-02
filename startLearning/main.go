package main

import (
	"fmt"
	"math"
)

func main() {
	programm2()
	fmt.Scanln()
}

func outliner() {
	fmt.Println()
	for i := 0; i < 20; i++ {
		fmt.Printf("][")
	}
	fmt.Println()
	fmt.Println()
}

func programm1() {
	fmt.Println("Hello! What is your name?")
	fmt.Printf("My name is - ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)
	outliner()
}

func programm2() {
	key := true
	for key {
		fmt.Println("hello, how i can help you?\n" +
			"(Choose number)\n" +
			"1) solve quadratic equation\n" +
			"2) is this simple number?\n" +
			"3) is this triangle possible?\n" +
			"4) exit")
		var choose byte
		fmt.Scanln(&choose)
		switch choose {
		case 1:
			{
				quadricEquation()
				outliner()
			}
		case 2:
			{
				simpleNumber()
				outliner()
			}
		case 3:
			{
				possibilityOfTriangle()
				outliner()
			}
		case 4:
			{
				key = false
			}
		default:
			{
				fmt.Println("i don't understand what you need")
			}
		}
	}

}

func quadricEquation() {
	var a, b, c int
	var x1, x2 float64

	fmt.Println("The formula is ax^2 + bx + c = 0\n" +
		"for example 5x^2 + 2x + 2 = 0")

	fmt.Printf("Enter the a : ")
	fmt.Scanln(&a)

	fmt.Printf("Enter the b : ")
	fmt.Scanln(&b)

	fmt.Printf("Enter the c : ")
	fmt.Scanln(&c)

	discriminant := math.Pow(float64(b), 2) - 4*float64(a)*float64(c)

	if discriminant > 0 {
		x1 = (-float64(b) + math.Sqrt(discriminant)) / (2 * float64(a))
		x2 = (-float64(b) - math.Sqrt(discriminant)) / (2 * float64(a))
	} else if discriminant == 0 {
		x1 = -float64(b) / (2 * float64(a))
		x2 = x1
	} else {
		fmt.Println("do nothing..." +
			" because d < 0")
		return
	}
	fmt.Printf("Your Solution x1 is %f, and x2 is %f \n", x1, x2)
}

func possibilityOfTriangle() {
	var a, b, c int
	fmt.Println("Enter the size of side the triangle (a, b, c)")

	fmt.Printf("Enter the a : ")
	fmt.Scanln(&a)

	fmt.Printf("Enter the b : ")
	fmt.Scanln(&b)

	fmt.Printf("Enter the c : ")
	fmt.Scanln(&c)

	if a < b+c && b < c+a && c < a+b {
		fmt.Println("The triangle is possible")
	} else {
		fmt.Println("The triangle is not possible")
	}
}

func simpleNumber() {
	var number uint
	fmt.Printf("Enter the positive number : ")
	fmt.Scanln(&number)
	isTrue := true
	if number > 2 {
		for i := 2; float64(i) < math.Sqrt(float64(number))+1; i++ {
			if number%uint(i) == 0 {
				isTrue = false
				break
			}
		}
	}

	if isTrue {
		fmt.Println("The number is simple")
	} else {
		fmt.Println("The number is not simple")
	}
}
