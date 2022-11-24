package broadcast

import (
	"fmt"
	"log"
	"time"
)

type Broadcaster struct {
	Registers   chan *Client
	Unregisters chan *Client
	Messages    chan string
	clients     map[*Client]struct{}
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		Registers:   make(chan *Client),
		Unregisters: make(chan *Client),
		Messages:    make(chan string),
		clients:     make(map[*Client]struct{}),
	}
}

func (b *Broadcaster) Run() {
	for {
		select {
		case client := <-b.Registers:
			// anyone has register
			b.clients[client] = struct{}{}
			log.Printf("%s registers, current online: %d", client.Name, len(b.clients))
		case client := <-b.Unregisters:
			// anyone has unregister
			delete(b.clients, client)
			close(client.Messages)
			log.Printf("%s unregisters, current online: %d", client.Name, len(b.clients))
		case msg := <-b.Messages:
			// any message from client
			for client, _ := range b.clients {
				client.Messages <- fmt.Sprintf("[%s] %s", time.Now().Format("2006/01/02 15:04:05"), msg)
			}
			log.Printf("has broadcasted message [%s] which from client", msg)
		}
	}
}
