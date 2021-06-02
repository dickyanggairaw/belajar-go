package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 3, 5, 5, 8)
	var msg = fmt.Sprintf("Rata-rata %.2f \n", avg)

	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0

	for _, number := range numbers {
		total += number
	}

	var average = float64(total) / float64(len(numbers))
	return average
}
