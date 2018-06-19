package Aliyuncs

type Signature interface {
	SignTopRequest(paramMap map[string]string, secret, signMethod string) (string, error)
}
