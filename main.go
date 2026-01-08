package main

import "github.com/AuraReaper/helix/cache"

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	server := NewServer(opts, cache.New())
}
