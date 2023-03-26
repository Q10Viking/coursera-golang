package scanner

import "fmt"

func RunScanner() {
	var appleNum int
	fmt.Println("Number of Apples?")
	fmt.Scan(&appleNum)
	fmt.Println("Applies is ", appleNum)
}
