package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mattn/cool-ascii-faces"
)

var (
	n = flag.Int("n", 0, "N faces")
	s = flag.Duration("s", 0, "Sleep")
)

func main() {
	flag.Parse()
	if *n == 0 {
		for c := range faces.Stream() {
			fmt.Println(c)
			time.Sleep(*s)
		}
	} else {
		for i := 0; i < *n; i++ {
			fmt.Println(faces.String())
		}
	}
}
