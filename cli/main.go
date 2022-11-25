package main

import (
	"flag"
	"time"
)

func main() {
	f := flag.String("text", "hello world", "Test to write")
	flag.Parse()

	text := *f

	chars := []rune(text)
	out := ""

	for i := 0; i < len(chars); i++ {
		char := chars[i]
		if char == rune(' ') {
			println(out + string(char))
			out += string(char)
			continue
		}
		for j := 97; j <= 122; j++ {
			println(out + string(j))
			time.Sleep(50 * time.Millisecond)
			if char == rune(j) {
				out += string(char)
				break
			}
		}
	}
}
