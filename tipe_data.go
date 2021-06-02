package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -12345555
	var decimal = 3.14
	var axist bool = true
	var message string = "Halo"

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negative: %d\n", negativeNumber)
	fmt.Printf("bilangan desimal %f\n", decimal)
	fmt.Printf("bilangan desimal %.3f\n", decimal)
	fmt.Printf("bilangan boolean %t\n", axist)
	fmt.Printf("message %s\n", message)
}
