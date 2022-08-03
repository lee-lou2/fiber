package http

import (
	"io"
	"io/ioutil"
	netHttp "net/http"
)

type ObjectType struct {
	Key   string
	Value string
}

func Request(method string, url string, payload io.Reader, headers ...*ObjectType) (string, error) {
	/*
		HTTP 요청
	*/
	client := &netHttp.Client{}

	req, err := netHttp.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}

	// 헤더 값 지정
	if len(headers) > 0 {
		for _, header := range headers {
			req.Header.Set(header.Key, header.Value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
