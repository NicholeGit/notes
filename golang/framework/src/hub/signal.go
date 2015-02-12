package main

import (
	"cfg"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//----------------------------------------------- handle unix signals
func SignalProc() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)

	for {
		msg := <-ch
		log.Println("Recevied signal:", msg)
		cfg.Reload()
	}
}
