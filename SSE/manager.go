package SSE

import (
	"strconv"
)

type Manager struct {
	Servers map[string]Server
}

type Request struct {
	AppName string
	Msg string
}

type SSEServerSettings struct {
	AppName string
	Port int
}

func (manager *Manager) SendMsg(request *Request, response *int) error {
	b := manager.Servers[request.AppName]
	b.Send(request.Msg)
	response = nil
	return nil
}

func NewManager(addrWithoutPort string, settings []SSEServerSettings) (manager Manager) {
	manager = Manager{Servers:map[string]Server{}}
	for _, setting := range settings {
		server := *NewServer()
		manager.Servers[setting.AppName] = server
		mServer := manager.Servers[setting.AppName]
		go mServer.Listen(addrWithoutPort+":"+strconv.Itoa(setting.Port))
	}
	return
}