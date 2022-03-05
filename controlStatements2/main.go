package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// random seeding
	rand.Seed(time.Now().UTC().UnixNano())

	var ageJohn, agePaul int = rand.Intn(110), rand.Intn(110)

	fmt.Println("John is ",ageJohn, " years old.")
	fmt.Println("Paul is ",agePaul, " years old.")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else if agePaul == ageJohn {
		fmt.Println("Paul and John have the same age")	
	} else {
		fmt.Println("Paul is younger than John")
	}

	// switch expression

	switch ageJohn {
	case 10: 
	fmt.Println("John is 10 years old")
	case 20:
	fmt.Println("John is 20 years old")
	case 100:
	fmt.Println("John is 100 years old")
	default:
	fmt.Println("John is neither 10,20 nor 100 years old")
	} 

	// switch statement; expression

	switch sumAge := ageJohn + agePaul; sumAge {
	case 10:
	fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40: 
	fmt.Println("ageJohn + agePaul = 20, 30 or 40")
	case 2 * agePaul:
	fmt.Println("ageJohn + agePaul = 2 times agePaul")
	}

	// switch (without statement, without expression)
	switch {
	case agePaul > ageJohn:
	fmt.Println("agePaul > ageJohn")
	case agePaul == ageJohn:
	fmt.Println("agePaul == ageJohn")
	case agePaul < ageJohn:	
	fmt.Println("agePaul > ageJohn")
	}



	
	
}