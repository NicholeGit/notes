package main

import (
	"runtime"
	"time"
)

import (
	//	"agent/gsdb"
	"helper"
	"misc/timer"
)

const (
	SYS_MQ_SIZE = 65535 // size of sys routine's message queue
	GC_INTERVAL = 300   // voluntary GC interval
)

//---------------------------------------------------------- system routine
func SysRoutine() {
	//	var sess Session
	//	sess.MQ = make(chan IPCObject, SYS_MQ_SIZE)
	//	gsdb.RegisterOnline(&sess, SYS_USR)

	// timer
	gc_timer := make(chan int32, 10)
	gc_timer <- 1

	for {
		select {
		//		case msg, ok := <-sess.MQ: // IPCObject to system routine
		//			if !ok {
		//				return
		//			}
		//			IPCRequestProxy(&sess, &msg)
		case <-gc_timer:
			runtime.GC()
			helper.INFO("== PERFORMANCE LOG ==")
			helper.INFO("Goroutine Count:", runtime.NumGoroutine())
			helper.INFO("GC Summary:", helper.GCSummary())
			//			INFO("Sysroutine MQ size:", len(sess.MQ))
			timer.Add(0, time.Now().Unix()+GC_INTERVAL, gc_timer)
		}
	}
}
