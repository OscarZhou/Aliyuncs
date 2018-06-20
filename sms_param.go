package Aliyuncs

type SMSParam struct {
	Extend         string
	Type           string
	FreeSignName   string
	Param          map[string]string
	ReceiverNumber string
	TemplateCode   string
	Method         string
	APPKey         string
	TargetAPPKey   string
	SignMethod     string
	Sign           string
	Session        string
	TimeStamp      string
	Format         string
	V              string
	PartnerID      string
	Simplify       bool
}
