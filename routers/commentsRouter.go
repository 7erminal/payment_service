package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["payment_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:AccountsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:CallbackController"] = append(beego.GlobalControllerRouter["payment_service/controllers:CallbackController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/process`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"],
        beego.ControllerComments{
            Method: "GetAllByBranch",
            Router: `/branch/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Expense_recordsController"],
        beego.ControllerComments{
            Method: "GetExpenseRecordCount",
            Router: `/count/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_categoriesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_historyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_methodsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Payment_typesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"],
        beego.ControllerComments{
            Method: "GetPaymentCount",
            Router: `/count/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"],
        beego.ControllerComments{
            Method: "NameInquiry",
            Router: `/name-inquiry`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"] = append(beego.GlobalControllerRouter["payment_service/controllers:PaymentsController"],
        beego.ControllerComments{
            Method: "UploadPaymentProof",
            Router: `/upload-payment-proof`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["payment_service/controllers:Request_moneyController"] = append(beego.GlobalControllerRouter["payment_service/controllers:Request_moneyController"],
        beego.ControllerComments{
            Method: "RequestMoneyViaMomo",
            Router: `/momo/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
