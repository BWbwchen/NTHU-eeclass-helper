package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	GET_REPORT_PAGE_AJAX_ANCHOR = "/ajax/sys.pages.homework_submitList/auditReport/?homeworkId=%s&reportId=%s&_lock=homeworkId%%2CreportId&ajaxAuth="
	SCORE_REQ_BODY_TEMPLATE     = "_fmSubmit=yes&formVer=3.0&formId=homework-audit-setting-form&auditScore=%d&worthWatch=&auditNote=%%3Cdiv%%3E%s%%3C%%2Fdiv%%3E%%0A&keepAudit=0&nextNotAuditReportId=&reportId=%s&csrf-t=%s"
)

func sendScore() {
	csv, err := ReadCsv("hw_" + HW_ID + ".csv")
	if err != nil {
		panic(err)
	}

	studentsInfo := stringToStudentStruct(csv)
	ajaxReportMap := getReportPageAjax(studentsInfo)

	type Rr struct {
		Status    bool   `json:"status"`
		Msg       string `json:"msg"`
		Focus     string `json:"focus"`
		TargetUrl string `json:"targetUrl"`
	}

	type Dd struct {
		Ret Rr `json:"ret"`
	}

	for _, student := range studentsInfo {
		//get send report page ajaxAuth first
		url := fmt.Sprintf(GET_REPORT_PAGE_AJAX_ANCHOR, HW_ID, student.SubmitID) + ajaxReportMap[student.SubmitID]
		resp, err := sendRequest("POST", BASE_URL+url, nil)
		if err != nil {
			panic(err)
		}
		respString := respBodyToString(resp.Body)

		var data Dd
		json.Unmarshal([]byte(respString), &data)

		resp, _ = sendRequest("GET", BASE_URL+data.Ret.TargetUrl, nil)

		doc, _ := goquery.NewDocumentFromReader(resp.Body)

		csrf := ""
		sendScorePath := ""
		doc.Find(`input[name="csrf-t"]`).Each(func(i int, s *goquery.Selection) {
			csrf, _ = s.Attr("value")
		})

		doc.Find(`form[id="homework-audit-setting-form"]`).Each(func(i int, s *goquery.Selection) {
			sendScorePath, _ = s.Attr("action")
		})
		if student.Score >= 0 && student.Score <= 100 {
			body := fmt.Sprintf(SCORE_REQ_BODY_TEMPLATE, student.Score, student.Comment, student.SubmitID, csrf)
			sendRequest("POST", BASE_URL+sendScorePath, strings.NewReader(body))
			fmt.Println("Send one score success!")
		}
	}
}

func getReportPageAjax(studentsInfo []Student) map[string]string {
	content := ""
	for i := 1; i < 5; i++ {
		resp, err := sendRequest("GET", EECLASS_HW_URL+HW_ID+"?page="+fmt.Sprint(i), nil)
		if err != nil {
			panic(err)
		}
		content += respBodyToString(resp.Body) + "\n"
	}

	ret := make(map[string]string)

	for _, student := range studentsInfo {
		anchor := fmt.Sprintf(GET_REPORT_PAGE_AJAX_ANCHOR, HW_ID, student.SubmitID)
		k := strings.Index(content, anchor)
		start := k + len(anchor)
		ajaxAuthGen := content[start : start+AJAX_AUTH_DOWNLOAD_SUBMIT_LEN]
		ret[student.SubmitID] = ajaxAuthGen
	}
	return ret
}

func respBodyToString(body io.Reader) string {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}

func stringToStudentStruct(csv [][]string) []Student {
	var ret []Student
	for _, row := range csv[1:] {
		var score int
		if row[3] == "" {
			score = -1
		} else {
			var err error
			score, err = strconv.Atoi(row[3])
			if err != nil {
				panic(err)
			}
		}

		if score > 100 {
			panic("invalid score")
		}

		d := Student{
			ID:       row[0],
			Name:     row[1],
			SubmitID: row[2],
			Score:    score,
			Comment:  row[4],
		}
		ret = append(ret, d)
	}
	return ret
}
