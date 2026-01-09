package main

import (
	"log"
	"net"
	"time"

	"github.com/AuraReaper/helix/cache"
)

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	go func() {
		time.Sleep(2 * time.Second)
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
		}

		conn.Write([]byte("SET master yashkr 2500"))
	}()

	server := NewServer(opts, cache.New())
	server.Start()
}
