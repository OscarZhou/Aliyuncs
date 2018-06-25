package Aliyuncs

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

// GenerateSignRequestString generates signature string
func GenerateSignRequestString(params interface{}) (string, error) {
	var keys []string
	s := reflect.ValueOf(params)
	if s.IsValid() && s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).Kind() == reflect.String {
				if s.Field(i).String() == "" {
					continue
				}
			}
			keys = append(keys, s.Type().Field(i).Tag.Get("json"))
		}
	} else {
		return "", errors.New("param is invalid")
	}

	sort.Strings(keys)
	keyMap := make(map[string]bool)
	for _, v := range keys {
		keyMap[v] = true
	}

	paramMap := make(map[string]string)
	for i := 0; i < s.NumField(); i++ {
		_, ok := keyMap[s.Type().Field(i).Tag.Get("json")]
		if !ok {
			continue
		}

		if s.Field(i).Kind() == reflect.String {
			paramMap[s.Type().Field(i).Tag.Get("json")] = s.Field(i).String()
		}

		if s.Field(i).Kind() == reflect.Slice {
			v, ok := s.Field(i).Interface().([]string)
			if !ok {
				return "", errors.New("phone numbers assert error")
			}

			paramMap[s.Type().Field(i).Tag.Get("json")] = strings.Join(v, ",")
		}
	}

	sortedQueryString := ""
	for _, k := range keys {
		sortedQueryString += "&"
		sortedQueryString += SpecialURLEncode(k)
		sortedQueryString += "="
		sortedQueryString += SpecialURLEncode(paramMap[k])

	}
	return sortedQueryString, nil
}

// SpecialURLEncode encodes the string with the rule of
// the special URL encode of Aliyun
func SpecialURLEncode(src string) string {
	dst := url.QueryEscape(src)
	dst = strings.Replace(dst, "+", "%20", -1)
	dst = strings.Replace(dst, "*", "%2A", -1)
	return strings.Replace(dst, "%7E", "~", -1)
}

// EncryptHmacSha1 encrypts the string with HMAC-SHA1 algorithm
func EncryptHmacSha1(message, secret string) string {
	mac := hmac.New(sha1.New, []byte(secret+"&"))
	mac.Write([]byte(message))

	s := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return SpecialURLEncode(s)
}
