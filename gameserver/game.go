package main

import (
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/middleware"
	"gameserver/server"
)

var (
	upgrader = websocket.Upgrader{}
)
var mkk = server.NewUm()

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

    GOB:
	for{
		dat := &server.UserMsg{}
		err := ws.ReadJSON(dat)
		if err != nil{
			println("sdaasdasd-->",err.Error())
			goto GOB
		}

		go server.WsConnect(ws,dat)
	}

}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./public") //创建服务
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
