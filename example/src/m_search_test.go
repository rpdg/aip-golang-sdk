package main

import (
	"aip-face-sdk/example/src/global"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestSearch(t *testing.T) {
	face_group := "blog"
	options := map[string]interface{}{
		"max_face_num":     1,
		"match_threshold":  80,
		"quality_control":  "NORMAL",
		"liveness_control": "LOW",
		"user_id":          "0000",
		"max_user_num":     3,
	}

	// Read the entire file into a byte slice

	//path,_ := aipFace.GetCurrentPath()
	//bytes, err := ioutil.ReadFile(path + "feng.jpg")

	bytes, err := ioutil.ReadFile("detect.jpg")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	//// Determine the content type of the image file
	//mimeType := http.DetectContentType(bytes)
	//
	//// Prepend the appropriate URI scheme header depending
	//// on the MIME type
	//switch mimeType {
	//case "image/jpeg":
	//	base64Encoding += "data:image/jpeg;base64,"
	//case "image/png":
	//	base64Encoding += "data:image/png;base64,"
	//}
	//
	//// Append the base64 encoded output
	//base64Encoding += toBase64(bytes)

	base64Encoding = global.ToBase64(bytes)

	image := base64Encoding
	imageType := "BASE64"

	result := global.AipFaceTest.Search(image, imageType, face_group, options)
	fmt.Println(result)
}
