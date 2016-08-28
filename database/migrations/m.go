package main

import(
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/migration"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func init(){
	orm.RegisterDataBase("default", "postgres","postgres://codex:123qwe@127.0.0.1/gobreak?sslmode=disable")
}

func main(){
	task := "upgrade"
	switch task {
	case "upgrade":
		if err := migration.Upgrade(1472333916); err != nil {
			os.Exit(2)
		}
	case "rollback":
		if err := migration.Rollback("Clients_20160827_213836"); err != nil {
			os.Exit(2)
		}
	case "reset":
		if err := migration.Reset(); err != nil {
			os.Exit(2)
		}
	case "refresh":
		if err := migration.Refresh(); err != nil {
			os.Exit(2)
		}
	}
}

