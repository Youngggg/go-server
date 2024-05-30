package enum

const (
	//provider
	IP_PROVIDER_MYSELF = "10086"
	IP_PROVIDER_PY     = "py"
)

const (
	CountryAny     = "ANY"
	IpCountryPrior = 1 // 对应国家优先，没有则取任意国家
	IpCountryAny   = 2 // 不要求国家
)
