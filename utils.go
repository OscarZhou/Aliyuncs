package Aliyuncs

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// SignTopRequest signs the params
func SignTopRequest(paramMap map[string]string, secret, signMethod string) (string, error) {
	var (
		keys      []string
		encrypted []byte
		err       error
	)
	for k := range paramMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	query := ""
	if signMethod == "md5" {
		query += secret
	}
	for _, k := range keys {
		query += strings.ToLower(k)
		query += paramMap[k]
	}
	if signMethod == "hmac" {
		encrypted, err = encryptHMAC(query, secret)
	} else {
		query += secret
		encrypted, err = encryptMD5(query)
		if err != nil {
			return "", err
		}
	}
	return hex.EncodeToString(encrypted), nil
}

// ExtractNotNullKeys gets the names of the struct variables
// whose values are not null.
func ExtractNotNullKeys(params interface{}) ([]string, error) {
	var keys []string
	s := reflect.ValueOf(params)
	if s.IsValid() && s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).Kind() == reflect.String {
				if s.Field(i).String() == "" {
					continue
				}
			}

			if s.Field(i).Kind() == reflect.Map {
				if s.Field(i).IsNil() {
					continue
				}
			}
			keys = append(keys, s.Type().Field(i).Name)
		}
	} else {
		return nil, errors.New("param is invalid")
	}

	return keys, nil
}

// GenerateMap generates a new map[string]string variable
func GenerateMap(params interface{}, keyMap map[string]bool) (map[string]string, error) {
	paramMap := make(map[string]string)
	pp := reflect.ValueOf(params)
	if pp.IsValid() {
		for i := 0; i < pp.NumField(); i++ {
			_, ok := keyMap[pp.Type().Field(i).Name]
			if !ok {
				continue
			}
			switch pp.Field(i).Kind() {
			case reflect.String:
				paramMap[pp.Type().Field(i).Tag.Get("json")] = pp.Field(i).String()
			case reflect.Bool:
				paramMap[pp.Type().Field(i).Tag.Get("json")] = strconv.FormatBool(pp.Field(i).Bool())
			case reflect.Map:
				m, ok := pp.Field(i).Interface().(map[string]string)
				if !ok {
					return nil, errors.New("assert failure")
				}
				v, err := json.Marshal(m)
				if err != nil {
					return nil, err
				}
				paramMap[pp.Type().Field(i).Tag.Get("json")] = string(v)
			}
		}
	} else {
		return nil, errors.New("params is invalid")
	}

	return paramMap, nil
}

// encryptMD5 encrypts message with MD5
func encryptMD5(plainText string) ([]byte, error) {
	if plainText == "" {
		return nil, errors.New("plain text can not be empty")
	}

	h := md5.New()
	io.WriteString(h, plainText)
	return h.Sum(nil), nil
}

// encryptHMAC encrypts message with hmac
func encryptHMAC(plainText, secret string) ([]byte, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(plainText))
	return mac.Sum(nil), nil
}
