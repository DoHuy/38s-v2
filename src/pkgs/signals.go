package pkgs

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// HandleSigterm catches sigterm(15) or sigint(2) signal
func HandleSigterm(onClose func()) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	signalsReceived := 0
	go func() {
		for {
			select {
			case s := <-sig:
				signalsReceived++
				fmt.Println(s)
				if signalsReceived < 2 {
					fmt.Println("Waiting for server to finish before shutting down")
					onClose()
				} else {
					// Force exit when user hits Ctrl+C second time in a row
					fmt.Println("Force exit application!")
					os.Exit(1)
				}
			}
		}
	}()
}


