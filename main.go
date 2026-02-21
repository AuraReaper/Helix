package main

import (
	"flag"

	"github.com/AuraReaper/helix/cache"
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

	server := NewServer(opts, cache.New())
	server.Start()
}
