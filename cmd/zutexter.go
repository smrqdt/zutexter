package main

import (
	"net/url"
	"os"

	"github.com/charmbracelet/log"

	"github.com/smrqdt/zutexter/pkg/display"
	powertexter "github.com/smrqdt/zutexter/pkg/powerTexter"
	"github.com/smrqdt/zutexter/pkg/texter"
)

func main() {
	displayURL, err := url.Parse(os.Getenv("DISPLAY_URL"))
	if err != nil {
		log.Fatalf("Could not parse URL: '%v' (%v)", os.Getenv("DISPLAY_URL"), err)
	}
	display := display.Display{URL: *displayURL}

	texters := []texter.Texter{
		&powertexter.PowerTexter{GrafanaURL: "https://c3power.top/api/ds/query"},
	}

	for _, texter := range texters {
		msg, err := texter.Text()
		if err != nil {
			log.Error("Could not generate message", "error", err, "texter", texter)
		}
		err = display.Queue(msg)
		if err != nil {
			log.Error("Could not quee message", "error", err, "message", msg)
		} else {
			log.Info("Sent message", "message", msg)
		}
	}
}
