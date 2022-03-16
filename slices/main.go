package main

import (
	"fmt"
	// "strings"
)

// // slice creation 2
// var s2 = []int{12,24}
func main(){

	// // slice creation
	// s:= make([]int,3)
	// s[0]= 10
	// s[2]= 20
	// fmt.Println(s,s2)

	// // slicing an array
	// customers :=[4]string{"Kanye West", "Lil Wayne", "Travis Scott", "Dr. Dre"}
	// customersSlice:=customers[0:1]
	// customersSlice2:=customers[2:4]
	// fmt.Println(customersSlice,customersSlice2)

	// Do slicing copy data?
	// customers :=[4]string{"Kanye West", "Lil Wayne", "Travis Scott", "Dr. Dre"}
	// customersSlice:=customers[0:1]
	// fmt.Println(customersSlice)
	// customers[0] = "Drake"
	// fmt.Println("After modification of original array")
	// fmt.Println(customersSlice)

	// // Slicing strings
	// hotelName:= "Go Dev Hotel"
	// s:= hotelName[0:6]
	// fmt.Println(s, len(s))
	// hotelName = "Java Dev Hotel"
	// fmt.Println(s)

	// // Length of a slice

	// randomSlice := []float64{31.213, 9.21, 6, 190.7}
	// fmt.Println(len(randomSlice))

	// // var randomSlice2 []float64
	// randomSlice2 := make([]float64,2)
	// randomSlice2 = []float64{12, 30, 50.6}
	// fmt.Println(randomSlice2)

	// // Capacity of a slice 
	// myArr := [4]float32{24, 40.2, 42, 11}
	// mySlice := myArr[1:3]
	// fmt.Println("len: ",len(mySlice))
	// fmt.Println("cap: ",cap(mySlice))

	// Slices as function paramters
	// s:= []int{1,2,3}
	// fmt.Println(s)
	// multiplyValues(s,2)
	// fmt.Println(s)

	// // Slice capacity common errors
	// languages := []string{"C", "Java", "Pyhton","C#"}

	// fmt.Println("Capacity:", cap(languages))
	// addGo(&languages)
	// fmt.Println("Capacity:", cap(languages))
	// fmt.Println(languages)
	// languages := [4]string{"C", "Java", "Pyhton","C#"}
	// languagesSlice := languages[1:3]
	// fmt.Println("Capacity:", cap(languagesSlice), len(languagesSlice))
	// addGo(languagesSlice)
	// fmt.Println(languagesSlice)
	// fmt.Println("Capacity:", cap(languagesSlice),len(languagesSlice))

	// builtin copy function

	// sourceSlice:=[]int{0,12,34,66,221}
	// destinationSlice := []int{1,1,1,1}
	// copiedNumber := copy(destinationSlice,sourceSlice)
	// fmt.Println(sourceSlice,destinationSlice,copiedNumber)


	// sourceSlice=[]int{0,12,34}
	// destinationSlice = []int{1,1,1,1}
	// copiedNumber = copy(destinationSlice,sourceSlice)
	// fmt.Println(sourceSlice,destinationSlice,copiedNumber)

	// sourceSlice=[]int{0,12,34,66,221}
	// destinationSlice = make([]int,5)
	// copiedNumber = copy(destinationSlice,sourceSlice)
	// fmt.Println(sourceSlice,destinationSlice,copiedNumber)

	// // append builtin function

	// a:= make([]int, 5,8)
	// // a:=[]int{1,2,3}
	// fmt.Println(cap(a))
	// a=append(a,10,20,30)
	// fmt.Println(a, cap(a))

	// // append builtin function slice growing example

	// s:=[]uint{10,20,30,40}
	// fmt.Printf(s)
	// fmt.Printf("Lenght: %d, capacity: %d\n", len(s),cap(s))
	// s = append(s,50)
	// fmt.Println(s)
	// fmt.Printf("Lenght: %d, capacity: %d\n", len(s),cap(s))
	// s = append(s,60)
	// fmt.Println(s)
	// fmt.Printf("Lenght: %d, capacity: %d\n", len(s),cap(s))
	// s = append(s,70)
	// fmt.Println(s)
	// fmt.Printf("Lenght: %d, capacity: %d\n", len(s),cap(s))
	// s = append(s,80)
	// fmt.Println(s)
	// fmt.Printf("Lenght: %d, capacity: %d\n", len(s),cap(s))
	// s = append(s,90)
	// fmt.Println(s)
	// fmt.Printf("Lenght: %d, capacity: %d\n", len(s),cap(s))
	// s = append(s,100)
	// fmt.Println(s)
	// fmt.Printf("Lenght: %d, capacity: %d\n", len(s),cap(s))
	// s = append(s,110)
	// fmt.Println(s)
	// fmt.Printf("Lenght: %d, capacity: %d\n", len(s),cap(s))
	
	// Performance impact 
	// the second one is way more efficient

	// s:=[]uint{10,20,30,40}
	// s=append(s,50)
	// s=append(s,60)
	// s=append(s,70)
	// s=append(s,80)
	// s=append(s,90)

	// fmt.Println(s)

	// s:= make([]int,0,10)
	// s= append(s,10,20,30,40)
	// s=append(s,50)
	// s=append(s,60)
	// s=append(s,70)
	// s=append(s,80)
	// s=append(s,90)

	// a:=[3]int{1,2,3}

	// fmt.Println(s)

	// elementAtIndex := a[5]

	// fmt.Println(elementAtIndex)

	// // for loop over slices

	// names := []string{"Bob","Nick", "Travis"}

	// // below will not modify original slice but makes a copy of name and modifies it
	// for _,name := range names {
	// 	name=strings.ToUpper(name)
	// }
	// fmt.Println(names)

	// // below will modify original slice

	// for i := range names{
	// 	names[i]=strings.ToUpper(names[i])
	// }
	// fmt.Println(names)

	// // merging slices

	// slice1:=[]string{"red","blue","green"}
	// slice2:=[]string{"purple", "white","black","gray"}

	// slice1=append(slice1,slice2...)
	// fmt.Println(slice1,slice2)

	// // removing an element from a slice

	// slice1:=[]int{1,2,3,4,5,6,7,8,9,10}
	// // remove 5th element
	// slice1 = append(slice1[:5],slice1[6:]...)
	// fmt.Println(slice1)

	// // adding an element to the slice at a specific index
	
	// letters:=[]string{"A","B","D","E","F","G"}

	// letters = append(letters,"")
	// copy(letters[3:],letters[2:])
	// letters[2]="C"
	// fmt.Println(letters)

	// b:=[]string{"john","jeanne","jean","josh"}
	// b=append(b,"")
	// fmt.Println(b)
	// copy(b[2:],b[1:])
	// fmt.Println(b)
	// b[1]="joe"
	// fmt.Println(b)

	// // remove all elements of a slice
	// a:= []int{0,1,2,3,4,5,6,7,8,9}
	// a=a[:0]
	// fmt.Println(a, len(a), cap(a))
	// a:=[]string{"a","b","c","d","e"}
	// a=nil
	// fmt.Println(a,len(a),cap(a))

	// // Two dimensional slices

	// innerSlice1 :=[]float64{2.30,4.01,6.99,8}
	// innerSlice2 := []float64{1,10.2,1.34,23}

	// my2DSlice := [][]float64{innerSlice1,innerSlice2}
	// fmt.Println(my2DSlice)

	// // lighter version of previous script
	// my2DSlice := [][]float64{[]float64{2.30, 4.01, 6.99, 8}, []float64{1,10.2,30,1.34,23}}
	// fmt.Println(my2DSlice)

	// // creation of two dimensional slices dynamically
	// myOther2DSlice:=[][]int{}
	// for i:=0; i<10;i++{
	// 	myOther2DSlice = append(myOther2DSlice,[]int{i})
	// }
	// fmt.Println(myOther2DSlice)

	// creation of three dimensional slices 

	mySlice1 := [][]float64{[]float64{1.01,2.02,3.03}, []float64{2.45,2.12,3.33}}
	mySlice2 := [][]float64{[]float64{10.23, 332, 12.2}, []float64{2.66,9.09}}
	my3DSlice := [][][]float64{mySlice2,mySlice1}
	fmt.Println(my3DSlice)
}

// func multiplyValues(s []int,factor int) []int{
// 	for i,_ :=range s{
// 		s[i] *= factor
// 	}
// 	return s
// }

// func addGo(s []string) {
// 	s = append(s, "Go")
// }