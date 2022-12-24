package sse

import "log"

type Broker struct {
	/*
	   Events are pushed to this channel by the main events-gathering          routine
	*/
	Notifier chan []byte

	// new client connections
	NewClients chan chan []byte

	// closed client connections
	ClosingClients chan chan []byte

	// client connections registry
	Clients map[chan []byte]bool
}

func New() *Broker {
	return &Broker{
		Notifier:       make(chan []byte, 1),
		NewClients:     make(chan chan []byte),
		ClosingClients: make(chan chan []byte),
		Clients:        make(map[chan []byte]bool),
	}
}

func (b *Broker) Listen() {
	for {
		select {
		case s := <-b.NewClients:
			// A new Client has joined
			b.Clients[s] = true
			log.Printf("Client added. %d registered clients", len(b.Clients))
		case s := <-b.ClosingClients:
			// A client has dettached
			// remove them from our clients map
			delete(b.Clients, s)
			log.Printf("Removed Client. %d registered clients", len(b.Clients))
		case event := <-b.Notifier:
			// case for getting a new msg
			// thus send it to all Clients
			for clientMsgChan := range b.Clients {
				clientMsgChan <- event
			}
		}
	}
}
