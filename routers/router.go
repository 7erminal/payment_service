// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"payment_service/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/payments",
			beego.NSInclude(
				&controllers.PaymentsController{},
			),
		),
		beego.NSNamespace("/payment-methods",
			beego.NSInclude(
				&controllers.Payment_methodsController{},
			),
		),
		beego.NSNamespace("/payment-categories",
			beego.NSInclude(
				&controllers.Payment_categoriesController{},
			),
		),
		beego.NSNamespace("/expenses",
			beego.NSInclude(
				&controllers.Expense_recordsController{},
			),
		),
		beego.NSNamespace("/callback",
			beego.NSInclude(
				&controllers.CallbackController{},
			),
		),
		beego.NSNamespace("/request-money",
			beego.NSInclude(
				&controllers.Request_moneyController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
