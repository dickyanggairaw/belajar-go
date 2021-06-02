package main

import "fmt"

func main() {
	var angka = [...]int{5, 2, 7, 8}

	for i := 0; i < len(angka); i++ {
		for j := 0; j < len(angka); j++ {
			if j < len(angka)-1 && angka[j] < angka[j+1] {
				// fmt.Println(angka[j])
				var temp = angka[j]

				angka[j] = angka[j+1]
				angka[j+1] = temp
			}
		}
	}

	fmt.Println(angka)
}
