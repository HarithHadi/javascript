package main

import "fmt"

func main() {
	counter := 18
	for i := 0; i <= 10; i++ {
		fmt.Printf("%v X %v = %v\n", counter, i, counter*i)
	}

}
