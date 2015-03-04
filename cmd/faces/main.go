package main

import (
	"fmt"
	"github.com/mattn/cool-ascii-faces"
)

func main() {
	for c := range faces.Stream() {
		fmt.Println(c)
	}
}
