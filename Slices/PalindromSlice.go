//Lege ein Integer Slice an- 
//Prüfe ob es sich um ein Palindrom handelt.

package main

import "fmt"

func main(){
	// Legen ein Integer Slice an
	numbers := []int{1, 2, 3, 2, 1}

	// Prüfen ob es sich um ein Palindrom handelt
	palindrome := true
	for i := 0; i < len(numbers)/2; i++ {
		if numbers[i] != numbers[len(numbers)-i-1] {
			palindrome = false
			break
		}
	}

	// Ausgabe
	if palindrome {
		fmt.Println("Es handelt sich um ein Palindrom")
	} else {
		fmt.Println("Es handelt sich nicht um ein Palindrom")
	}
}