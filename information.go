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

type Student struct {
	ID       string `json:"student_id"`
	Name     string `json:"student_name"`
	SubmitID string `json:"submit_id"`
}

func getAllSubmitList() ([]string, string, string) {
	client := &http.Client{}

	var ret []string
	var data []Student
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

		doc.Find(`tr[class=" "]`).Each(func(i int, s *goquery.Selection) {
			var d Student
			s.Find(`input[name="chkbox"]`).Each(func(i int, ss *goquery.Selection) {
				content, _ := ss.Attr("value")
				d.SubmitID = content
				ret = append(ret, content)
			})
			s.Find(`td[class="    col-photoAccount"] span[class="text "]`).Each(func(i int, ss *goquery.Selection) {
				content := ss.Text()
				d.Name = content
				ret = append(ret, content)
			})
			s.Find(`div[class="fs-hint"]`).Each(func(i int, ss *goquery.Selection) {
				content := ss.Text()
				d.ID = content
				ret = append(ret, content)
			})

			data = append(data, d)
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

	generateStudentScoreCSVTemplate(data)
	return ret, ajaxAuthGen, ajaxAuthDownload
}
