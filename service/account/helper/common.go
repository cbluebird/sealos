package helper

const (
	GROUP                    = "/account/v1alpha1"
	GetAccount               = "/account"
	GetPayment               = "/payment"
	GetHistoryNamespaces     = "/namespaces"
	GetProperties            = "/properties"
	GetRechargeAmount        = "/costs/recharge"
	GetConsumptionAmount     = "/costs/consumption"
	GetPropertiesUsed        = "/costs/properties"
	GetAPPCosts              = "/costs/app"
	SetPaymentInvoice        = "/payment/set-invoice"
	GetUserCosts             = "/costs"
	SetTransfer              = "/transfer"
	GetTransfer              = "/get-transfer"
	GetRegions               = "/regions"
	GetOverview              = "/cost-overview"
	GetAppList               = "/cost-app-list"
	GetAppTypeList           = "/cost-app-type-list"
	GetBasicCostDistribution = "/cost-basic-distribution"
	GetAppCostTimeRange      = "/cost-app-time-range"
	CheckPermission          = "/check-permission"
	GetInvoice               = "/invoice/get"
	ApplyInvoice             = "/invoice/apply"
	SetStatusInvoice         = "/invoice/set-status"
	GetInvoicePayment        = "/invoice/get-payment"
	UseGiftCode              = "/gift-code/use"
)

// env
const (
	ConfigPath         = "/config/config.json"
	EnvMongoURI        = "MONGO_URI"
	ENVGlobalCockroach = "GLOBAL_COCKROACH_URI"
	ENVLocalCockroach  = "LOCAL_COCKROACH_URI"
	EnvLocalRegion     = "LOCAL_REGION"
)
