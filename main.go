package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"iFei/config/baseCon"
	"iFei/controller"
	"iFei/forever"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	//logrus.SetLevel(logrus.DebugLevel)
	//iFcon := baseCon.InitIFcon()
	logrus.SetLevel(logrus.InfoLevel)

}

func main() {
	IFconBase := baseCon.LoadBaseConfig()
	//logrus.SetLevel(IFconBase.LogLevel)
	forever.MysqlRegister()

	gin.SetMode(gin.DebugMode)
	router := controller.RegisterRouterMap()
	//IFconBase :=
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", IFconBase.Host, IFconBase.Port),
		Handler: router,
	}
	ExitServerHandler(server)

	if err := server.ListenAndServe(); err != nil {
		//fmt.Println("serve listens failed: ", err)
		logrus.Errorf("serve listens failed: %v", err)
	}
}

func ExitServerHandler(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logrus.Infof("[signal %v]exiting IFei now...", s)
		if err := server.Close(); nil != err {
			logrus.Errorf("server close failed:%v", err)
		}

		//Todo: resources should be disconnected like databases
		forever.MysqlUnRegister()
		os.Exit(0)
	}()
}
