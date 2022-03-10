package main

import (
	"fmt"
	//"strings"
)

// Methods with pointer receivers practice
type Cat struct{
		Name string
		Age uint8
		Color string
	}

func main() {

	// // Manipulating slices practice
	// EUcountries := []string{"Austria", "Belgium", "Bulgaria"}
	// addCountries(&EUcountries)
	// log.Println(EUcountries)

	// // Function/methods with a slice as parameter/receiver practice 
	// type Item struct{
	// 	SKU string
	// 	Quantity int
	// }
	// type Cart struct {
	// 	ID string
	// 	Paid bool
	// 	Items []Item
	// }
	// cart := Cart{
	// 	ID : "123413",
	// 	Paid: false,
	// }
	// cartPtr := &cart
	// (*cartPtr).Items=[]Item{
	// 	{SKU: "13243", Quantity:12},
	// 	{SKU: "989783", Quantity:1},
	// }
	// log.Println(cart.Items)

	// // Methods with pointer receivers practice
	cat := Cat{
		Name: "Bob",
		Age: 12,
		Color: "White",
	}
	cat.Meow()
	cat.Rename("Browe")
	fmt.Println(cat.Name)
	cat.Rename2("Browe")
	fmt.Println(cat.Name)

}

// // Manipulating slices practice
// func addCountries(EUcountriesPtr *[]string) {
// 	*EUcountriesPtr=append(*EUcountriesPtr, []string{"Crotia","Italy","Poland"}...)
// }

// // Methods with pointer receivers practice
func(c Cat)Meow(){

	fmt.Println("Meow")
}

func(c Cat)Rename(newName string){
	c.Name=newName
}

func(c *Cat)Rename2(newName string){
	c.Name=newName
}

