package main 
import "fmt"
func main() {
	var a, b int
	fmt.Print("input your number1:")
	fmt.Scanln(&a)
	fmt.Print("input your number2:")
	fmt.Scanln(&b)
	
	
	result := a
	fmt.Println("=:", result)
	result += b
	fmt.Println("+=:", result)
	result -=b
	fmt.Println("-=:", result)
	result *= b
	fmt.Println("*=:", result)
	result /= b
	fmt.Println("/=:", result)
	result %= b
	fmt.Println("%=:",result)
}