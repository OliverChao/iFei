package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"iFei/config/baseCon"
	"iFei/controller"
	"iFei/forever"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	//"github.com/valyala/fasthttp"
)

func init() {
	//logrus.SetLevel(logrus.DebugLevel)
	//iFcon := baseCon.InitIFcon()
	logrus.SetLevel(logrus.DebugLevel)

}

func main() {
	IFconBase := baseCon.LoadBaseConfig()
	//logrus.SetLevel(IFconBase.LogLevel)
	forever.MysqlRegister()
	forever.RedisRegister()
	forever.LoadMarkdownEngine()

	//gin.SetMode(gin.DebugMode)
	router := controller.RegisterRouterMap()
	//IFconBase :=
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", IFconBase.Host, IFconBase.Port),
		Handler: router,
	}
	ExitServerHandler(server)

	if err := server.ListenAndServe(); err != nil {
		logrus.Errorf("serve listens failed: %v", err)
	}
}

// yoooo~~~ close Server gracefully
func ExitServerHandler(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logrus.Infof("[signal %v]exiting IFei now...", s)
		if err := server.Close(); nil != err {
			logrus.Errorf("server close failed:%v", err)
		}

		// unregister here
		forever.MysqlUnRegister()
		forever.RedisUnRegister()

		os.Exit(0)
	}()
}
