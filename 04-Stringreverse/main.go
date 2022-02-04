//Author - Lalith Pallage
//String reverse

package main

import "fmt"

func main() {
	var userInput string

	fmt.Print("Please enter a string : ")
	fmt.Scanln(&userInput)

	fmt.Println("User Entered :", userInput)

	reverse := ReverseString(userInput)
	fmt.Println("In reverse order :", reverse)

}

//ReverseString - function to reverse string
func ReverseString(input string) (reverse string) {
	for _, v := range input {
		reverse = string(v) + reverse
	}
	return reverse
}
