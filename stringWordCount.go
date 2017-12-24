package main
import
(
	"fmt"
	"strings"
)
func WordCount (s string) map[string]int {
	wordsNcounts := make(map[string]int)
	words := strings.Fields(s)
	for i := range words {
		_, present := wordsNcounts[words[i]]
		if present {
			wordsNcounts[words[i]]++
		} else {
			wordsNcounts[words[i]] = 1
		}
	}
	return wordsNcounts
}
func main () {
	fmt.Println(WordCount("test repeat test check"))
}
