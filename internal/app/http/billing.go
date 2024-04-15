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
	BASEURL = "https://api.openai.com/"
)

func init() {
	client = resty.New()
	env = configs.GetBillingConf()
}

func GetKeyInfo(key string, baseUrl string) model.KeyInfo {
	BASEURL = baseUrl
	subData, err := getSub(&key)
	if err != nil {
		log.Fatal(err)
	}
	usage, err := getUsage(&key)
	if err != nil {
		log.Fatal(err)
	}
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

func getSub(key *string) (SubData, error) {
	url := BASEURL + env.UrlSub
	log.Println("请求地址：", url)
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+*key).
		SetHeader("Content-type", "application/json").
		Get(url)
	if err != nil {
		log.Fatal(err, " ---> 请求失败")
		return SubData{}, err
	}
	subData := SubData{}
	err = util.Parse(resp.String(), &subData)
	if err != nil {
		return SubData{}, err
	}
	return subData, nil
}

func getUsage(key *string) (Usage, error) {
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

	if err != nil {
		log.Fatal(url, " 解析失败")
		return Usage{}, err
	}
	usage := Usage{}
	err = util.Parse(resp.String(), &usage)
	if err != nil {
		return Usage{}, err
	}
	usage.TotalUsage /= 100
	return usage, nil
}
