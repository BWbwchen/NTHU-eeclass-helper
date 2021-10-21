package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	EECLASS_HW_URL                = "https://eeclass.nthu.edu.tw/homework/submitList/"
	AJAX_ANCHOR_DOWNLOAD_SUBMIT   = "&_lock=homeworkId&ajaxAuth="
	AJAX_ANCHOR_DOWNLOAD_ZIP      = ""
	AJAX_AUTH_DOWNLOAD_SUBMIT_LEN = 32
)

func getAllSubmitList() ([]string, string, string) {
	client := &http.Client{}

	var ret []string
	ajaxAuthGen := ""
	ajaxAuthDownload := ""

	for i := 1; i < 5; i++ {
		req, _ := http.NewRequest("GET", EECLASS_HW_URL+HW_ID+"?page="+fmt.Sprint(i), nil)
		req.Header.Set("cookie", COOKIE)

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(`input[name="chkbox"]`).Each(func(i int, s *goquery.Selection) {
			content, _ := s.Attr("value")
			ret = append(ret, content)
		})

		if i == 1 {
			d, _ := goquery.OuterHtml(doc.Selection)
			// generate file request ajaxAuth
			k := strings.Index(d, AJAX_ANCHOR_DOWNLOAD_SUBMIT)
			start := k + len(AJAX_ANCHOR_DOWNLOAD_SUBMIT)
			ajaxAuthGen = d[start : start+AJAX_AUTH_DOWNLOAD_SUBMIT_LEN]

			// download request ajaxAuth
			j := strings.Index(d, AJAX_ANCHOR_DOWNLOAD_ZIP)
			start2 := j + len(AJAX_ANCHOR_DOWNLOAD_ZIP)
			ajaxAuthDownload = d[start2 : start2+AJAX_AUTH_DOWNLOAD_SUBMIT_LEN]
		}

		defer resp.Body.Close()
	}
	return ret, ajaxAuthGen, ajaxAuthDownload
}
