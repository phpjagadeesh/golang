package main

import "fmt"

func main() {
	// For loop one method
	for i := 1; i <= 10; i++ {
		fmt.Println("Value is", i)
	}

	var j = 1
	// For loop another method, same as while loop
	for j <= 10 {
		fmt.Println(j)
		j++
	}

}
