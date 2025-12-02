package functions

import (
	"bytes"
	"encoding/json"
	"io"
	"payment_service/api"
	"payment_service/structs/requests"
	"payment_service/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func HubtelRequestViaMobileMoney(c *beego.Controller, req requests.HubtelMomoPaymentRequestDTO) (responses.HubtelPaymentRequestApiResponseDTO, error) {
	host, _ := beego.AppConfig.String("hubtelReceiveBaseUrl")
	salesId, _ := beego.AppConfig.String("hubtelSalesID")
	authorizationKey, _ := beego.AppConfig.String("authorizationKey")

	logs.Info("Sending phone number ", req.CustomerMsisdn)
	logs.Info("Network is ", req.Network)
	logs.Info("Callback URL is ", req.PrimaryCallbackUrl)
	logs.Info("Amount is ", req.Amount)
	logs.Info("Destination is ", req.Channel)

	// serviceId, _ := helpers.GetServiceId(req.Network)

	reqText, _ := json.Marshal(req)

	logs.Info("Request to process payment request: ", string(reqText))

	request := api.NewRequest(
		host,
		"/"+salesId+"/receive/mobilemoney",
		api.POST)
	request.HeaderField["Authorization"] = "Basic " + authorizationKey
	request.InterfaceParams["CustomerName"] = req.CustomerName
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["PrimaryCallbackUrl"] = req.PrimaryCallbackUrl
	request.InterfaceParams["ClientReference"] = req.ClientReference
	request.InterfaceParams["CustomerMsisdn"] = req.CustomerMsisdn
	request.InterfaceParams["CustomerEmail"] = req.CustomerEmail
	request.InterfaceParams["Channel"] = req.Channel
	request.InterfaceParams["Description"] = req.Description

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.HubtelPaymentRequestApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data, nil
}

func HubtelSendMoneyViaMobileMoney(c *beego.Controller, req requests.HubtelMomoPaymentRequestDTO) (responses.HubtelSendPaymentApiResponseDTO, error) {
	host, _ := beego.AppConfig.String("hubtelSendBaseUrl")
	prepaidId, _ := beego.AppConfig.String("hubtelPrepaidDepositID")
	authorizationKey, _ := beego.AppConfig.String("authorizationKey")

	logs.Info("Sending phone number ", req.CustomerMsisdn)
	logs.Info("Network is ", req.Network)
	logs.Info("Callback URL is ", req.PrimaryCallbackUrl)
	logs.Info("Amount is ", req.Amount)
	logs.Info("Destination is ", req.Channel)

	// serviceId, _ := helpers.GetServiceId(req.Network)

	reqText, _ := json.Marshal(req)

	logs.Info("Request to process payment request: ", string(reqText))

	request := api.NewRequest(
		host,
		"/"+prepaidId+"/send/mobilemoney",
		api.POST)
	request.HeaderField["Authorization"] = "Basic " + authorizationKey
	request.InterfaceParams["RecipientName"] = req.CustomerName
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["PrimaryCallbackUrl"] = req.PrimaryCallbackUrl
	request.InterfaceParams["ClientReference"] = req.ClientReference
	request.InterfaceParams["RecipientMsisdn"] = req.CustomerMsisdn
	request.InterfaceParams["CustomerEmail"] = req.CustomerEmail
	request.InterfaceParams["Channel"] = req.Channel
	request.InterfaceParams["Description"] = req.Description

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.HubtelSendPaymentApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data, nil
}

func HubtelNameInquiry(c *beego.Controller, mobileNumber string, channel string) (responses.HubtelNameInquiryApiResponseDTO, error) {
	host, _ := beego.AppConfig.String("hubtelNameInquiryBaseUrl")
	prepaidId, _ := beego.AppConfig.String("hubtelPrepaidDepositID")
	authorizationKey, _ := beego.AppConfig.String("authorizationKey")

	logs.Info("Sending phone number ", mobileNumber)

	// serviceId, _ := helpers.GetServiceId(req.Network)

	request := api.NewRequest(
		host,
		"/"+prepaidId+"/mobilemoney/verify?channel="+channel+"&customerMsisdn="+mobileNumber,
		api.POST)
	request.HeaderField["Authorization"] = "Basic " + authorizationKey

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.HubtelNameInquiryApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data, nil
}
