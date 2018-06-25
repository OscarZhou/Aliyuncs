package Aliyuncs

// SmsParam contains system and business parameters
type SmsParam struct {
	AccessKeyID      string   `json:"AccessKeyId"`
	Timestamp        string   `json:"Timestamp"`
	Format           string   `json:"Format"`
	SignatureMethod  string   `json:"SignatureMethod"`
	SignatureVersion string   `json:"SignatureVersion"`
	SignatureNonce   string   `json:"SignatureNonce"`
	Signature        string   `json:"Signature"`
	Action           string   `json:"Action"`
	Version          string   `json:"Version"`
	RegionID         string   `json:"RegionID"`
	PhoneNumbers     []string `json:"PhoneNumbers"`
	SignName         string   `json:"SignName"`
	TemplateCode     string   `json:"TemplateCode"`
	TemplateParam    string   `json:"TemplateParam"`
	OutID            string   `json:"OutID"`
}
