package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"znfz/server/protocol"
	"znfz/web_server/client"
)

// the orders struct
type Order struct {
	UsrAddress string
	Psw        string
	Desp       string
	Money      uint64
	JobAddress string
}

// save orders form Thirty-party api
func SaveOrder(c *gin.Context, cli *client.Client) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	order := &Order{}
	err = json.Unmarshal(body, order)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	req := &protocol.ReqSetContent{
		UserAddress: order.UsrAddress,
		PassWord:    order.Psw,
	}
	cli.Conn <- req
}

// get datas from the Thire-party apis and divide the money to
// eth accounts
func SaveBill(c *gin.Context, cli *client.Client) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	order := &Order{}
	err = json.Unmarshal(body, order)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	req := &protocol.ReqSetContent{
		UserAddress: order.UsrAddress,
		PassWord:    order.Psw,
	}
	cli.Conn <- req
}
