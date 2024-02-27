package main

import (
	"time"

	"github.com/NikhilSharmaWe/go-blockchain/network"
)

// Server
// Transport => tcp, udp
// block
func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello"))
			time.Sleep(time.Second)
		}
	}()
	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal, trRemote},
	}

	s := network.NewServer(opts)

	s.Start()
}
