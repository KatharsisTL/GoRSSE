package GoRSSE

import (
	"net/rpc"
	"net"
	"log"
	"net/http"
	"strconv"
	s "github.com/KatharsisTL/GoRSSE/SSE"
)

type Server struct {
	manager s.Manager
}

func StartServer(addrWithoutPort string, port int, settings []s.SSEServerSettings) {
	manager := s.NewManager(addrWithoutPort, settings)
	rpc.Register(&manager)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", addrWithoutPort+":"+strconv.Itoa(port))
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}