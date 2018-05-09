package manager

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"time"
	"znfz/server/protocol"
	"znfz/web_server/client"
	"znfz/web_server/models"
)

// the orders struct
type Order struct {
	UsrAddress  string    // Pay Address
	Description string    // desciption of the orders
	Bill        string    // bills
	CreateTime  time.Time // Creating time of the order
	Company     string    // company name
	BranchShop  string    // branch shop name
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
	req := &protocol.ReqPay{
		UserAddress: order.UsrAddress,
	}
	models.SaveOrders(order.CreateTime, order.Company, order.BranchShop, order.Description)
	cli.C.Pay(context.Background(), req)
}
