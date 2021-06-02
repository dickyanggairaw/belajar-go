package main

import (
	"fmt"
)

// func main() {
// 	var names = []string{"John", "Wick"}
// 	printMessage("halo", names)
// }

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")

// 	fmt.Println(message, nameString)
// }

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	var randomValue int

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random Number", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random Number", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random Number", randomValue)
// }

// func randomWithRange(min, max int) int {
// 	var value = rand.Int()%(max-min+1) + min
// 	return value
// }

func main() {
	divideNumber(5, 2)
	divideNumber(10, 1)
	divideNumber(-10, 0)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid devider %d cannot devide by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
