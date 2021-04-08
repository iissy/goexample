// 反序列化，结构体的字段必须大些开头，不然无法承载值
package main

import (
	"encoding/json"
	"fmt"
)

type PSMessageBody struct {
	DeviceType  int32          `json:"deviceType"`
	DeviceSn    string         `json:"deviceSn"`
	WarehouseID string         `json:"warehouseID"`
	Channel     MessageChannel `json:"channel"`
	MsgType     string         `json:"type"`
	MsgMode     MessageMode    `json:"msgMode"`
	MsgSeq      uint64         `json:"msgSeq"`
	Params      string         `json:"params"`
}

type MessageMode uint32
type MessageChannel uint32

func main() {
	var bodyMap PSMessageBody
	str := `{"channel":1,"deviceSN":"jc0014","deviceType":3,"msgMode":2,"msgSeq":345486499292119053,"params":"p","type":"drop","warehouseID":"232414789256686833"}`
	err := json.Unmarshal([]byte(str), &bodyMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bodyMap)
}
