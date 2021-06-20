package core

import (
	"aip-face-sdk/example/src/global"
	"aip-face-sdk/sdk"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func RenderHome(w http.ResponseWriter, r *http.Request, title string) {
	data := Page{}
	user, _ := GetSession(w, r).GetAttr("user")

	if user != nil {
		data.Title = title
		data.UserID = user.(*User).ID
		data.Username = user.(*User).Username
	}

	t, err := template.ParseFiles("templates/index.html")
	global.CheckError(err)

	err = t.Execute(w, data)
	global.CheckError(err)
}

func RenderLogin(w http.ResponseWriter, r *http.Request, title string) {

	data := Page{}
	success := false
	url := "templates/login.html"

	if r.Method == "GET" {

	}

	if r.Method == "POST" {
		log.Println("r.Method: POST")

		ispersis := r.PostFormValue("ispersis")
		log.Println(ispersis)

		formFile, _, err := r.FormFile("image")
		defer func() {
			_ = formFile.Close()
		}()

		if err != nil {
			log.Println("err: ", err)
			data.Error = "参数错误"
		} else {
			//人脸检测
			aipFace := sdk.NewAipFace()
			aipFace.Construct(global.APP_ID, global.API_KEY, global.SECRET_KEY)
			aipFace.Client.SetConf(global.MTransport, nil)

			face_group := "blog"
			options := map[string]interface{}{
				"max_face_num":     1,
				"match_threshold":  80,
				"quality_control":  "NORMAL",
				"liveness_control": "HIGH",
				"user_id":          "0000",
				"max_user_num":     3,
			}

			bytes, err := ioutil.ReadAll(formFile)
			if err != nil {
				log.Fatal(err)
			}

			image := global.ToBase64(bytes)
			imageType := "BASE64"

			result := global.AipFaceTest.Search(image, imageType, face_group, options)

			errorCode := result["error_code"].(float64)

			//人脸识别成功
			if errorCode == 0 {
				//此处直接赋值管理员账号密码
				username := "admin"
				password := "admin"

				user := FindUserByUsernameAndPassword(username, password)
				if user == nil {
					//message(w, r, "用户名或密码错误，登录失败！")
					data.Error = "用户名或密码错误，登录失败！"
				} else {
					success = true
					url = "/"
					// 登陆成功
					sess := GetSession(w, r)
					sess.SetAttr("user", user)
					log.Println("Login success")
				}
			} else {
				data.Error = result["error_msg"].(string)
			}
		}

	}

	if !success {
		t, err := template.ParseFiles(url)
		global.CheckError(err)
		err = t.Execute(w, data)
		global.CheckError(err)
	} else {
		http.Redirect(w, r, url, 302)
	}
}

func RenderLogout(w http.ResponseWriter, r *http.Request, title string) {
	sess := GetSession(w, r)
	sess.DelAttr("user")
	http.Redirect(w, r, "/", 302)
}

func RenderAdmin(w http.ResponseWriter, r *http.Request, title string) {
	sess := GetSession(w, r)
	user, exist := sess.GetAttr("user")
	if !exist {
		http.Redirect(w, r, "/", 302)
		return
	}

	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/userinfo.html")
		global.CheckError(err)
		err = t.Execute(w, user)
		global.CheckError(err)
		return
	}

	// POST 更新用户信息
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	if global.IsEmpty(username, password, email) {
		message(w, r, "字段不能为空")
		return
	}

	switch user := user.(type) {
	case *User:
		user.Username = username
		user.Password = password
		user.Email = email
		UpdateUser(user)
	default:
		log.Println(":userinfo:user.(type)", user)
	}
	http.Redirect(w, r, "/admin", 302)
}

func message(w http.ResponseWriter, r *http.Request, message string) {
	t, err := template.ParseFiles("templates/message.html")
	global.CheckError(err)

	err = t.Execute(w, map[string]string{"Message": message})
	global.CheckError(err)
}
