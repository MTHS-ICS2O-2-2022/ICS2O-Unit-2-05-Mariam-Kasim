// Created by: Mariam Kasim
// Created on: March 2023
//
// This file contains the main function for the go_app application

package main

import (
	"fmt"
)

func main() {
	const rate float64 = 0.18
	var hours float64
	var pay float64
	var governmentpay float64
	var salary float64

	// input
	fmt.Println("Enter Hours:")
	fmt.Scanln(&hours)
	fmt.Print("Enter the salary:")
	fmt.Scanln(&salary)

	pay = (hours * salary) * (1 - rate)
	governmentpay = (hours * salary) * rate
	roundedpay := fmt.Sprintf("%.2f", pay)
	roundedgovernmentpay := fmt.Sprintf("%.2f", governmentpay)

	fmt.Println("Your pay is: ", roundedpay)
	fmt.Println("Your government pay is: ", roundedgovernmentpay)

	fmt.Println("\nDone.")
}
