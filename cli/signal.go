package cli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// catchInterrupt exits gracefully upon receiving a SIGINT (^C)
func catchInterrupt(s *spinner.Spinner) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	<-sigs
	if s.Active() {
		s.FinalMSG = color.GreenString("👋 Bye!\n")
		s.Stop()
	}

	os.Exit(0)
}
