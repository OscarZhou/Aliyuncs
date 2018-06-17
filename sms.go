package Aliyuncs

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"time"
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
		TimeStamp:  time.Now().Format("2016-01-01 12:00:00"),
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

func (sms *SMS) SendSMS() (int, error) {
	url := "https://eco.taobao.com/router/rest/" + sms.PublicParam.Method

	body, err := json.Marshal(sms.SMSParam)
	if err != nil {
		return 500, err
	}

	var result SMSResult
	statusCode, err := DoRequest("POST", url, body, &result)
	if err != nil {
		return statusCode, err
	}
	return 200, nil
}

func (sms *SMS) Sign() error {
	var keys []string

	s := reflect.ValueOf(sms.PublicParam)
	if s.IsValid() && s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).Kind() == reflect.String {
				if s.Field(i).String() == "" {
					continue
				}
			}
			keys = append(keys, s.Type().Field(i).Name)
		}
	}

	s = reflect.ValueOf(sms.SMSParam)
	if s.IsValid() && s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).Kind() == reflect.String {
				if s.Field(i).String() == "" {
					continue
				}
			}

			if s.Field(i).Kind() == reflect.Map {
				if s.Field(i).IsNil() {
					continue
				}
			}
			keys = append(keys, s.Type().Field(i).Name)
		}
	}

	sort.Strings(keys)

	fmt.Println(keys)
	return nil
}
