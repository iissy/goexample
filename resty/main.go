package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

var client *resty.Client

func init() {
	client = resty.New()
}

func main() {
	for i := 1; i < 77; i++ {
		r := fmt.Sprintf("T3-%d", i)
		params := fmt.Sprintf(`{"warehouseID": "232414789256686855","regionID": "314420765674045445","containerCode": "","skuID": "","railCodes": "%s"}`, r)
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(params).
			SetResult(&ContainerStockVos{}).
			Post("http://10.128.35.203/pms/api/common/queryContainerStockDetail")
		if err != nil {
			log.Panic(err)
		}

		result := resp.Result().(*ContainerStockVos)

		log.Printf("%+v", result)
	}
}

type ContainerStock struct {
	SkuID            string `json:"skuID"`
	ExpireTime       int64  `json:"expireTime"`
	FreezeNum        int    `json:"freezeNum"`
	Num              int    `json:"num"`
	ContainerCode    string `json:"containerCode"`
	LocationPosition string `json:"locationPosition"`
	RailCode         string `json:"railCode"`
	UpBoxNum         int    `json:"upBoxNum"`
}

type ContainerStockVos struct {
	Data map[string][]ContainerStock `json:"data"`
}
