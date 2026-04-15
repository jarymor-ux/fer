package hub

import (
	"github.com/jarymor-ux/fer/internal/io/ws/client"
)

type Hub struct {
	clients    map[*client.Client]bool //nolint:unused
	broadcast  chan []byte             //nolint:unused
	register   chan *client.Client     //nolint:unused
	unregister chan *client.Client     //nolint:unused
}

func NewHub() *Hub {
	return new(Hub) //TODO:переписать
}
