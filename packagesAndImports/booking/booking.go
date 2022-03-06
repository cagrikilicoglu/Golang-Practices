package booking

import "fmt"

const vatRate = 20.0

func printVatRate(){
	fmt.Printf("%.2f %%", vatRate)
}