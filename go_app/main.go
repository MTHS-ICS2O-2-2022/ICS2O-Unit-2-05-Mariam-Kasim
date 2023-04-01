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
	var governmentPay float64
	var salary float64
	
	// input
	fmt.Print("Enter Hours: ")
	fmt.Scanln(&hours)
	fmt.Print("Enter the salary: ")
	fmt.Scanln(&salary)

	// process
	pay = (hours * salary) * (1 - rate)
	governmentPay = (hours * salary) * rate
	roundedpay := fmt.Sprintf("%.2f", pay)
	roundedGovernmentPay := fmt.Sprintf("%.2f", governmentPay)

	// output
	fmt.Println("")
	fmt.Println("Your pay is: $", roundedpay)
	fmt.Println("Your government pay is: $", roundedGovernmentPay)

	fmt.Println("\nDone.")
}
