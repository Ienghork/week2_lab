package main

import "fmt"

func main () {
	for{
		

		fmt.Println("  mini calculator:")
		fmt.Println("1) add")
		fmt.Println("2) subtract")
		fmt.Println("3) multuply")
		fmt.Println("4) devide")
		fmt.Println("5) modulus")
		fmt.Println("6) exit")

		var choice int
		fmt.Println("choose: ")
		fmt.Scan(&choice)


		if choice == 6 {
			fmt.Println("goodbye")
			break
		}

		if choice < 1 || choice > 5 {
			fmt.Print("error")
			continue
		}
		var a, b int
		fmt.Print(" please, input your first number:")

		fmt.Scan(&a)
		fmt.Println("please input your second number")

		fmt.Scan(&b)

		if choice == 1{
			fmt.Println("result:", a+b)
		}else if choice ==2 {
			fmt.Println("result:", a-b)
		}else if choice == 3{
			fmt.Println("result:", a*b)
		}else if choice == 4{
			if b == 0{
				fmt.Println("error: canot divide 0")
			}else{
				fmt.Println("result", a/b)
			}
		}else if choice == 5 {
			if b == 0 {
				fmt.Println("errror: cannot divide 0")
			}else {
				fmt.Println("result", a%b)
			}

		}
		

	}
}