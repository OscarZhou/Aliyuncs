package Aliyuncs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Request interface {
	DoRequest(method, url string, body []byte, object interface{}) (int, error)
}

func DoRequest(method, url string, body []byte, object interface{}) (int, error) {
	var reader io.Reader
	switch method {
	case "GET", "get":

	case "POST", "post":
		if len(body) == 0 {
			return 411, errors.New("Length required: body. ")
		}
		reader = bytes.NewReader(body)
	default:
		return 400, errors.New("http method invalid. ")
	}

	fmt.Println(reader)
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return 500, errors.New("Internal server error")
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if resp == nil {
		return 500, errors.New("Internal server error. return null. ")
	}

	if err != nil {
		return resp.StatusCode, errors.New(resp.Status)
	}

	defer resp.Body.Close()
	if resp.StatusCode > 300 && resp.StatusCode < 200 {
		return resp.StatusCode, errors.New(resp.Status)
	}

	respBody, err := ioutil.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return resp.StatusCode, errors.New(resp.Status)
	}

	fmt.Println(string(respBody))

	if !reflect.ValueOf(object).IsValid() {
		return 400, errors.New("the object is invalid ")
	}

	err = json.Unmarshal(respBody, object)
	if err != nil {
		return 500, errors.New("Unmarshal error")
	}

	return resp.StatusCode, nil

}
