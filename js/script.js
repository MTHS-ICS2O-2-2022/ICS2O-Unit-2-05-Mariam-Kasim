// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Mariam Kasim
// Created on: Mar 2023
// This file contains the JS functions for index.html

"use strict"
/**
 * This function calculates how much income tax you have to pay
 */

function myButtonClicked() {
  // input
  const hoursWorked = parseFloat(document.getElementById("hours-worked").value)
  const hourlyWage = parseFloat(document.getElementById("hourly-wage").value)

  // process
  const pay = hoursWorked * hourlyWage * (1 - 0.18)
  const governmentPay = hoursWorked * hourlyWage * 0.18

  // output
  document.getElementById("your-pay").innerHTML = "You get $ " + pay.toFixed(2)
  document.getElementById("tax-paid").innerHTML =
    "The goverment gets $" + governmentPay.toFixed(2)
}
