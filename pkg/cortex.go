package pkg

import (
	"log"
	"os"
	"sshnot/internal"
	"sshnot/pkg/dispatcher/telegram"
	"sshnot/pkg/formatter"
)

// Cortex ties all the components together and performs the
// whole workload.
func Cortex(options *internal.Options) {
	if !isCorrectEvent() {
		return
	}

	scraper := NewScrape(options)
	formatted := formatter.Format(*scraper.Login)
	output, err := telegram.NewTelegramBot(options)
	if err != nil {
		log.Panic("Could not create telegram bot")
	}
	output.Send(formatted)
}

// isCorrectEvent checks whether the correct PAM event has happened
// for out notification script.
func isCorrectEvent() bool {
	event := os.Getenv("PAM_TYPE")
	// We are only interested in the "open_session" event. If we don't
	// distinct this, it's possible that messages are being send on
	// disconnects too.
	return len(event) > 1 && event == "open_session"
}
