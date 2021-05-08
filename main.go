package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "grids/routers"
)

func main() {

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))

	beego.Run()
}

