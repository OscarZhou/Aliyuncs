package Aliyuncs

// SmsConfig stores the setting parameters of the sms API
type SmsConfig struct {
	PhoneNumbers    []string
	SignName        string
	TemplateCode    string
	TemplateParam   string
	AccessKeyID     string
	AccessKeySecret string
}
