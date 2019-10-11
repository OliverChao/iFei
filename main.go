package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"iFei/controller"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router := controller.RegisterRouterMap()
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	ExitServerHandler(server)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("serve listens failed: ", err)
	}

}

func ExitServerHandler(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		fmt.Printf("[signal %v]exiting pipe now...\n", s)
		if err := server.Close(); nil != err {
			fmt.Printf("server close failed:%v\n", err)
		}

		//Todo: resources should be disconnected like databases

		os.Exit(0)
	}()
}
