package network

import (
	"fmt"
	"sync"
)

// for testing without tcp and other requirements
type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.RWMutex

	peers map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		lock:      sync.RWMutex{},
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr Transport) error {

	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport)

	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	// Lock(): only one go routine read/write at a time by acquiring the
	// lock. RLock(): multiple go routine can read(not write) at a time
	// by acquiring the lock.

	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		fmt.Println("ERROR")
		return fmt.Errorf("%s: could not send message to %s", t.addr, to)
	}

	peer.consumeCh <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
