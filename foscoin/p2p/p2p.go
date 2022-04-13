package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jon890/foscoin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// Port :3000 will upgrade the request from :4000
	openPort := r.URL.Query().Get("openPort")
	ip := utils.Splitter(r.RemoteAddr, ":", 0)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}

	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	initPeer(conn, ip, openPort)
}

func AddPeer(address, port, openPort string) {
	// from :4000 is requesting an upgrade from the port :3000
	urlStr := fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:])
	conn, _, err := websocket.DefaultDialer.Dial(urlStr, nil)
	utils.HandleErr(err)
	p := initPeer(conn, address, port)
	sendNewestBlock(p)
}
