package Aliyuncs

type SMSConfig struct {
	APPKey    string
	APPSecret string

	// The sms signature. The value must be as same
	// as the accessible signature in alidayu management
	// center. The sample: 阿里大于. Required field
	FreeSignName string

	// The sms template parameters. key-value format.
	// The key must be as same as ones in the template.
	// The sample: {"code":"1234","product":"alidayu"}
	// Optional field.
	Param map[string]string

	// Phone number. The samle: 13000000000.
	// Required field.
	PhoneNumber string

	// Sms template ID. The value must be as same as
	// the accessible signature in alidayu management
	// center. The sample: SMS_585014
	TemplateCode string
}
