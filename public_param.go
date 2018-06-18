package Aliyuncs

import "reflect"

type PublicParam struct {
	Method       string
	APPKey       string
	TargetAPPKey string
	SignMethod   string
	Sign         string
	Session      string
	TimeStamp    string
	Format       string
	V            string
	PartnerID    string
	Simplify     bool
}

func (pp PublicParam) ExtractKeys() ([]string, error) {
	var keys []string
	s := reflect.ValueOf(pp)
	if s.IsValid() && s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).Kind() == reflect.String {
				if s.Field(i).String() == "" {
					continue
				}
			}
			keys = append(keys, s.Type().Field(i).Name)
		}
	}
	return keys, nil
}
