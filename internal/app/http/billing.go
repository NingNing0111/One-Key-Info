package http

import (
	"log"
	"onekeyinfo/configs"
	"onekeyinfo/internal/app/model"
	"onekeyinfo/internal/app/util"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

var client *resty.Client = nil
var env = configs.Billing{}
var (
	BASEURL = ""
)

func init() {
	client = resty.New()
	env = configs.GetBillingConf()
}

func GetKeyInfo(key string, baseUrl string) model.KeyInfo {
	BASEURL = baseUrl
	subData := getSub(&key)
	usage := getUsage(&key)
	remaining := subData.HardLimitUsd - usage.TotalUsage
	key = key[:5] + "**********************" + key[len(key)-5:]
	accessUntil := ""
	if subData.AccessUntil == 0 {
		accessUntil = "永不过期"
	} else {
		accessUntil = time.Unix(subData.AccessUntil, 0).Format("2006-01-02")
	}
	info := model.KeyInfo{
		KeyValue:    key,
		TotalAmount: subData.HardLimitUsd,
		TotalUsage:  usage.TotalUsage,
		AccessUntil: accessUntil,
		Remaining:   remaining,
	}

	return info
}

func getSub(key *string) SubData {
	url := BASEURL + env.UrlSub
	log.Println("请求地址：", url)
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+*key).
		SetHeader("Content-type", "application/json").
		Get(url)

	subData := SubData{}

	if err != nil {
		log.Fatal(err, " ---> 请求失败")
	} else {
		util.Parse(resp.String(), &subData)
	}
	return subData
}

func getUsage(key *string) Usage {
	url := BASEURL + env.UrlUsage
	//// 开始时间与截止时间
	start := strconv.Itoa(time.Now().Year()) + "-" + "01-01"
	end := strconv.Itoa(time.Now().Year()) + "-" + strconv.Itoa(int(time.Now().Month())) + "-" + strconv.Itoa(time.Now().Day())
	log.Println("请求地址：", url)
	log.Println("查询时间：", start, "--->", end)
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+*key).
		SetHeader("Content-type", "application/json").
		SetQueryParam("start_date", time.Now().Format(start)).
		SetQueryParam("end_date", time.Now().Format(end)).
		Get(url)

	usage := Usage{}
	if err != nil {
		log.Fatal(url, " 解析失败")
	} else {
		util.Parse(resp.String(), &usage)
	}
	usage.TotalUsage /= 100
	return usage
}
