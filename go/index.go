package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 2.01

	fmt.Printf("a:%4T %v \t b:%T %v \n", a, a, b, b) // 4T , 8T tab space
	fmt.Printf("a:%T %[1]v \t b:%T %[2]v \n", a, b)
	// var sum float64
	// var n int

	// for { // infinite loop
	// 	var val float64
	// 	_, err := fmt.Fscanln(os.Stdin, &val)
	// 	if err != nil {
	// 		break
	// 	}
	// 	sum += val
	// 	n++
	// }

	// if n == 0 {
	// 	fmt.Fprintln(os.Stderr, "No values")
	// 	os.Exit(-1)
	// }

	// //go run index.go < nums.txt
	// // cat nums.txt | go run index.go

	// fmt.Println("the average is: ", sum/float64(n))

	str := "Hello, 世界" // Contains English and Chinese characters

	// Iterate over runes in the string
	for _, r := range str {
		fmt.Printf("%v ", r) // Print each rune
	}
}
