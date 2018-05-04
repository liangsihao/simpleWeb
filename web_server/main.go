package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"znfz/web_server/api"
	"znfz/web_server/client"
)

// Reset handlefunc,change owners function to 'gin style' handle function by using golang Anonymous
// functions,this function have recover.
func Handler(cli *client.Client, f func(c *gin.Context, cli *client.Client)) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorln("panic", err)
			}
		}()
		f(c, cli)
	}
}

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	glog.Infoln("starting web service")
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// starts a new Grpc Client
	cli := client.NewClient("localhost:8089")
	go cli.Run()

	// Define my handlers
	r.POST("/saveorder", Handler(cli, api.SaveOrder)) // Save orders from Three_party api
	r.POST("/savebill", Handler(cli, api.SaveBill))   // Save bills from Three_party api

	// Handle all requests using net/http
	http.Handle("/", r)
	r.Run("localhost:8088")
}
