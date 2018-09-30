package GoRSSE

import (
	"net/rpc"
	"log"
	"github.com/KatharsisTL/GoRSSE/SSE"
)

func SendMsg(addrWithPort string, appName string, msg string) {
	client, err := rpc.DialHTTP("tcp", addrWithPort)
	if err != nil {
		log.Println("dialing: ", err)
		return
	}
	request := SSE.Request{appName,msg}
	var reply int
	err = client.Call("Manager.SendMsg", request, &reply)
	if err != nil {
		log.Println("arith error: ", err)
		return
	}
	e := client.Close()
	if e != nil {
		log.Println(e.Error())
	}
	return
}
