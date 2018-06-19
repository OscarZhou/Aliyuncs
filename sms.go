package Aliyuncs

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
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

func (sms *SMS) SignTopRequest(paramMap map[string]string, secret, signMethod string) (string, error) {
	var (
		keys      []string
		encrypted []byte
		err       error
	)
	for k, _ := range paramMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	query := ""
	if signMethod == "md5" {
		query += secret
	}
	for _, k := range keys {
		query += k
		query += paramMap[k]
	}

	if signMethod == "hmac" {

	} else {
		query += secret
		encrypted, err = encryptMD5(query)
		if err != nil {
			return "", err
		}
	}
	return hex.EncodeToString(encrypted), nil
}

func (sms *SMS) Sign() error {
	var keys []string
	k, err := ExtractKeys(sms.PublicParam)
	if err != nil {
		return err
	}
	keys = append(keys, k...)

	k, err = ExtractKeys(sms.SMSParam)
	if err != nil {
		return err
	}
	keys = append(keys, k...)

	sort.Strings(keys)

	keyMap := make(map[string]bool)
	for _, v := range keys {
		keyMap[v] = true
	}

	paramMap := make(map[string]string)
	pMap, err := GenerateMap(sms.PublicParam, keyMap)
	if err != nil {
		return err
	}

	for k, v := range pMap {
		paramMap[k] = v
	}

	pMap, err = GenerateMap(sms.SMSParam, keyMap)
	if err != nil {
		return err
	}

	for k, v := range pMap {
		paramMap[k] = v
	}

	plainText := sms.APPSecret
	for _, k := range keys {
		v, ok := paramMap[k]
		if ok {
			plainText += k
			plainText += v
		}
	}
	plainText += sms.APPSecret
	if sms.PublicParam.SignMethod == "md5" {

	}

	fmt.Println(paramMap)

	return nil
}

func ExtractKeys(params interface{}) ([]string, error) {
	var keys []string
	s := reflect.ValueOf(params)
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
	} else {
		return nil, errors.New("param is invalid")
	}

	return keys, nil
}

func GenerateMap(params interface{}, keyMap map[string]bool) (map[string]string, error) {
	paramMap := make(map[string]string)
	pp := reflect.ValueOf(params)
	if pp.IsValid() {
		for i := 0; i < pp.NumField(); i++ {
			_, ok := keyMap[pp.Type().Field(i).Name]
			if !ok {
				continue
			}
			switch pp.Field(i).Kind() {
			case reflect.String:
				paramMap[pp.Type().Field(i).Name] = pp.Field(i).String()
			case reflect.Bool:
				paramMap[pp.Type().Field(i).Name] = strconv.FormatBool(pp.Field(i).Bool())
			case reflect.Map:
				m, ok := pp.Field(i).Interface().(map[string]string)
				if !ok {
					return nil, errors.New("assert failure")
				}
				v, err := json.Marshal(m)
				if err != nil {
					return nil, err
				}
				paramMap[pp.Type().Field(i).Name] = string(v)
			}
		}
	} else {
		return nil, errors.New("params is invalid")
	}

	return paramMap, nil
}

func encryptMD5(plainText string) ([]byte, error) {
	if plainText == "" {
		return nil, errors.New("plain text can not be empty")
	}

	h := md5.New()
	io.WriteString(h, plainText)
	return h.Sum(nil), nil
}

func encryptHMAC(plainText, secret string) ([]byte, error) {
	return nil, nil
}
