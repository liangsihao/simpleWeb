package db

import (
	"testing"
	"time"
)

type HistoryDeploy struct {
	Id         uint      "gorm:PRIMARY KEY"
	CreateTime time.Time //创建时间
	Address    string    // 地址
	Account    string    // 账户
	Version    string    // 版本
	SubVersion string    // 子版本號
	RequsetIp  string    // 請求地址
	Menthon    string    // 請求方式
}

type Company struct {
	ID               uint      "gorm:PRIMARY KEY"
	CompanyName      string    `gorm:"not null"`
	CompanyBanchName string    `gorm:"not null"`
	CompanyPassWord  string    `gorm:"not null"`
	CompanyAddress   string    `gorm:"not null"`
	CompanyPhone     string    `gorm:"not null"`
	CompanyDespcrite string    `gorm:"type:text;not null"`
	CreateTime       time.Time `gorm:"not null"`
}

type Orders struct {
	ID         uint      `gorm:"primary_key"`
	CreateTime time.Time `gorm:"not null"` // 创建时间
	Company    string    `gorm:"not null"` // company name
	BranchShop string    `gorm:"not null"` // branch shop name
	Content    string    `gorm:"not null"` // infomation of the order
	Bill       string    `gorm:"not null"` // information of the bill
	Price      string    `gorm:"not null"` // price of the order
	GasPrice   string    `gorm:"not null"` // gas price
}

type Binder struct {
	ID             uint      "gorm:PRIMARY KEY"
	CompanyId      uint      `gorm:"not null"`
	CompanyName    string    `gorm:"not null"`
	BrenchName     string    `gorm:"not null"`
	UserAddress    string    `gorm:"not null"`
	CompanyAddress string    `gorm:"not null"`
	BindTime       time.Time `gorm:"not null"`
}

func TestInitDb(t *testing.T) {
	InitDb("launch:launch@tcp(39.108.80.66:3306)/test_order?charset=utf8&parseTime=true&loc=Local")
	//DbClient.Client.CreateTable(&Orders{})
	for i := 0; i < 10; i++ {
		DbClient.Client.CreateTable(&Binder{})
		DbClient.Client.CreateTable(&Company{})
		//err:=DbClient.Client.Model(&Orders{}).Create(&Orders{
		//	CreateTime: time.Now(),
		//	Company:    "launch",
		//	BranchShop: "yanjiuyuan",
		//	Content:    time.Now().String(),
		//})
		//if err!=nil{
		//	glog.Errorln(err)
		//}
		//DbClient.Client.Delete(&Orders{})
	}
}
