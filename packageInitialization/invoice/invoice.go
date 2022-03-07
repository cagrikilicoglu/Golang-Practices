package invoice

import (
	"fmt"
	"packageInitializationPractice/currency"
)

var c = func() string {
	fmt.Println("variable c initialized", b)
	return "value of c"
} ()

// var a will initialize first since it does not depend on any other variable
var a = func() string {
	fmt.Println("variable a initialized")
	return "value of a"
} ()

var b = func() string {
	fmt.Println("variable b initialized", a)
	return "value of b"
} ()

func init() {
	fmt.Println("invoice init", c)
}

func Print() {
	fmt.Println("Invoice Number X")
	fmt.Println(54, currency.EuroSymbol())
}