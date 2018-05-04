package client

import (
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	ctx "golang.org/x/net/context"
	"time"
	"znfz/server/protocol"
)

// 测试
func (this *Client) Check() {
	this.C.SayHello(ctx.Background(), &protocol.Req{
		Name: "liangsihao",
	})
}

// 注册服务
func (this *Client) SetRegister(password string) *protocol.RespRegister {
	usr, _ := this.C.Register(ctx.Background(), &protocol.ReqRegister{
		PassWord: password,
	})
	glog.Infoln("1--注册服务", usr)
	glog.Infoln("1--注册服务", usr.GetAccountDescribe())
	time.Sleep(10 * time.Second)
	return usr
}

func (this Client) band(psw, name, role, desp, addr, phone string) {
	glog.Infoln(desp)
	req := &protocol.ReqBand{
		Name:            name,
		Role:            role,
		PassWord:        psw,
		Phone:           phone,
		AccountDescribe: desp,
		UserAddress:     addr,
	}
	glog.Infoln("1.1 1", req)
	resp, _ := this.C.Band(ctx.Background(), req)
	glog.Infoln("1.1 2", resp)
}

// 注册服务
func (this *Client) GetRegister(address string) {
	acc, _ := this.C.CheckAccount(ctx.Background(), &protocol.ReqCheckAccount{
		UserAddress: address,
	})
	glog.Infoln("2", acc, acc.StatusCode)
}

// 发布排班
func (this *Client) SetSchedule(addr, pass, desp, name, timestamp, company string) {

	glog.Infoln("1", timestamp)
	resp2, _ := this.C.SetSchedule(ctx.Background(), &protocol.ReqScheduling{
		UserAddress:     addr,
		PassWord:        pass,
		AccountDescribe: desp,
		Company:         company,
		TimeStamp:       timestamp,
		Jobs: []*protocol.Job{
			&protocol.Job{
				JobAddress: "",
				Role:       uint32(2),
				Count:      uint32(1),
				Radio:      uint64(1),
			}},
	})
	glog.Infoln("3", resp2)
}

// 查询排班
func (this *Client) GetSchedule(userAddr, company, timestamp string) {
	resp3, err := this.C.GetSchedule(ctx.Background(), &protocol.ReqGetSchedue{
		UserAddress: userAddr,
		CompanyName: company,
		TimeStamp:   timestamp,
	})
	glog.Infoln("3", resp3, err)
}

// 获取以太坊余额
func (this *Client) GetEthBalance(usr string) {
	resp, err := this.C.GetEthBalance(ctx.Background(), &protocol.ReqGetEthBalance{
		UserAddress: usr,
	})
	glog.Infoln("获取以太坊余额", resp, err)
}

func (this *Client) RpcTest() {
	t1 := time.Now()
	resp, err := this.C.SayHello(ctx.Background(), &protocol.Req{Name: "sihao "})
	if err != nil {
		glog.Errorln("Do Format error:" + err.Error())
	} else {
		glog.Infoln("client recving msg ->>", resp, " time:", time.Now().Sub(t1))
	}
}

func (this *Client) FindJob(name, cookerAddr, jobAddress, psw, desp string, role uint32) {
	resp4, _ := this.C.ApplyJob(ctx.Background(), &protocol.ReqFindJob{
		UserAddress:     cookerAddr,
		PassWord:        psw,
		AccountDescribe: desp,
		MyJob: &protocol.Job{
			JobAddress: jobAddress,
			Role:       role,
		},
	})
	glog.Infoln("4", resp4.StatusCode)
}

func (this *Client) GetCanapplyJob(company, cookerAddr string) {
	resp, _ := this.C.GetJob(ctx.Background(), &protocol.ReqGetCanApply{
		UserAddress: cookerAddr,
		CompanyName: company,
		TimeStamp:   "",
	})
	glog.Infoln("4", resp)
}

func (this *Client) GetFindJob(usr, job string) {
	resp, err := this.C.CheckIsOkApplication(ctx.Background(), &protocol.ReqCheckIsOkApplication{
		UserAddress: usr,
		JobAddress:  job,
	})
	glog.Infoln("4", resp, err)
}


func (this *Client) Order(id, Context, addr string, table uint32, money float64) {
	//order := &protocol.Order{
	//	Table:     uint32(table),
	//	TimeStamp: time.Now().String(),
	//	Money:     float64(money),
	//	Content:   Context,
	//}
}

func (this *Client) GetOrder(id, addr string) {
	resp, err := this.C.GetContent(ctx.Background(), &protocol.ReqGetContent{
		OrderId:    id,
		JobAddress: addr,
	})
	glog.Infoln("5", err, string(resp.GetContent().GetContent()), resp)
}

func (this *Client) Pay(msg proto.Message) {
	pay := msg.(*protocol.ReqPay)
	order := &protocol.Order{
		TimeStamp: time.Now().String(),
		Money:     float64(pay.GetMoney()),
	}
	resp, err := this.C.Pay(ctx.Background(), &protocol.ReqPay{
		UserAddress:     pay.GetUserAddress(),
		PassWord:        pay.GetPassWord(),
		AccountDescribe: pay.GetAccountDescribe(),
		Money:           pay.GetMoney(),
		Content:         order,
		JobAddress:      pay.GetJobAddress(),
	})
	glog.Infoln("6", err, resp)
}

func (this *Client) GetMoney(job_addr, account_addr string) {
	resp, err := this.C.GetBalance(ctx.Background(), &protocol.ReqGetBalance{
		SchedueAddress: job_addr,
		UserAddress:    account_addr,
	})
	glog.Infoln("6", resp.GetMoney(), " err:", err)
}

func (this *Client) GetApply(job_addr string) {
	resp, err := this.C.GetApply(ctx.Background(), &protocol.ReqGetStaff{
		JobAddress: job_addr,
	})
	glog.Infoln("5", resp.GetStaffs(), " err:", err)
}

func (this *Client) Login(phone string) {
	resp, err := this.C.Login(ctx.Background(), &protocol.ReqLogin{
		Phone: phone,
	})
	glog.Infoln("8", resp, " err:", err)
}

func (this *Client) GetAllOrder(user, job string) {
	resp, _ := this.C.GetAllOrderBySchedule(ctx.Background(), &protocol.ReqGetAllOrder{
		UserAddress: user,
		JobAddress:  job,
	})
	glog.Infoln("8", resp.GetSum(), " v = ", resp)
}

func (this *Client) GetAllMoney(user, company string) {
	resp, _ := this.C.GetAllMoney(ctx.Background(), &protocol.ReqGetAllMoney{
		UserAddress: user,
		CompanyName: company,
	})
	glog.Infoln("9 ", resp)
}

func (this *Client) HistoryJoin(uaddr, comp string) {
	resp, _ := this.C.HistoryJoin(ctx.Background(), &protocol.ReqHistoryJoin{
		UserAddress: uaddr,
		Company:     comp,
	})
	glog.Infoln("HistoryJoin", resp)
}

func (this *Client) GetAllIncome(uaddr, comp string) {
	resp, _ := this.C.GetAllIncome(ctx.Background(), &protocol.ReqGetAllIncome{
		UserAddress: uaddr,
		CompanyName: comp,
	})
	glog.Infoln("GetAllIncome", resp)
}
