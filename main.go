package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("env")
	loadGlobalVariable()
}

func loadGlobalVariable() {
	HW_ID = os.Getenv("COURSE_ID")
	COOKIE = os.Getenv("COOKIE")
	EECLASS_DOWNLOAD_ZIP_URL = "https://eeclass.nthu.edu.tw/homework/package/" + fmt.Sprint(HW_ID) + "/?ajaxAuth="
	AJAX_ANCHOR_DOWNLOAD_ZIP = "/homework/package/" + fmt.Sprint(HW_ID) + "/?ajaxAuth="
}

var (
	EECLASS_LOGIN_URL             = "https://oauth.ccxp.nthu.edu.tw/v1.1/authorize.php?response_type=code&client_id=eeclass&redirect_uri=https%3A%2F%2Feeclass.nthu.edu.tw%2Fservice%2Foauth%2F&scope=lmsid+userid&state=&ui_locales=zh-CH"
	EECLASS_BASE_URL              = "https://eeclass.nthu.edu.tw/dashboard"
	EECLASS_HW_URL                = "https://eeclass.nthu.edu.tw/homework/submitList/"
	EECLASS_DOWNLOAD_SUBMIT_URL   = "https://eeclass.nthu.edu.tw/ajax/sys.pages.homework_submitList/package/?homeworkId="
	EECLASS_DOWNLOAD_ZIP_URL      = ""
	EECLASS_LOGIN_TYPE            = "application/x-www-form-urlencoded"
	AJAX_ANCHOR_DOWNLOAD_SUBMIT   = "&_lock=homeworkId&ajaxAuth="
	AJAX_ANCHOR_DOWNLOAD_ZIP      = ""
	AJAX_AUTH_DOWNLOAD_SUBMIT_LEN = 32
	HW_ID                         = ""
	COOKIE                        = ""
)

func main() {
	allList, ajaxAuthGen, ajaxAuthDownload := getAllSubmitList()
	fmt.Println(allList)
	downloadAllSubmit(allList, ajaxAuthGen, ajaxAuthDownload)
}

func getAllSubmitList() ([]string, string, string) {
	client := &http.Client{}

	var ret []string
	ajaxAuthGen := ""
	ajaxAuthDownload := ""

	for i := 1; i < 5; i++ {
		req, _ := http.NewRequest("GET", EECLASS_HW_URL+fmt.Sprint(HW_ID)+"?page="+fmt.Sprint(i), nil)
		req.Header.Set("cookie", COOKIE)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("QQQQQQQQQ")
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
			fmt.Println(d[k : k+100])
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
