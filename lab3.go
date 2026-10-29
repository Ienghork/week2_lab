package main 

import "fmt"

func main () {
	a := 4
	b := 6

	fmt.Println("XOR:", a^b)

	fmt.Println("not a:", ^a)
	fmt.Println("not b:", ^b)

	fmt.Println("OR:", a|b)


	fmt.Println("and:", a&b)


	fmt.Println("left shift a:", a<<1)
	
	
	
	fmt.Println("right shift a:", a >>1)


}