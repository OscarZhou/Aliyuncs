package Aliyuncs

// SmsReturnStatus presents sms response information
type SmsReturnStatus struct {
	RequestID string `json:"RequestId"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	BizID     string `json:"BizId"`
}

// LookupCode returns a Chinese description about Code
func (smsr *SmsReturnStatus) LookupCode() string {
	desc := ""
	switch smsr.Code {
	case "OK":
		desc = "请求成功"
	case "isp.RAM_PERMISSION_DENY":
		desc = "RAM权限DENY"
	case "isv.OUT_OF_SERVICE":
		desc = "业务停机"
	case "isv.PRODUCT_UN_SUBSCRIPT":
		desc = "未开通云通信产品的阿里云客户"
	case "isv.PRODUCT_UNSUBSCRIBE":
		desc = "产品未开通"
	case "isv.ACCOUNT_NOT_EXISTS":
		desc = "账户不存在"
	case "isv.ACCOUNT_ABNORMAL":
		desc = "账户异常"
	case "isv.SMS_TEMPLATE_ILLEGAL":
		desc = "短信模板不合法"
	case "isv.SMS_SIGNATURE_ILLEGAL":
		desc = "短信签名不合法"
	case "isv.INVALID_PARAMETERS":
		desc = "参数异常"
	case "isp.SYSTEM_ERROR":
		desc = "系统错误"
	case "isv.MOBILE_NUMBER_ILLEGAL":
		desc = "非法手机号"
	case "isv.MOBILE_COUNT_OVER_LIMIT":
		desc = "手机号码数量超过限制"
	case "isv.TEMPLATE_MISSING_PARAMETERS":
		desc = "模板缺少变量"
	case "isv.BUSINESS_LIMIT_CONTROL":
		desc = "业务限流"
	case "isv.INVALID_JSON_PARAM":
		desc = "JSON参数不合法，只接受字符串值"
	case "isv.BLACK_KEY_CONTROL_LIMIT":
		desc = "黑名单管控"
	case "isv.PARAM_LENGTH_LIMIT":
		desc = "参数超出长度限制"
	case "isv.PARAM_NOT_SUPPORT_URL":
		desc = "不支持URL"
	case "isv.AMOUNT_NOT_ENOUGH":
		desc = "账户余额不足"
	}
	return desc
}
