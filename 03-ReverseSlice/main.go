//Authot - Lalith Pallage
//Rever a slice

package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Before Reversing :", si)

	reverse := ReverseSlice(si)
	fmt.Println("Reversed : ", reverse)

}

//ReverseSlice method to reverse the slice
func ReverseSlice(si []int) []int {
	//get the length
	length := len(si)

	//return slice if length is 1
	if length == 1 {
		return si
	}

	for i := 0; i < length/2; i++ {
		SwapValues(&si[i], &si[(length-1)-i])
	}

	return si
}

//SwpValues
func SwapValues(val1, val2 *int) {
	temp := *val1
	*val1 = *val2
	*val2 = temp
}
