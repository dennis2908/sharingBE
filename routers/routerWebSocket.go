package routers

import (
	"api_beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ws", &controllers.WebSocketController{}, "get:Get")
	beego.Router("/socket.io/", &controllers.WebSocketController{}, "get:GetIO")
	beego.Router("/socket.io/", &controllers.WebSocketController{}, "post:GetIO")
}
