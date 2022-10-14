package main

import (
	"fmt"

	"net/http"

	"muyusvr/cfg"
	muyucfg "muyusvr/cfg"
	"muyusvr/muyulog"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//gin引擎
var engine *gin.Engine

func main() {
	fmt.Println("start muyu server")
	//初始化服务配置文件
	e, c := muyucfg.InitedSvr()
	if e != nil {

		muyulog.Log.Error("初始化配置文件失败", zap.String("err", e.Error()))
		return
	}
	fmt.Printf("linsten address %s:%d\n", c.ServerIni.Addr, c.ServerIni.Port)
	//初始化log
	muyucfg.InitedLog(c.ServerIni)
	//初始化数据库连接
	e = muyucfg.InitedDB(c.DBServer)
	if e != nil {

		muyulog.Log.Error("初始化数据库失败", zap.String("err", e.Error()))
		return
	}
	//初始化gin
	initGin(c)
}

func initGin(c *cfg.ServerConf) {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	//使用日志
	engine.Use(gin.Logger())
	//使用Panic处理方案
	engine.Use(gin.Recovery())

	engine.GET("/", index)

	engine.Run(fmt.Sprintf("%s:%d", c.ServerIni.Addr, c.ServerIni.Port))
}

func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": ""})
}
