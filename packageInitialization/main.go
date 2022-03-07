package main

import (
	"fmt"
	"packageInitializationPractice/invoice"
)

func init() {
	fmt.Println(("main"))
}

func main() {
	fmt.Println("--program start--")
	invoice.Print()
}
