package main

import (
	"fmt"
)

func isPalindrome(n int) bool { //bool gives the ans in true or false, that the num is palindrome or not
	num := n
	rev_num := 0
	// if its palindrome then all steps
	for n != 0 {
		lastDigit := n % 10
		rev_num = (rev_num * 10) + lastDigit
		n /= 10
	}

	return num == rev_num
}

func main() {
	var num int
	fmt.Print("Enter a number: ") //if we use Print then input will take like this Enter a number: <= put value here
	fmt.Scanln(&num)              // if we use Println the it acts as print + \n(new line) its prints the line then take input

	if isPalindrome(num) {
		fmt.Println(num, "is a palindrome!")
	} else {
		fmt.Println(num, "is not a palindrome!")
	}
}
