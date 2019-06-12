package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterModel(
		new(FinanceItem),
		new(Ledger),
		new(FinanceSubjcet))
	orm.RunSyncdb("default", false, true)
}
