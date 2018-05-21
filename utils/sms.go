package utils

import (
	"fmt"
	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/gin-gonic/gin"
	"github.com/tobyzxj/uuid"
)

func SendSMS(Phone, Template, Dict string) (isOK bool, msg string, errID string) {
	if gin.Mode() == "test" {
		if Phone == "17777766667" || Phone == "" {
			return false, "Testing err sms", ""
		}
		return true, "Done", ""
	}
	dysms.HTTPDebugEnable = true
	dysms.SetACLClient(AppConfig.Dayu.AccessID, AppConfig.Dayu.AccessKey)

	// send to some one
	respSendSms, err := dysms.SendSms(uuid.New(), Phone, "联创团队", Template, Dict).DoActionWithException()
	respond := respSendSms.GetRequestID()
	if (err != nil) || (respSendSms.GetMessage() != "OK") {
		fmt.Println("send sms failed", err, respSendSms.Error())
		return false, respSendSms.GetMessage(), respond
	}
	return true, "Done", respond
}
