package main 
import "fmt"
func main()  {
	var a, b int
	fmt.Print("input your number:")
	fmt.Scanln(&a)
	fmt.Print("input your number:")
	fmt.Scanln(&b)
	fmt.Println("both positive (&&):", a >0 && b> 0)
	fmt.Println("one greather than the other (||), a> b || b> a")
	fmt.Println("not equal(!):" , !(a==b))
}