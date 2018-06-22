package Aliyuncs

import (
	"errors"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

type Sms struct {
	Product         string `json:"product"`
	Domain          string `json:"domain"`
	AccessKeyId     string
	AccessKeySecret string
	SmsParam        SmsParam
}

func NewSms(config SmsConfig) (*Sms, error) {
	if config.AccessKeyId == "" || config.AccessKeySecret == "" {
		return nil, errors.New("key or secret can not be empty")
	}

	// loc, err := time.LoadLocation("Asia/Shanghai")
	// if err != nil {
	// 	return nil, err
	// }
	// now := time.Now().In(loc).Format("2006-01-02T15:04:05Z")

	smsParam := SmsParam{
		SignatureMethod:  "HMAC-SHA1",
		SignatureNonce:   uuid.NewV4().String(),
		AccessKeyId:      config.AccessKeyId,
		SignatureVersion: "1.0",
		Timestamp:        time.Now().In(time.UTC).Format("2006-01-02T15:04:05Z"),
		Format:           "JSON",
		Action:           "SendSms",
		Version:          "2017-05-25",
		RegionID:         "cn-hangzhou",
		PhoneNumbers:     config.PhoneNumbers,
		SignName:         config.SignName,
		TemplateParam:    config.TemplateParam,
		TemplateCode:     config.TemplateCode,
		OutID:            "123",
	}

	sms := &Sms{
		Product:         "Dysmsapi",
		Domain:          "dysmsapi.aliyuncs.com",
		AccessKeyId:     config.AccessKeyId,
		AccessKeySecret: config.AccessKeySecret,
		SmsParam:        smsParam,
	}
	return sms, nil
}

func (sms *Sms) SendSms() (int, error) {
	sortedQueryString, err := GenerateSignRequestString(sms.SmsParam)
	if err != nil {
		return 500, err
	}
	stringToSign := "GET&" + specialUrlEncode("/") + "&" + specialUrlEncode(sortedQueryString[1:])
	sms.SmsParam.Signature = EncryptHmacSha1(stringToSign, sms.AccessKeySecret)
	fmt.Println(sms.SmsParam.Signature)
	url := "http://" + sms.Domain + "/?Signature=" + sms.SmsParam.Signature + sortedQueryString
	return DoRequest("GET", url, nil)
}
