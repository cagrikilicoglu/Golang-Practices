package room

import "fmt"

// display information about a room

func PrintDetails (roomNumber, size, nights int ) {
	fmt.Println(roomNumber, " : ", size, " people / ", nights, " nights ")
}