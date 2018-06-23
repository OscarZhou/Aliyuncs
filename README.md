# Aliyuncs

Alibaba cloud communication system APIs implementing three primary cloud communication products: *sms*, *voice* and *flow*.  
The current version implements the sms sending function.  

*Related API document:*  
1. [Aliyuncs HTTP protocol and signature](https://help.aliyun.com/document_detail/56189.html?spm=a2c4g.11186623.6.581.kGwdh9)
2. [SendSms API](https://help.aliyun.com/document_detail/55284.html?spm=a2c4g.11186623.2.7.fvbDcw)

**Example Code:**  

```
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

 ```


 *Similar API:* [Alidayu API](https://github.com/OscarZhou/Alidayu)