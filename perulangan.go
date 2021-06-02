package main

import "fmt"

func main() {
	var kata string = "kamu dan aku bukan apa apa"
	for i := 0; i < len(kata); i++ {
		fmt.Println("perulangan ke ", i)
	}
}
