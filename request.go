package Aliyuncs

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// DoRequest sends http request
func DoRequest(method, url string, body []byte) (int, error) {
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

	return resp.StatusCode, nil
}
