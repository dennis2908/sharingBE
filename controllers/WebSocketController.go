package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	websockett "github.com/googollee/go-socket.io/engineio/transport/websocket"
	websocket "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: time.Duration(time.Second * 1),
}

// Easier to get running with CORS. Thanks for help @Vindexus and @erkie
var allowOriginFunc = func(r *http.Request) bool {
	return true
}

type WebSocketController struct {
	beego.Controller
}

func (api *WebSocketController) handle(conn *websocket.Conn) {
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			beego.Error("error", err)
			return
		}
		fmt.Printf("Received message: %s\n", msg)
		reply := string(msg)
		err = conn.WriteMessage(websocket.TextMessage, []byte(reply))
		if err != nil {
			beego.Error("error", err)
			return
		}
	}

}

func (api *WebSocketController) GetIO() {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websockett.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	server.OnConnect("", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		fmt.Println("chat:", msg)
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	defer server.Close()

}

func (api *WebSocketController) Get() {
	conn, err := upgrader.Upgrade(api.Ctx.ResponseWriter, api.Ctx.Request, nil)
	if err != nil {
		return
	}
	go api.handle(conn)
}
