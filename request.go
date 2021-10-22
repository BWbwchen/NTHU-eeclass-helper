package main

import (
	"io"
	"net/http"
)

func sendRequest(method string, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("cookie", COOKIE)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(req)
}
