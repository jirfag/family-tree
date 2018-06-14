package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"hash"
	"io"
	"time"
)

var accessKeyID = AppConfig.OSS.AccessKeyID
var accessKeySecret = AppConfig.OSS.AccessKeySecret
var host = AppConfig.OSS.Host
var expireTime = int64(60)
var uploadDir = "upload/"

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func getGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

// ConfigStruct is a struct of config
type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

// PolicyToken load Policy
type PolicyToken struct {
	AccessKeyID string `json:"accessid"  example:"LTAasdI0TzWRZrhibAA"`
	Host        string `json:"host" example:"https://asdasda.oasdasdangzhou.aliyuncs.com"`
	Expire      int64  `json:"expire" example:"1528651096"`
	Signature   string `json:"signature" example:"bfeHZsrF7qnuBNNlTO0odCDw6SM="`
	Policy      string `json:"policy" example:"eyJleHBpcmF0aW9uIjoiMjAxOC0wNi0xMVQwMToxODoxNloiLCJjb25kaXRpb25zIjpbWyJzdGFydHMtd2l0aCIsIiRrZXkiLCJ1cGxvYWQvIl1dfQ=="`
	Directory   string `json:"dir" example:"upload/"`
	Callback    string `json:"callback" example:"https://fmt.fredliang.cn/files/callback"`
}

// CallbackParam is a struct load Callback Params
type CallbackParam struct {
	CallbackURL      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}

// CallBackBody to save callback content
type CallBackBody struct {
	FilePath string `json:"filePath"`
	Table    string `json:"table"`  // including: user, company, group
	Field    string `json:"field"`  // example: avatar, logo, images
	Action   string `json:"action"` // including: init, append{
}

// GetPolicyToken is a func to get PolicyToken
func GetPolicyToken(username string) PolicyToken {
	now := time.Now().Unix()
	expireEnd := now + expireTime
	var tokenExpire = getGmtIso8601(expireEnd)

	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, uploadDir+username+"/")
	config.Conditions = append(config.Conditions, condition)

	var callbackParam CallbackParam
	callbackParam.CallbackURL = AppConfig.OSS.CallBack

	callbackData, _ := json.Marshal(CallBackBody{
		FilePath: "${object}",
		Table:    "CallbackTable",
		Action:   "CallbackAction",
		Field:    "CallbackField"})

	callbackParam.CallbackBody = string(callbackData)
	//"filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/json"
	callbackStr, err := json.Marshal(callbackParam)
	if err != nil {
		fmt.Println("callback json err:", err)
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)

	//calucate signature
	result, _ := json.Marshal(config)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(accessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var policyToken PolicyToken
	policyToken.AccessKeyID = accessKeyID
	policyToken.Host = host
	policyToken.Expire = expireEnd
	policyToken.Signature = string(signedStr)
	policyToken.Directory = uploadDir + username + "/"
	policyToken.Policy = string(debyte)
	policyToken.Callback = string(callbackBase64)

	return policyToken
}
