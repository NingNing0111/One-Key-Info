package http

type SubData struct {
	SoftLimitUsd float32 `json:"soft_limit_usd"`
	HardLimitUsd float32 `json:"hard_limit_usd"`
	AccessUntil  int64   `json:"access_until"`
}

type Usage struct {
	TotalUsage float32 `json:"total_usage"`
}
