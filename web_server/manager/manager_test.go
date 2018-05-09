package manager

import (
	"bytes"
	"github.com/gin-gonic/gin/json"
	"github.com/golang/glog"
	"net/http"
	"testing"
)

func TestManager(t *testing.T) {
	glog.Infoln("start testing")
	//req := apply()
	//req := bind()
	req := getbind()
	c := http.Client{}
	_, err := c.Do(req)
	if err != nil {
		glog.Errorln("err", err)
	}
}

func apply() *http.Request {
	glog.Infoln("start testing apply")
	j, _ := json.Marshal(&ReqApplyCompany{
		CompanyName: "hhhh",
		BranchName:  "01",
		UserAddress: "0x1111",
		PassWord:    "5555",
		Phone:       "6666",
	})
	buf := bytes.NewBuffer(j)
	req, err := http.NewRequest("POST", "http://localhost:8088/setapplycompany", buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}

func bind() *http.Request {
	glog.Infoln("start testing bind")
	j, _ := json.Marshal(&ReqBinder{
		CompanyName: "hhhh",
		BranchName:  "01",
		UserAddress: "0x1111",
		Phone:       "6666",
	})
	buf := bytes.NewBuffer(j)
	req, err := http.NewRequest("POST", "http://localhost:8088/setbindmsg", buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}

func getbind() *http.Request {
	glog.Infoln("start testing getbind")
	j, _ := json.Marshal(&ReqBindMsg{
		UserAddress: "0x1111",
		Phone:       "6666",
	})
	buf := bytes.NewBuffer(j)
	req, err := http.NewRequest("POST", "http://localhost:8088/getbindmsg", buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}
