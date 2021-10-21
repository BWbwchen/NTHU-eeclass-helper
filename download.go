package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadAllSubmit(allList []string, ajaxAuthGen string, ajaxAuthDownload string) {
	body := ""
	for _, s := range allList {
		body += "reportIds%5B%5D=" + s + "&"
	}
	l := len(body)
	body = body[:l-1]
	client := &http.Client{}

	req, _ := http.NewRequest("POST", EECLASS_DOWNLOAD_SUBMIT_URL+fmt.Sprint(HW_ID)+AJAX_ANCHOR_DOWNLOAD_SUBMIT+ajaxAuthGen, strings.NewReader(body))
	req.Header.Set("cookie", COOKIE)
	resp, _ := client.Do(req)

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
	fmt.Println(url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("cookie", COOKIE)
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	f, e := os.Create(fileName)
	if e != nil {
		panic(e)
	}
	fmt.Println(resp.ContentLength)
	fmt.Println(resp.Header)
	fmt.Println(resp.Request.Header)
	defer f.Close()
	io.Copy(f, resp.Body)
}
