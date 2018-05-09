package manager

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"znfz/server/lib"
	"znfz/server/protocol"
	"znfz/web_server/client"
	"znfz/web_server/db"
	"znfz/web_server/models"
)

type ReqApplyCompany struct {
	CompanyName string
	BranchName  string
	UserAddress string
	PassWord    string
	Phone       string
}

type RespApplyCompany struct {
	Address string
}

// Set apply company
func Setapplycompany(c *gin.Context, cli *client.Client) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	req := &ReqApplyCompany{}
	err = json.Unmarshal(body, req)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	glog.Infoln(lib.Log("api", "", "Setapplycompany"), req)

	// apply chain to add new account
	resp, err := cli.C.Register(context.Background(), &protocol.ReqRegister{
		PassWord: req.PassWord,
	})
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	// save it to mysql
	db := db.DbClient.Client.Create(&models.Company{
		CompanyName:      req.CompanyName,
		CompanyDespcrite: resp.GetAccountDescribe(),
		CompanyPassWord:  req.PassWord,
		CompanyBanchName: req.BranchName,
		CompanyAddress:   resp.GetUserAddress(),
		CompanyPhone:     req.Phone,
	})
	if db.Error != nil {
		glog.Error(db.Error)
	}
	glog.Infoln(lib.Log("api", "", "Setapplycompany"), resp)
}

type GetApplyCompany struct {
	UserAddress string
}

type CompanyAccount struct {
	CompanyName      string
	CompanyPassWord  string
	CompanyPhone     string
	CompanyDespcrite string
}

// Get apply company
func Getapplycompany(c *gin.Context, cli *client.Client) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	req := &GetApplyCompany{}
	err = json.Unmarshal(body, req)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	company := &models.Company{}
	db.DbClient.Client.Where(&ReqApplyCompany{
		UserAddress: req.UserAddress,
	}).First(company)
	glog.Infoln(lib.Log("api", "", "Getapplycompany"), company)
}
