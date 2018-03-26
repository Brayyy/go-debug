package main

import "time"
import "github.com/visionmedia/go-debug"

var debug = godebug.Debug("single")

func main() {
	for {
		debug("sending mail")
		debug("send email to %s", "tobi@segment.io")
		debug("send email to %s", "loki@segment.io")
		debug("send email to %s", "jane@segment.io")
		time.Sleep(500 * time.Millisecond)
	}
}
