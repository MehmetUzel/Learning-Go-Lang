package main

import "fmt"

func main() {
	fmt.Println("Welcome to Quiz game")

	// var name string = "Mehmet"

	// surname := "Uzel"

	// var age int = 23

	// var score float64 = 74.89

	// var isFinished bool = false

	//fmt.Printf("Hello %v %v, how r u",name,surname)

	// var name string

	// fmt.Scan(&name)

	// fmt.Println(name)

	var age int

	fmt.Printf("Enter\n your age: ")
	fmt.Scan(&age)

	if age >= 14 {
		fmt.Println("Yess you are old enough!")
	}else {
		fmt.Printf("Next yearrr")
		return
	}

	var answer string
	var answer2 string

	fmt.Printf("do you prefer 13 mini or 13 pro ?")
	fmt.Scan(&answer,&answer2)

	if answer + " " + answer2 == "13 mini" {
		fmt.Println("Bravo u are smart")
	} else if answer + " " + answer2 == "13 MINI"{
		fmt.Println("Bravo u are smart")
	} else {
		fmt.Printf("Are u sure that u need pro features ?")
	}


}
