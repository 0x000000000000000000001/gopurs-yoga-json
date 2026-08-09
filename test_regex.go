package main
import (
	"fmt"
	"regexp"
)
func main() {
	s := `[\u0000-\u002f\u003a-\u0040\u005b-\u0060\u007b-\u007f]+`
	re1 := regexp.MustCompile(`\\u([0-9a-fA-F]{4})`)
	re2 := regexp.MustCompile(`\\u\{([0-9a-fA-F]+)\}`)
	s = re1.ReplaceAllString(s, `\x{${1}}`)
	s = re2.ReplaceAllString(s, `\x{${1}}`)
	fmt.Println(s)
	_, err := regexp.Compile(s)
	fmt.Println("Err:", err)
}
