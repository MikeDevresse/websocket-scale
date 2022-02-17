package websocket

type server struct {
	clients map[string]*client
}

func NewServer() *server {
	return &server{
		clients: make(map[string]*client),
	}
}

func (server *server) AddClient(client *client) {
	server.clients[client.identifier] = client
}
