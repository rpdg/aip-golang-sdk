package global

import (
	"aip-face-sdk/sdk"
	"encoding/base64"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	APP_ID     = "your id"
	API_KEY    = "your key"
	SECRET_KEY = "your secret key"

	USER_NAME     = "your name"
	SUCCESS_SCORE = 90.00
)

var Header = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36",
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
	"Accept-Encoding": "utf-8",
	"Connection":      "keep-alive",
}

//uri, err := url.Parse("http://username:password@inproxy.sjtu.edu.cn:8000")
//
//if err != nil{
//	log.Fatal("parse url error: ", err)
//}

var MTransport = &http.Transport{
	IdleConnTimeout:       time.Second * 2048,
	ResponseHeaderTimeout: time.Second * 5,
	DialContext: (&net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 60 * time.Second,
	}).DialContext,
}

var AipFaceTest = sdk.NewAipFace()

func init() {
	AipFaceTest.Construct(APP_ID, API_KEY, SECRET_KEY)
	AipFaceTest.Client.SetConf(MTransport, nil)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func IsEmpty(strs ...string) (isEmpty bool) {
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if str == "" || len(str) == 0 {
			isEmpty = true
			return
		}
	}
	isEmpty = false
	return
}
