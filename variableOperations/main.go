package main

import "fmt"

func main() {
	var hotelName string = "Gopher Hotel"
	const longitude, latitude = 24.806078, -78.243027
	var occupancy int = 12
	fmt.Println(hotelName, "\n",longitude, "\n",latitude, "\n",occupancy)

}