package main
import "fmt"

func main (){
	var a,b int

	//membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar a dan b
	a,b=b,a

	//Output
	fmt.Println(a)
	fmt.Println(b)
	
}