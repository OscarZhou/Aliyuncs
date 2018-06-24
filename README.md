# Aliyuncs

Alibaba cloud communication system Golang SDK - Go 语言开源阿里云通讯SDK

The current version implements the sms sending function.  



*参考文档:*  
1. [Aliyuncs HTTP protocol and signature](https://help.aliyun.com/document_detail/56189.html?spm=a2c4g.11186623.6.581.kGwdh9)
2. [SendSms API](https://help.aliyun.com/document_detail/55284.html?spm=a2c4g.11186623.2.7.fvbDcw)

**安装**  

```
$ go get github.com/OscarZhou/Aliyuncs
```


**示例代码**  

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


 *类似SDK:* [阿里大于](https://github.com/OscarZhou/Alidayu)