package main

import (
	"flag"
	"fmt"
	"github.com/mattn/cool-ascii-faces"
)

var n = flag.Int("n", 0, "N faces")

func main() {
	flag.Parse()
	if *n == 0 {
		for c := range faces.Stream() {
			fmt.Println(c)
		}
	} else {
		for i := 0; i < *n; i++ {
			fmt.Println(faces.String())
		}
	}
}
