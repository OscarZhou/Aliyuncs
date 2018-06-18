package Aliyuncs

import "reflect"

type SMSParam struct {
	Extend         string
	Type           string
	FreeSignName   string
	Param          map[string]string
	ReceiverNumber string
	TemplateCode   string
}

func (sp SMSParam) ExtractKeys() ([]string, error) {
	var keys []string
	s := reflect.ValueOf(sp)
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
	}
	return keys, nil
}
