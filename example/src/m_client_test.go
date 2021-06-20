package main

import (
	"aip-face-sdk/example/src/global"
	"testing"
)

type myData map[string]interface{}

func TestClientGet(t *testing.T) {
	params := map[string]string{
		"city": "1",
		"key":  "051b1dc0059e92f",
	}

	global.AipFaceTest.Client.Get("http://apis.juhe.cn/simpleWeather/query", params, global.Header)
}

func TestClientPost(t *testing.T) {
	datas := []myData{
		map[string]interface{}{
			"city": 1,
		},
		map[string]interface{}{
			"city": 2,
		},
	}

	global.AipFaceTest.Request("http://apis.juhe.cn/simpleWeather/query", datas, map[string]string{
		"Content-Type": "application/json",
	})
}
