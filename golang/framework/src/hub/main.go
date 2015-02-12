package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"

	"cfg"
	_ "db"
)

func init() {
	//	runtime.GOMAXPROCS(runtime.NumCPU())
}

const (
	DEFAULT_SERVICE = ":8889"
)

func main() {
	//收集错误信息
	defer func() {
		if x := recover(); x != nil {
			log.Println("caught panic in main()", x)
		}
	}()

	// start logger
	config := cfg.Get()
	if config["hub_log"] != "" {
		cfg.StartLogger(config["hub_log"])
	}

	//	aa := NextVal("USERID_GEN")
	//	log.Println("NextVal", aa)

	log.Println("Starting HUB")
	// signal handler
	go SignalProc()

	// sys routine
	go SysRoutine()

	// Listen
	service := DEFAULT_SERVICE
	if config["hub_service"] != "" {
		service = config["hub_service"]
	}

	log.Println("Hub Service:", service)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	log.Println("HUB Server OK.")
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		conn.SetLinger(-1)
		go handleClient(conn)
	}

}

//----------------------------------------------- handle hub request
func handleClient(conn net.Conn) {
	defer conn.Close()
	header := make([]byte, 2)
	ch := make(chan []byte, 100000)

	go StartAgent(ch, conn)

	//读取消息
	for {
		// header
		n, err := io.ReadFull(conn, header)
		if err != nil {
			log.Println("error receiving header, bytes:", n, "reason:", err)
			break
		}

		// data
		size := binary.BigEndian.Uint16(header)
		data := make([]byte, size)
		n, err = io.ReadFull(conn, data)
		if err != nil {
			log.Println("error receiving msg, bytes:", n, "reason:", err)
			break
		}
		ch <- data
	}

	close(ch)
}

func checkError(err error) {
	if err != nil {
		log.Printf("Fatal error: %v", err)
		os.Exit(-1)
	}
}
