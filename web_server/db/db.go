package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var DbClient *Db

type Db struct {
	addr   string       // the addr of db server
	Lock   sync.RWMutex // lock
	Client *gorm.DB     // mysql client
}

func InitDb(addr string) {
	glog.Infoln("starting db")
	mydb := &Db{}
	mydb.addr = addr
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		glog.Errorln("db initing fail", err)
		return
	}
	err = db.DB().Ping()
	if err != nil {
		glog.Errorln("db ping fail", err)
		return
	}else {
		glog.Infoln("connecting db success!")
	}
	mydb.Client = db
	DbClient = mydb

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(false)
}

func (this *Db) CreateTable(models interface{}) {
	this.Client.CreateTable(models)
}
