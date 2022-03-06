// refactoring the code in controlStatements Hotel App in terms of functions

package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {

	const hotelName = "Gopher Paris Inn"
	const totalRooms = 134
	const maxAvailableRooms = 24

	// for simplicity room number of available rooms starts from 110 and increase by one as the number of available rooms
	const firstRoomNumber = 110

	// seeding random library with UNIX Epoch
	rand.Seed(time.Now().UTC().UnixNano())
	roomsAvailable := rand.Intn(maxAvailableRooms)
	occupancyRate := computeOccupancyRate(totalRooms,roomsAvailable)
	occupancyLevel := computeOccupancyLevel(occupancyRate)

	fmt.Println("Hotel: ", hotelName)
	fmt.Println("Number of Rooms : ", totalRooms)
	fmt.Println("Rooms Available : ", roomsAvailable)
	fmt.Println("Occupancy Level : ", occupancyLevel)
	fmt.Printf("Occupancy Rate : %0.2f %%\n", occupancyRate)
	

	if roomsAvailable == 0 {
		fmt.Println("No rooms available for tonight")
	} else {
	fmt.Println("Rooms :")
	for i:=0; i<roomsAvailable; i++ {
		roomNumber := firstRoomNumber+i
		availablePeople := rand.Intn(9) + 1
		availableNights := rand.Intn(9) + 1
		printRoomDetails(roomNumber,availablePeople,availableNights)
	}
}

}

func computeOccupancyRate(totalRooms, roomsAvailable int) float32 {

	roomsOccupied := totalRooms - roomsAvailable	
	return (float32(roomsOccupied)/float32(totalRooms)) * 100
}
func computeOccupancyLevel(occupancyRate float32) (occupancyLevel string) {

	switch {
	case occupancyRate < 0.3:
		occupancyLevel = "Low"
	case occupancyRate > 0.7:
		occupancyLevel = "High"
	default:
		occupancyLevel = "Medium"
	}
	return
}

func printRoomDetails(roomNumber, people, nights int) {

	nightText := " nights"
	peopleText := " people / "


	if people == 1 {
		peopleText = " person / "
	}
	if nights == 1 {
		nightText = " night"
	}
	
	fmt.Println(roomNumber, " : ", people, peopleText, nights, nightText )

} 