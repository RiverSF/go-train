package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

type Adstxt struct {
	ID      int    `orm:"column(id);auto"`
	Domain  string `orm:"column(domain)"`
	MediaId int    `orm:"-"` //忽略该字段
	//UpdatedAt 	time.Time	`orm:"auto_now;type(timestamp)"`
}

func (a *Adstxt) TableName() string {
	return "saas_adx_adstxt"
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(Adstxt))

	// need to register db driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// need to register default database
	orm.RegisterDataBase("default", "mysql", "ssp_test:Abc@1234@tcp(10.12.1.6:3306)/ssp_test?charset=utf8mb4")
}

func main() {
	// automatically build table
	//orm.RunSyncdb("default", false, true)

	orm.Debug = true

	// create orm object
	o := orm.NewOrm()

	// data
	//ads := new(Adstxt)

	txt := &Adstxt{}
	txt.ID = 1
	o.Read(txt)
	fmt.Println(*txt)

	// insert data
	//o.Insert(ads)
}
