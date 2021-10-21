package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

var (
	EECLASS_BASE_URL            = "https://eeclass.nthu.edu.tw/dashboard"
	EECLASS_DOWNLOAD_SUBMIT_URL = "https://eeclass.nthu.edu.tw/ajax/sys.pages.homework_submitList/package/?homeworkId="
	EECLASS_DOWNLOAD_ZIP_URL    = ""
)

func downloadAllSubmit(allList []string, ajaxAuthGen string, ajaxAuthDownload string) {
	body := ""
	for _, s := range allList {
		body += "reportIds%5B%5D=" + s + "&"
	}
	l := len(body)
	body = body[:l-1]

	resp, err := sendRequest("POST", EECLASS_DOWNLOAD_SUBMIT_URL+HW_ID+AJAX_ANCHOR_DOWNLOAD_SUBMIT+ajaxAuthGen, strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	b, _ := io.ReadAll(resp.Body)

	type Rr struct {
		Status   bool   `json:"status"`
		Msg      string `json:"msg"`
		Focus    string `json:"focus"`
		ErrorMsg string `json:"errorMsg"`
	}

	type Dd struct {
		Ret Rr `json:"ret"`
	}

	var data Dd
	json.Unmarshal(b, &data)

	url := EECLASS_DOWNLOAD_ZIP_URL + ajaxAuthDownload + "&file=" + data.Ret.Msg
	download(url, fmt.Sprint(HW_ID)+".zip")
}

func download(url string, fileName string) {
	resp, err := sendRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(fileName, []byte(body), 0644)
}
