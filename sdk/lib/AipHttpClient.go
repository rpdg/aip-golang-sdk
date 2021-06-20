package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type array map[string]string

/**
HttpClient
headers HTTP header
*/
type AipHttpClient struct {
	headers   map[string]string
	transport *http.Transport
}

func NewAipHttpClient() *AipHttpClient {
	return &AipHttpClient{
		headers: make(map[string]string),
	}
}

/**
配置
*/
func (client *AipHttpClient) SetConf(ts *http.Transport, proxyUrl *url.URL) {
	client.transport = ts

	if proxyUrl != nil {
		client.transport.Proxy = http.ProxyURL(proxyUrl)
	}

}

/**
Post请求
@param string url
@param map data
@param map params HTTP URL
@param map headers HTTP header
*/
func (client *AipHttpClient) Post(url string, data interface{}, params array, headers array) map[string]interface{} {

	obj := make(map[string]interface{})

	url = buildUrl(url, params)
	log.Println(url)
	client.buildHeaders(headers)
	log.Println(client.headers)

	//转为json格式
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("json.Marshal err: ", err)
		return obj
	}

	c := http.Client{
		Transport: client.transport,
	}

	reader := bytes.NewReader(b)
	log.Println(string(b))
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Println("HttpRequest err: ", err)
		return obj
	}
	log.Println("Method: ", req.Method)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.Do(req)

	defer func() {
		resp.Body.Close()
	}()

	if err != nil {
		log.Println("Client Do err: ", err)
		return obj
	}

	if resp.StatusCode == 0 {
		log.Println("Code err: ", 0)
		return obj
	}

	body, err := ioutil.ReadAll(resp.Body)
	content, _ := UnescapeUnicode(body)

	obj["code"] = resp.StatusCode
	obj["content"] = string(content)

	return obj

}

/**
Get请求
@param string url
@param map params HTTP URL
@param map headers HTTP header
*/
func (client *AipHttpClient) Get(url string, params array, headers array) map[string]interface{} {

	obj := make(map[string]interface{})

	url = buildUrl(url, params)
	log.Println(url)
	client.buildHeaders(headers)

	c := http.Client{
		Transport: client.transport,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("HttpRequest err: ", err)
		return obj
	}
	log.Println("Method: ", req.Method)

	for k, v := range client.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.Do(req)

	defer func() {
		resp.Body.Close()
	}()

	if err != nil {
		log.Println("Client Do err: ", err)
		return obj
	}

	if resp.StatusCode == 0 {
		log.Println("Code err: ", 0)
		return obj
	}

	body, err := ioutil.ReadAll(resp.Body)
	content, _ := UnescapeUnicode(body)

	log.Println(string(content))

	obj["code"] = resp.StatusCode
	obj["content"] = string(content)

	return obj

}

/**
构造 header
@param map headers
*/
func (client *AipHttpClient) buildHeaders(headers map[string]string) {
	for k, v := range headers {
		client.headers[k] = v
	}
}

/**
@param string url
@param map params 参数
@return string
*/
func buildUrl(url string, params map[string]string) (param_str string) {
	params_arr := make([]string, 0, len(params))
	for k, v := range params {
		params_arr = append(params_arr, fmt.Sprintf("%s=%s", k, v))
	}
	//fmt.Println(params_arr)
	param_str = strings.Join(params_arr, "&")

	if len(param_str) > 0 {
		param_str = strings.Join([]string{url, param_str}, "?")
	} else {
		param_str = url
	}

	return param_str
}
