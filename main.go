package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/AuraReaper/helix/cache"
	"github.com/AuraReaper/helix/client"
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
		client, err := client.New(":3000", client.Options{})
		if err != nil {
			log.Fatal(err)
		}

		client.Set(context.Background(), []byte("user"), []byte("yash"), 0)

		client.Close()
	}()

	server := NewServer(opts, cache.New())
	server.Start()
}

func SendCommand(c *client.Client) {
	_, err := c.Set(context.Background(), []byte("master"), []byte("yash"), 0)
	if err != nil {
		log.Fatal(err)
	}
}
