package main

import (
	"context"
	"fmt"
	"net"

	"github.com/xiaonanln/pktconn"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:14572")

	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			pc := pktconn.NewPacketConn(context.TODO(), conn)
			fmt.Printf("client connected: %s\n", pc.RemoteAddr())

			for pkt := range pc.Recv() {
				fmt.Printf("recv packet: %d\n", pkt.GetPayloadLen())
				pc.Send(pkt) // send packet back to the client
				pkt.Release()
			}

			fmt.Printf("client disconnected: %s", pc.RemoteAddr())
		}()
	}
}
