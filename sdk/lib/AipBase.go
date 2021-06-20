package lib

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	/**
	 * 获取access token url
	 * @var string
	 */
	ACCESS_TOKEN_URL = "https://aip.baidubce.com/oauth/2.0/token"

	/**
	 * 反馈接口
	 * @var string
	 */
	REPORT_URL = "https://aip.baidubce.com/rpc/2.0/feedback/v1/report"

	SCOPE = "brain_all_scope"
)

type AipBase struct {
	/**
	  appId
	  @var string
	*/
	AppId string

	/**
	  apiKey
	  @var string
	*/
	ApiKey string

	/**
	  secretKey
	  @var string
	*/
	SecretKey string

	/**
	  权限
	  @var array
	*/
	Scope string

	//+++++++++++++++++++++++++++++++++++
	isCloudUser bool
	Client      *AipHttpClient
	version     string
}

func (base *AipBase) Construct(appId string, apiKey string, secretKey string) {
	base.AppId = strings.Trim(appId, " ")
	base.ApiKey = strings.Trim(apiKey, " ")
	base.SecretKey = strings.Trim(secretKey, " ")
	base.Client = NewAipHttpClient()
	base.isCloudUser = false
	base.version = "2_2_20"
	base.Scope = SCOPE
}

/**
查看版本
@return string
*/
func (base *AipBase) GetVersion() (ver string) {
	return base.version
}

/**
格式检查
@param string url
@param string data
@return mix
*/
func (base *AipBase) validate(url string, data interface{}) (mix bool) {
	return true
}

/**
处理请求参数
@param string url
@param map params
@param string data
@param map headers
*/
func (base *AipBase) proccessRequest(url string, params map[string]string, data interface{}, headers map[string]string) {
	//使用php接口。。。
	params["aipSdk"] = "golang"
	params["aipSdkVersion"] = base.version
}

/**
返回 access token 路径
@return string
*/
func (base *AipBase) getAuthFilePath() (path string) {
	dir, err := base.GetCurrentPath()
	if err != nil {
		log.Println("os.Getwd err: ", err)
		return ""
	}

	return dir + Md5(base.ApiKey)
}

/**
写入本地文件
@param map obj
@return
*/
func (base *AipBase) writeAuthObj(obj map[string]interface{}) {
	if len(obj) == 0 {
		return
	}

	if obj["is_read"] != nil && obj["is_read"] == true {
		return
	}

	obj["time"] = time.Now().Unix()
	obj["is_cloud_user"] = base.isCloudUser

	//写入文件
	b, err := json.Marshal(obj)
	if err != nil {
		log.Println("json.Marshal err: ", err)
		return
	}

	err = ioutil.WriteFile(base.getAuthFilePath(), b, 0644)
	if err != nil {
		log.Println("ioutil.WriteFile err: ", err)
	}
}

/**
读取本地缓存
*/
func (base *AipBase) readAuthObj() map[string]interface{} {

	obj := make(map[string]interface{})

	content, err := ioutil.ReadFile(base.getAuthFilePath())
	if err != nil {
		log.Println("ioutil.ReadFile err: ", err)
		return obj
	}

	if content != nil {
		err = json.Unmarshal(content, &obj)
		if err != nil {
			log.Println("json.Unmarshal err: ", err)
			return obj
		}

		val, _ := obj["is_cloud_user"].(bool)
		base.isCloudUser = val

		obj["is_read"] = true

		t_time, _ := obj["time"].(int64)
		t_expires_in, _ := obj["expires_in"].(int64)
		t := t_time + t_expires_in - 30

		if base.isCloudUser || t > time.Now().Unix() {
			return obj
		}
	}

	return obj
}

/**
认证
*/
func (base *AipBase) auth(refresh bool) map[string]interface{} {

	obj := make(map[string]interface{})

	//非过期刷新
	if !refresh {
		obj = base.readAuthObj()
		if len(obj) != 0 {
			return obj
		}
	}

	response := base.Client.Get(ACCESS_TOKEN_URL, map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     base.ApiKey,
		"client_secret": base.SecretKey,
	}, nil)

	data, _ := response["content"].(string)
	obj = base.processResult(data)

	base.isCloudUser = !base.isPermission(obj)
	return obj
}

/**
判断认证是否有权限
@param map authObj
@return bool
*/
func (base *AipBase) isPermission(authObj map[string]interface{}) (ok bool) {
	if len(authObj) == 0 || authObj["scope"] == nil {
		return false
	}

	scopes, ok := authObj["scope"].(string)
	log.Println("Scopes: ", scopes)

	return strings.Contains(scopes, base.Scope)
}

/**
Api请求
@param string url
@param string data
*/
func (base *AipBase) Request(url string, data interface{}, headers map[string]string) map[string]interface{} {

	obj := make(map[string]interface{})

	result := base.validate(url, data)
	if result != true {
		return obj
	}

	params := make(map[string]string)
	authObj := base.auth(false)

	if base.isCloudUser == false {
		token, _ := authObj["access_token"].(string)
		params["access_token"] = token
	}

	//特殊处理
	base.proccessRequest(url, params, data, headers)

	headers = base.getAuthHeaders("POST", url, params, headers)
	response := base.Client.Post(url, data, params, headers)

	respStr, _ := response["content"].(string)
	obj = base.processResult(respStr)

	log.Println("auth error_code", obj["error_code"])

	errorCode := obj["error_code"].(float64)
	if !base.isCloudUser && obj["error_code"] != nil && errorCode == 110 {
		authObj = base.auth(true)

		token, _ := authObj["access_token"].(string)
		params["access_token"] = token
		response = base.Client.Post(url, data, params, headers)
		respStr, _ := response["content"].(string)
		obj = base.processResult(respStr)
	}

	if len(obj) == 0 || obj["error_code"] == nil {
		base.writeAuthObj(authObj)
	}

	return obj
}

/**
@param string method
@param string url
@param map params
@param map headers
*/
func (base *AipBase) getAuthHeaders(method string, url string, params map[string]string, headers map[string]string) map[string]string {

	//不是云的用户不需要在header中签名认证
	if base.isCloudUser == false {
		return headers
	}

	/**
	未修改完。。。
	*/

	return headers
}

/**
反馈
*/
func (base *AipBase) Report(feedback map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	data["feedback"] = feedback

	return base.Request(REPORT_URL, data, nil)
}

func (base *AipBase) GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return "", errors.New(`Can‘t find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func (base *AipBase) processResult(content string) map[string]interface{} {
	data := make(map[string]interface{})

	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Println("json.Unmarshal err: ", err)
		return data
	}

	return data
}
