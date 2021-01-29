package uhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {

}

func Post(url string, params map[string]interface{}) (res map[string]interface{}) {
	jsonByte, err := json.Marshal(params)
	if err != nil {
		fmt.Println("解析json错误")
		return
	}
	reader := bytes.NewReader(jsonByte)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	_ = json.Unmarshal(respBytes, &res)
	return
}
