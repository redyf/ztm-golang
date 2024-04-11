//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favoriteColor string = "blue"
	fmt.Println("My favorite color is", favoriteColor)
	birthYear, ageInYears := 2003, 20
	fmt.Println("My birth year is:", birthYear, "and my age is:", ageInYears)
	var (
		firstInitial = "M"
		lastInitial  = "P"
	)
	var ageInDays int
	ageInDays = 20 * 365
	fmt.Println("My first name initial is:", firstInitial, "and the last one is:", lastInitial)
	fmt.Println("my age in days is:", ageInDays)
}
