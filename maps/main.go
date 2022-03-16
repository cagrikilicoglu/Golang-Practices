package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)
type employee struct{
	name string
	genre string
	position string
}

func main(){
	file, err := os.Open("/Users/cagrikilicoglu/Desktop/Golang-Practices/maps/employees.csv")
	if err != nil {
		log.Fatalf("impossible to open file %s", err)
	}

	defer file.Close()

	// creating a reader to read the file line by line
	r:=csv.NewReader(file)

	employees:=make(map[string]employee)
	orderSlice := make([]string, 0, 10)

	// start reading with the for loop
	for {

		// record variable is a slice of strings
		record, err:= r.Read()
		// since Read will populate an EOF error at the end of the file, check if reader is at the end of the file and break for loop there
		if err == io.EOF {
			break
		} 
		// if reader has an error other than EOF, log this error
		if err != nil {
			log.Fatal(err)
		}
		employee := employee{
			name: record[1],
			genre: record[2],
			position: record[3],
		}
		employeeID := record[0]
		employees[employeeID] = employee
		orderSlice =append(orderSlice, employeeID)
	
		
	}


	value, ok :=employees["V45657"] 
	if ok{
		
		fmt.Println(value)
		
	} else {
		fmt.Println("There is no such employee")
	}

	for _,v := range orderSlice{
		fmt.Printf("Key %s, value: %v\n",v,employees[v])
	}

	// for _,value:= range employees{
	// 	// fmt.Printf("employee %s is a %s and working as a %s\n", value.name, value.genre, value.position )
	// 	fmt.Printf("%s\n", value)
	// }

	// fmt.Println(employees)
	// fmt.Println(len(employees))

	// delete(employees,"V45657")

	// fmt.Println(employees)
	// fmt.Println(len(employees))
}