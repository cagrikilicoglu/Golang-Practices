package main

import (
	"packagesAndImportsRefactoring/email"
	"packagesAndImportsRefactoring/invoice"
)

func main() {
	customerName := "Doe"
	customerEmail := "john.doe@example.com"
	var nights uint = 12

	emailContents := email.Contents("M", customerName, nights)
	email.Send(emailContents, customerEmail)
	invoice.Create(customerName, nights, 145.32)

}