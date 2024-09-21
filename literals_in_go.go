package main

import "fmt"


var integerLiteral = 43 	// Integer Literal

var floatLiteral = 3.24 	// Floating-point Literal

var hexLiteral = 0x1F 		// Hexadecimal Literal (32 in decimal)

var octalLiteral = 0775 		// Octal Literal (493 in decimal)

var stringLiteral = " Hello, World!" // String Literal

var boolLiteralFalse = false 	// Boolean Literal

var boolLiteralTrue = true 		// Boolean Literal



ascii := 'a'		// Runes Literal
unicode := 'B' 		// Runes Literal
newline := '\n'		// Runes Literal

// when printed using the fmt package we hace the below format

func main()  {
	fmt.Printf("%d %[1]c %[1]q\n", ascii)		// "97 a 'a'"
	fmt.Printf("%d %[1]c %1[1]q\n", unicode)	// "68 D 'D'"
	fmt.Printf("%d %[1]q\n", newline) 			// "10 '\n'"
}