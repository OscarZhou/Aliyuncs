package Aliyuncs

import (
	"testing"
)

func TestSms(t *testing.T) {
	smsConfig := SmsConfig{
		AccessKeyID:     "",
		AccessKeySecret: "",
		PhoneNumbers:    "15000000000",
		SignName:        "云通信(短信签名)",
		TemplateCode:    "SMS_0000(短信模板ID)",
		TemplateParam:   `{"code":"1234","product":"ytx"}`,
	}
	sms, err := NewSms(smsConfig)
	if err != nil {
		t.Error(err)
	}

	statusCode, err := sms.SendSms()
	if err != nil {
		t.Errorf("status code is %d, error is %s\n", statusCode, err)
	}
}
