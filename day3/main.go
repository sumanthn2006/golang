package main
import "fmt"
func main(){
	
x, y := 10, 3
fmt.Println(x + y) // 13
fmt.Println(x - y) // 7
fmt.Println(x * y) // 30
fmt.Println(x / y) // 3
fmt.Println(x % y) // 1
fmt.Println(x >y)  // true
    a, b := 5, 10
               fmt.Println(a < b && b > 5) // true
               fmt.Println(a > b || b > 5) // true
               fmt.Println(!(a == b))      // true
}