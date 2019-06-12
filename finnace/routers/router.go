package routers

import (
	"finnace/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/finance", &controllers.FinanceController{})
	beego.Router("/ledger", &controllers.LedgerController{})
	beego.Router("/financeSubject", &controllers.FinanceSubjectController{})

}
