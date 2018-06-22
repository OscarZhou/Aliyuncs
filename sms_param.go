package Aliyuncs

type SmsParam struct {
	AccessKeyId      string
	Timestamp        string
	Format           string
	SignatureMethod  string
	SignatureVersion string
	SignatureNonce   string
	Signature        string
	Action           string
	Version          string
	RegionID         string
	PhoneNumbers     string
	SignName         string
	TemplateCode     string
	TemplateParam    string
	OutID            string
}
