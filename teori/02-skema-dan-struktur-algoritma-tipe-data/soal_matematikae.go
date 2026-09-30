package main
 
import "fmt"
 
func main() {
	var x, y int
	var hasil float64
 
	fmt.Scan(&x, &y)
	hasil = float64(5*x*x-2*x*y) + float64(y*y*y)/float64(x+1)
	fmt.Println(hasil)
}