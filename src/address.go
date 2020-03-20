package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type location struct {
	Lng float64 `json:"lng"` // 经度
	Lat float64 `json:"lat"`  // 维度
}

type result struct {
	Location      location `json:"location"`
	Precise       int      `json:"precise"`
	Confidence    int      `json:"confidence"`
	Comprehension int      `json:"comprehension"`
	Level         string   `json:"level"`
}

// AddInfo 地址信息.
type AddInfo struct {
	Status int8   `json:"status"`
	Result result `json:"result"`
}

// GetAddressJSON : 将百度地图api返回的JSON字符串转为结构体.
func GetAddressJSON(address string) (add *AddInfo, err error) {
	// 访问地址
	ak := "您的ak"
	url := fmt.Sprintf("http://api.map.baidu.com/geocoding/v3/?address=%s&output=json&ak=%s", address,ak)
	
	// 建立http请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("访问错误，", err)
	}

	// 获取返回结构体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取内容错误，", err)
	}
	
	add = &AddInfo{}
	err = json.Unmarshal(body, add)
	return add, err
}