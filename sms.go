package Aliyuncs

import (
	"encoding/json"
	"errors"
	"time"
)

type SMS struct {
	APPKey    string
	APPSecret string
	SMSParam  SMSParam
}

func NewSMS(config SMSConfig) (*SMS, error) {
	if config.APPKey == "" || config.APPSecret == "" {
		return nil, errors.New("key and secret can not be empty")
	}

	smsParam := SMSParam{
		Type:           "normal",
		FreeSignName:   config.FreeSignName,
		Param:          config.Param,
		ReceiverNumber: config.PhoneNumber,
		TemplateCode:   config.TemplateCode,
		Method:         "alibaba.aliqin.fc.sms.num.send",
		APPKey:         config.APPKey,
		TimeStamp:      time.Now().Format("2016-01-01 12:00:00"),
		SignMethod:     "md5",
		Format:         "json",
		V:              "2.0",
	}
	sms := &SMS{
		APPKey:    config.APPKey,
		APPSecret: config.APPSecret,
		SMSParam:  smsParam,
	}
	return sms, nil
}

func (sms *SMS) SendSMS() (int, error) {
	url := "https://eco.taobao.com/router/rest"

	err := sms.SignParams()
	if err != nil {
		return 500, err
	}

	body, err := json.Marshal(sms.SMSParam)
	if err != nil {
		return 500, err
	}
	return DoRequest("POST", url, body)
}

// Sign generates the body that the http request needs
func (sms *SMS) SignParams() error {
	var keys []string

	k, err := ExtractNotNullKeys(sms.SMSParam)
	if err != nil {
		return err
	}
	keys = append(keys, k...)

	keyMap := make(map[string]bool)
	for _, v := range keys {
		keyMap[v] = true
	}

	pMap, err := GenerateMap(sms.SMSParam, keyMap)
	if err != nil {
		return err
	}
	sms.SMSParam.Sign, err = SignTopRequest(pMap, sms.APPSecret, sms.SMSParam.SignMethod)
	if err != nil {
		return err
	}
	return nil
}
