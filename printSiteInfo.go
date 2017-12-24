package main
import (
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)
type Lang struct {
	Name string
	Year int
	URL string
}
func do(f func(Lang)) {
	input, error := os.Open("lang.json")
	if error != nil {
		log.Fatal(error)
	}
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		error := dec.Decode(&lang)
		if error != nil {
			if error == io.EOF {
				break
			}
			log.Fatal(error)
		}
		f(lang)
	}
}
func count(name, url string, c chan<- string) {
	start := time.Now()
	r, error := http.Get(url)
	if error != nil {
		c <- fmt.Sprintf("%s: %s\n", name, error)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, dt)
}
func main() {
	start := time.Now()
	c := make(chan string)
	n := 0
	do(func(lang Lang) {
		n++
		go count(lang.Name, lang.URL, c)
	})
	for i := 0; i < n; i++ {
		fmt.Print(<-c)
	}
	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}
