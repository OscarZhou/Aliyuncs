package Aliyuncs

type SMSParam struct {
	Extend         string
	Type           string
	FreeSignName   string
	Param          map[string]string
	ReceiverNumber string
	TemplateCode   string
	ExtraParam     map[string]interface{}
}
