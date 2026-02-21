package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/AuraReaper/helix/cache"
	"github.com/AuraReaper/helix/client"
	"github.com/AuraReaper/helix/protocol"
)

func main() {
	var (
		listenAddr = flag.String("listenaddr", ":3000", "listern  address of server")
		leaderAddr = flag.String("leaderaddr", "", "listen address of the leader")
	)
	flag.Parse()

	opts := ServerOpts{
		ListenAddr: *listenAddr,
		IsLeader:   len(*leaderAddr) == 0,
		LeaderAddr: *leaderAddr,
	}

	go func() {
		time.Sleep(2 * time.Second)
		for i := 0; i < 10; i++ {
			SendCommand()
			time.Sleep(200 * time.Millisecond)
		}
	}()

	server := NewServer(opts, cache.New())
	server.Start()
}

func SendCommand() {
	cmd := &protocol.CommandSet{
		Key:   []byte("user"),
		Value: []byte("yash"),
	}

	client, err := client.New(":3000", client.Options{})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Set(context.Background(), []byte("user"), []byte("yash"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write(cmd.Bytes())
}
