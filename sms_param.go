package Aliyuncs

type SMSParam struct {
	Extend         string            `json:"extend"`
	Type           string            `json:"sms_type"`
	FreeSignName   string            `json:"sms_free_sign_name"`
	Param          map[string]string `json:"sms_param"`
	ReceiverNumber string            `json:"rec_num"`
	TemplateCode   string            `json:"sms_template_code"`
	Method         string            `json:"method"`
	APPKey         string            `json:"app_key"`
	TargetAPPKey   string            `json:"target_app_key"`
	SignMethod     string            `json:"sign_method"`
	Sign           string            `json:"sign"`
	Session        string            `json:"session"`
	TimeStamp      string            `json:"timestamp"`
	Format         string            `json:"format"`
	V              string            `json:"v"`
	PartnerID      string            `json:"partner_id"`
	Simplify       bool              `json:"simplify"`
}
