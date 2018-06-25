# Aliyuncs

Alibaba cloud communication system Golang SDK - Go 语言开源阿里云通讯SDK

**To-do list:**  

:white_check_mark: 短信发送（包括群发，群发上限为1000）
:black_square_button: 短信查询
:black_square_button: 语音
:black_square_button: 流量




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
		PhoneNumbers:    []string{"15000000000", "15000000001"},
		SignName:        "云通信(短信签名)",
		TemplateCode:    "SMS_0000(短信模板ID)",
		TemplateParam:   `{"code":"1234","product":"ytx"}`,
	}
	sms, err := NewSms(smsConfig)
	if err != nil {
		fmt.Errorf(err)
	}

	statusCode, err := sms.SendSms()
	if err != nil {
		fmt.Errorf("status code: %d, error: %s\n", statusCode, err.Error())
		fmt.Errorf("Aliyun error code: %s, description: %s\n", sms.SmsReturnStatus.Code, sms.SmsReturnStatus.Message)
		fmt.Errorf("code description: %s\n", sms.SmsReturnStatus.LookupCode())
	}

 ```


 *类似SDK:* [阿里大于](https://github.com/OscarZhou/Alidayu)