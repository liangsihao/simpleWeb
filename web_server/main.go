package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"znfz/web_server/client"
	"znfz/web_server/config"
	"znfz/web_server/db"
	m "znfz/web_server/manager"
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
	flag.Parse()
	var path string
	flag.StringVar(&path, "config", "../conf/config.toml", "config path")
	config.ParseToml(path) // initing config
	glog.Infoln("starting web service")

	db.InitDb(config.Opts().MysqlStr)

	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// starts a new Grpc Client
	cli := client.NewClient("192.168.83.139:8888")
	go cli.Run()

	// Three_party handlers
	r.POST("/saveorder", Handler(cli, m.SaveOrder)) // Save orders from Three_party api
	r.POST("/savebill", Handler(cli, m.SaveOrder))  // Save bills from Three_party api

	// Web handlers
	r.POST("/setbindmsg", Handler(cli, m.Setbindmsg)) // Bind UserAddress to CompanyAddrss
	r.POST("/getbindmsg", Handler(cli, m.Getbindmsg)) // Get bind message

	r.POST("/setapplycompany", Handler(cli, m.Setapplycompany)) // Apply CompanyAddrss
	r.POST("/getapplycompany", Handler(cli, m.Getapplycompany)) // Get Apply CompanyAddrss

	r.POST("/getmoney", Handler(cli, m.SaveOrder)) // GetMoney Request

	// Handle all requests using net/http
	http.Handle("/", r)
	r.Run("localhost:8088")
}
