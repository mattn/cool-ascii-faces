package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/mattn/cool-ascii-faces"
)

func shot(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(faces.String())
	}
}

func stream(s time.Duration) {
	ctx, cancel := context.WithCancel(context.Background())

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		<-sc
		cancel()
	}()
	for c := range faces.Stream(ctx) {
		fmt.Println(c)
		time.Sleep(s)
	}
}

func main() {
	var n int
	var s time.Duration
	flag.IntVar(&n, "n", 0, "N faces")
	flag.DurationVar(&s, "s", 0, "Sleep")
	flag.Parse()
	if n == 0 {
		stream(s)
	} else {
		shot(n)
	}
}
