package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sun-fight/zinx-websocket/global"
	"github.com/sun-fight/zinx-websocket/znet"
)

func main() {
	server := znet.NewServer()

	global.InitGormReadMysql()
	global.InitGormWriteMysql()
	global.InitRedis()
	fmt.Println(global.MysqlRead)
	fmt.Println(global.MysqlWrite)
	fmt.Println(global.Redis)

	bindAddress := fmt.Sprintf("%s:%d", global.GlobalObject.Host, global.GlobalObject.TCPPort)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/", server.Serve)
	router.Run(bindAddress)
}