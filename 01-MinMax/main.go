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
	fmt.Println("Find Max Loop", FindMaxLoop(si))

	//Print FindMinLoop
	fmt.Println("Find Min Loop", FindMinLoop(si))

	//Print FindMax results
	fmt.Println("Find Max", FindMax(si))

	//Print FindMin results
	fmt.Println("Find Min :", FindMin(si))
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

//Find Min Value
func FindMinLoop(si []int) int {
	length := len(si)

	//set the first value as min
	minValue := si[0]

	//if its the first element, return it
	if length == 1 {
		return minValue
	}

	//Loop through to find min value
	for i := 0; i < length; i++ {
		if minValue > si[i] {
			minValue = si[i]
		}
	}

	//return min value
	return minValue
}

//Easy way
//FindMin function
func FindMin(si []int) int {
	//sort the slice
	sort.Ints(si)

	return si[0]
}
