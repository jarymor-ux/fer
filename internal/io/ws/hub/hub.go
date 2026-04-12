package ws

import (
	"github.com/jarymor-ux/fer/internal/io/ws/client"
)

type Hub struct {
	clients    map[*client.Client]bool
	broadcast  chan []byte
	register   chan *client.Client
	unregister chan *client.Client
}

func NewHub() *Hub {
	return new(Hub) //TODO:переписать
}
