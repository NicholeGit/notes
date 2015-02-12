package main

import (
	"log"

	"helper"
	"misc/packet"
	"stats/protos"
)

func HandleRequest(reader *packet.Packet) {
	defer helper.PrintPanicStack()
	b, err := reader.ReadS16()
	if err != nil {
		log.Println("read protocol error")
		return
	}

	handle := protos.ProtoHandler[b]
	//DEBUG("=== stats protocal====", b)
	if handle == nil {
		log.Println("service not bind", b)
		return
	}
	handle(reader)
}
