package main
import "fmt"
func main() {
	x, y := 0, 7
	f := func() int {
		x++
		y++
		return x
	}
	fmt.Println(f(), y, f())
}