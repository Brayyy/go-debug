package main

import "time"
import "github.com/visionmedia/go-debug"

var a = godebug.Debug("multiple:a")
var b = godebug.Debug("multiple:b")
var c = godebug.Debug("multiple:c")

func work(debug godebug.DebugFunction, delay time.Duration) {
	for {
		debug("doing stuff")
		time.Sleep(delay)
	}
}

func main() {
	q := make(chan bool)

	go work(a, 1000*time.Millisecond)
	go work(b, 250*time.Millisecond)
	go work(c, 100*time.Millisecond)

	<-q
}
