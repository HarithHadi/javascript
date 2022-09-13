package main

import "fmt"

func main() {
	fmt.Println("Datatypes")
	// declare a variable in go
	// var variableName datatype
	var score int
	var cityName string
	var check bool
	var lifeline float32
	score = 88
	cityName = "KL"
	check = true
	lifeline = 88.88

	fmt.Println("Score : ", score) //you cant use + for (String & int)
	fmt.Println("City Name : ", cityName)
	fmt.Println("Check : ", check)
	fmt.Println("Lifeline : ", lifeline)
	fmt.Println("My score is ", score, ". I live in ", cityName)

	// use format specifier
	fmt.Printf("My score is %v. I live in %v\n", score, cityName)
	fmt.Printf("My lifeline is %v\n", lifeline)

	// %V format specifier -> get the datatype of variable
	fmt.Printf("Datatype of lifeline variable is %V \n", lifeline)

	// autotype inference
	// let go detect the datatype of the variable automatically
	// use :=
	autoscore := "pass"
	fmt.Printf("My autoScore is %v. \n", autoscore)
	fmt.Printf("Datatype of autoScore")

}
