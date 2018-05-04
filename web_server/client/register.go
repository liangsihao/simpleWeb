package client

import "znfz/server/protocol"

func (this *Client) Init() {
	this.Register(this.Pay, &protocol.ReqPay{})
}
