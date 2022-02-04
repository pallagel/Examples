//Author - Lalith Pallage
//Check if a string is a Palindrome

package main

import "fmt"

func main() {
	var input string

	//Accept user input
	fmt.Print("Enter a string : ")
	fmt.Scanln(&input)

	fmt.Println(IsPalindrome(input))
}

//IsPalindrome - function to test a string is a Palindrome
func IsPalindrome(input string) bool {
	j := len(input) - 1 //sset the variable j to length of the string

	//loop through input mid-length and check characters at begging and end to see if they are the same
	for i := 0; i < j/2; i++ {
		if (input[i]) != input[j-i] {
			return false
		}
	}
	return true
}
