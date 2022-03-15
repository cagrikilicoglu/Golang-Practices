package guestList

import (
	"github.com/Pallinder/go-randomdata"
	// "math/rand"
	// "time
)
func Generate() [10][2]string {
	guestList := [10][2]string{}
	for i,_:=range guestList{
		guestList[i]= roomGuests(i)
	}
	return guestList
}

func roomGuests(roomID int)[2]string{
	arr := [2]string{}
	// rand.Seed(time.Now().UTC().UnixNano())
	arr[0] = randomdata.SillyName()
	arr[1] = randomdata.SillyName()
	return arr
}