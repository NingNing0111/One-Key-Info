package model

type KeyInfo struct {
	KeyValue    string  `json:"keyValue"`
	TotalAmount float32 `json:"totalAmount"`
	TotalUsage  float32 `json:"totalUsage"`
	AccessUntil string  `json:"accessUntil"`
	Remaining   float32 `json:"remaining"`
}
