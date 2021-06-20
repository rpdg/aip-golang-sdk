package main

import (
	"aip-face-sdk/example/src/global"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestDetect(t *testing.T) {
	options := map[string]interface{}{
		"quality_control":  "NORMAL",
		"liveness_control": "LOW",
		"max_face_num":     2,
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

	result := global.AipFaceTest.Detect(image, imageType, options)
	fmt.Println(result)
}
