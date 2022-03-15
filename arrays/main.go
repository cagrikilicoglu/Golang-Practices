package main

import (
	"fmt"
	// "arrayPractices/demo"
)

func main(){

	// // standard definiton
	// var myArr [2]int
	// myArr[0]= 156
	// myArr[1]=147

	// // array literal
	// myArr := [2]int{156,147}
	// fmt.Println(myArr)

	// // when you create an empty array, Go will fill it with zero values of type (if not defined)

	// var myA [2]int
	// fmt.Println(myA)
	// var myB [3]string

	// len and cap of an array
	// fmt.Println(len(myB))
	// fmt.Println(cap(myA))

	// // accesing elements of an array

	// b:=[2]string{"FR", "US"}
	// firstElement := b[0]
	// secondElement := b[1]
	// fmt.Println(firstElement,secondElement)

	// // iterating over an array
	// myArr := [3]string{"monday","tuesday","wednesday"}
	// for index, value := range myArr {
	// 	fmt.Printf("element at index %d is %s\n", index,value)
	// }

	// // ignore the index
	// for _, v := range myArr{
	// 	fmt.Println(v)
	// }
	// for index, _ := range myArr{
	// 	fmt.Println(index)
	// }

	// // looping between certain index 
	// for i:=0; i<len(myArr); i++{
	// 		fmt.Printf("element in index %d is %s\n",i,myArr[i])
	// }

	// //iterate in descending order 
	// for i:= len(myArr)-1; i >=0 ; i-- {
	// 	fmt.Printf("element in index %d is %s\n",i,myArr[i])
	// }

	// // find a certain element in an array

	// for i, v := range myArr{
	// 	if v == "tuesday"{
	// 		fmt.Println("element ",v, "is at index ", i)
	// 	}
	// }

	// use function defined outside
	// getIndex(myArr, "wednesday")

	// // comparison of arrays
	// myA := [3]float32{0, 3.24, 198.23}
	// myB := [3]float32{0, 3.24, 198.23}

	// fmt.Println(myA == myB)

	// passing array to functions
	// myArray := [4]uint{12,148,31,20}

	// fmt.Println(increaseValues(&myArray))
	// fmt.Println(myArray)
	
	// myArr := [3]string{"travis", "scott", "astroworld"}
	// fmt.Println(demo.UpdateValue(myArr))

	// 3 dimensional arrays

	myArr := [2][2][2]string{}
	myArr[0][0][0] = "element 0,0,0"
	myArr[0][1][0] = "element 0,1,0"
	myArr[0][0][1] = "element 0,0,1"
	fmt.Println(myArr)


}

// func getIndex(arr [3]string, day string)int{
// 	for i, v:=range arr{
// 		if v == day {
// 			fmt.Println(i)
// 			return i
// 		} 
// 	}
// 	return -1
// }

// func increaseValues(arr *([4]uint)) [4]uint {
// 	for i,v := range *arr {
// 		arr[i] = v+1
// 	}
// 	return *arr
// } 