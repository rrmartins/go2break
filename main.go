package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "github.com/rrmartins/gobreak/go2break/routers"
)

func init() {
	beego.Info("aqui 0.1")
	orm.RegisterDriver("postgres", orm.DRPostgres)
	beego.Info("aqui 0.2")
	orm.RegisterDataBase("default", "postgres", "postgres://codex:123qwe@127.0.0.1/gobreak?sslmode=disable")
	beego.Info("aqui 0.3")
}

func main() {
	beego.Run()
}
