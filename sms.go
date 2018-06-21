package Aliyuncs

import (
	"errors"
	"net/url"
	"strings"
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
		TimeStamp:      time.Now().Format("2006-01-02 15:04:05"),
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

	body, err := sms.GetURLQuery()
	if err != nil {
		return 500, err
	}

	return DoRequest("POST", url, []byte(body))
}

// Sign generates the body that the http request needs
func (sms *SMS) GetURLQuery() (string, error) {
	var keys []string

	k, err := ExtractNotNullKeys(sms.SMSParam)
	if err != nil {
		return "", err
	}
	keys = append(keys, k...)

	keyMap := make(map[string]bool)
	for _, v := range keys {
		keyMap[v] = true
	}

	pMap, err := GenerateMap(sms.SMSParam, keyMap)
	if err != nil {
		return "", err
	}
	sms.SMSParam.Sign, err = SignTopRequest(pMap, sms.APPSecret, sms.SMSParam.SignMethod)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("sign", strings.ToUpper(sms.SMSParam.Sign))
	for k, v := range pMap {
		params.Set(k, v)
	}
	return params.Encode(), nil
}
