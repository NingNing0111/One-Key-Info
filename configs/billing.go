package configs

type Billing struct {
	UrlSub   string
	UrlUsage string
	UrlModel string
}

var b = Billing{}

func init() {
	b = Billing{
		UrlSub:   "v1/dashboard/billing/subscription",
		UrlUsage: "v1/dashboard/billing/usage",
		UrlModel: "v1/models",
	}
}

func GetBillingConf() Billing {
	return b
}
