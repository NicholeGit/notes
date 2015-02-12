package main

import (
	"cfg"
	"helper"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			helper.ERR("caught panic in main()", x)
		}
	}()

	go func() {
		//为了检测性能与”net/http/pprof“配合
		helper.INFO(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	// start basic services
	startup()

	// Listen
	config := cfg.Get()
	service := ":8080"
	if config["service"] != "" {
		service = config["service"]
	}

	helper.INFO("Service:", service)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	helper.INFO("Game Server OK.")

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			helper.WARN("accept failed", err)
			continue
		}
		_ = conn
		//		go handleClient(conn)
	}

}

func checkError(err error) {
	if err != nil {
		helper.ERR("Fatal error:", err)
		os.Exit(-1)
	}
}
