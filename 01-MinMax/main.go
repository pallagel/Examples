//Author Lalith Pallage
//Find Max and Min values in a slice

package main

import (
	"fmt"
	"sort"
)

func main() {
	si := []int{11, 56, 87, 9, 1, 127, 99}
	fmt.Println("Slice : ", si)

	//Print FindMaxLoop results
	fmt.Println(FindMaxLoop(si))

	//Print FindMax results
	fmt.Println(FindMax(si))
}

//Easy way to find max value
//FindMax - function to find max
func FindMax(si []int) int {
	length := len(si) - 1
	//sort the arry
	sort.Ints(si)

	//return the last element, which is the highest number
	return si[length]
}

//Loop through slice and compare items to find max
//FindMaxLoop
func FindMaxLoop(si []int) int {
	//get the array length
	length := len(si)

	//set the first element as max
	currentMax := si[0]

	//return currentMax if slice size is 1
	if length == 1 {
		return currentMax
	}

	//loop through to find the max value
	for i := 0; i < length; i++ {
		if currentMax < si[i] {
			currentMax = si[i]
		}
	}
	//return currentMax value
	return currentMax
}
