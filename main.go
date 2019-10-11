package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"iFei/controller"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	log.SetLevel(log.DebugLevel)

}

func main() {
	router := controller.RegisterRouterMap()
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	ExitServerHandler(server)

	if err := server.ListenAndServe(); err != nil {
		//fmt.Println("serve listens failed: ", err)
		log.Errorf("serve listens failed: %v", err)
	}

}

func ExitServerHandler(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		//fmt.Printf("[signal %v]exiting pipe now...\n", s)
		log.Infof("[signal %v]exiting pipe now...", s)
		if err := server.Close(); nil != err {
			//fmt.Printf("server close failed:%v\n", err)
			log.Errorf("server close failed:%v", err)
		}

		//Todo: resources should be disconnected like databases

		os.Exit(0)
	}()
}
