package Aliyuncs

import (
	"errors"
	"net/http"
)

type SMS struct {
	APPKey      string
	APPSecret   string
	PublicParam PublicParam
	SMSParam    SMSParam
}

func NewSMS(config SMSConfig) (*SMS, error) {
	if config.APPKey == "" || config.APPSecret == "" {
		return nil, errors.New("key and secret can not be empty")
	}

	publicParam := PublicParam{
		Method:     "alibaba.aliqin.fc.sms.num.send",
		APPKey:     config.APPKey,
		SignMethod: "md5",
		Format:     "json",
		V:          "2.0",
	}

	smsParam := SMSParam{
		Type:           "normal",
		FreeSignName:   config.FreeSignName,
		Param:          config.Param,
		ReceiverNumber: config.PhoneNumber,
		TemplateCode:   config.TemplateCode,
	}

	sms := &SMS{
		APPKey:      config.APPKey,
		APPSecret:   config.APPSecret,
		PublicParam: publicParam,
		SMSParam:    smsParam,
	}
	return sms, nil
}

func (sms *SMS) SendSMS() (http.Response, error) {

	return http.Response{}, nil
}
