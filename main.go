package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	EECLASS_LOGIN_URL  = "https://oauth.ccxp.nthu.edu.tw/v1.1/authorize.php?response_type=code&client_id=eeclass&redirect_uri=https%3A%2F%2Feeclass.nthu.edu.tw%2Fservice%2Foauth%2F&scope=lmsid+userid&state=&ui_locales=zh-CH"
	EECLASS_LOGIN_TYPE = "application/x-www-form-urlencoded"
	HW_ID              = ""
	COOKIE             = ""
)

func init() {
	godotenv.Load("env")
	loadGlobalVariable()
}

func loadGlobalVariable() {
	HW_ID = os.Getenv("COURSE_ID")
	COOKIE = os.Getenv("COOKIE")
	EECLASS_DOWNLOAD_ZIP_URL = "https://eeclass.nthu.edu.tw/homework/package/" + HW_ID + "/?ajaxAuth="
	AJAX_ANCHOR_DOWNLOAD_ZIP = "/homework/package/" + HW_ID + "/?ajaxAuth="
}

func main() {
	os.Exit(mainLoop())
}

const (
	DOWNLOAD_SUBMIT_FILE = iota
	UPLOAD_SCORE
)

func mainLoop() int {
	err := promptInputString(func(in interface{}) error {
		s := in.(string)
		if s != "" {
			HW_ID = s
		}
		return nil
	})
	if err != nil {
		return 1
	}

	jobID, err := promptSelectJob()
	if err != nil {
		return 1
	}

	switch jobID {
	case DOWNLOAD_SUBMIT_FILE:
		downloadAllSubmit(getAllSubmitList())
	case UPLOAD_SCORE:
		fmt.Println("Working now...")
	default:
		return 1
	}
	return 0
}
