package main
import 
(
	"fmt"
	"time"
)
func fibonacci() func() int {
	first := 0
	second := 1
	sum := 0
	return func() int {
		sum = first + second
		first = second
		second = sum
		return sum
	}
}
func main () {
	start := time.Now()
	f := fibonacci()
	for i := 0; i < 20000; i++ {
		fmt. Println(f())
	}
	fmt.Println(time.Since(start))
}
