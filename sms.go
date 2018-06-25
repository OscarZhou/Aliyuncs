package Aliyuncs

import (
	"errors"
	"time"

	"github.com/satori/go.uuid"
)

// Sms is a sms handler
type Sms struct {
	Product         string `json:"product"`
	Domain          string `json:"domain"`
	AccessKeyID     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	SmsParam        SmsParam
	SmsReturnStatus SmsReturnStatus
}

// NewSms creates a sms handler
func NewSms(config SmsConfig) (*Sms, error) {
	if config.AccessKeyID == "" || config.AccessKeySecret == "" ||
		len(config.PhoneNumbers) == 0 || config.SignName == "" ||
		config.TemplateCode == "" {
		return nil, errors.New("Missing configuration parameters")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	smsParam := SmsParam{
		SignatureMethod:  "HMAC-SHA1",
		SignatureNonce:   id.String(),
		AccessKeyID:      config.AccessKeyID,
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
		AccessKeyID:     config.AccessKeyID,
		AccessKeySecret: config.AccessKeySecret,
		SmsParam:        smsParam,
		SmsReturnStatus: SmsReturnStatus{},
	}
	return sms, nil
}

// SendSms sends sms request to Aliyun
func (sms *Sms) SendSms() (int, error) {
	sortedQueryString, err := GenerateSignRequestString(sms.SmsParam)
	if err != nil {
		return 500, err
	}
	stringToSign := "GET&" + SpecialURLEncode("/") + "&" + SpecialURLEncode(sortedQueryString[1:])
	sms.SmsParam.Signature = EncryptHmacSha1(stringToSign, sms.AccessKeySecret)
	url := "http://" + sms.Domain + "/?Signature=" + sms.SmsParam.Signature + sortedQueryString
	return DoRequest("GET", url, nil, &sms.SmsReturnStatus)
}
