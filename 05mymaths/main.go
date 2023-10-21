package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
	//"time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var num1 int = 2
	// var num2 float64 = 4

	// fmt.Println("The sum is: ", num1 + int(num2))

	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//random number using cryto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
