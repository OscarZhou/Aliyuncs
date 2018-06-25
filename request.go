package Aliyuncs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

// DoRequest sends http request
func DoRequest(method, url string, body []byte, response interface{}) (int, error) {
	var reader io.Reader
	if reflect.ValueOf(response).Kind() != reflect.Ptr {
		return 400, errors.New("response object should be pointer")
	}
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

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return 500, errors.New("Internal server error")
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if resp == nil {
		return 500, errors.New("Internal server error. return null. ")
	}

	if err != nil {
		return resp.StatusCode, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 300 || resp.StatusCode < 200 {
		return resp.StatusCode, errors.New(resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, errors.New(resp.Status + "," + string(respBody))
	}

	err = json.Unmarshal(respBody, response)
	if err != nil {
		return 500, err
	}

	return resp.StatusCode, nil
}
