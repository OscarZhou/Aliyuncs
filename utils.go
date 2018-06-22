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

// ExtractNotNullKeys gets the names of the struct variables
// whose values are not null.
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
			keys = append(keys, s.Type().Field(i).Name)
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
		_, ok := keyMap[s.Type().Field(i).Name]
		if !ok {
			continue
		}

		if s.Field(i).Kind() == reflect.String {
			paramMap[s.Type().Field(i).Name] = s.Field(i).String()
		}
	}
	sortedQueryString := ""
	for _, k := range keys {
		sortedQueryString += "&"
		sortedQueryString += specialUrlEncode(k)
		sortedQueryString += "="
		sortedQueryString += specialUrlEncode(paramMap[k])

	}

	return sortedQueryString, nil
}

func specialUrlEncode(src string) string {
	dst := url.QueryEscape(src)
	dst = strings.Replace(dst, "+", "%20", -1)
	dst = strings.Replace(dst, "*", "%2A", -1)
	return strings.Replace(dst, "%7E", "~", -1)
}

func EncryptHmacSha1(message, secret string) string {
	mac := hmac.New(sha1.New, []byte(secret+"&"))
	mac.Write([]byte(message))

	s := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return specialUrlEncode(s)
}
