# NTHU eeclass TA helper 

[<img alt="github" src="https://img.shields.io/badge/github-NTHU--eeclass--helper-blue?style=for-the-badge&logo=github" height="20">](https://github.com/BWbwchen/NTHU-eeclass-helper)

This tool is a CLI-interactive tool for TA who use eeclass platform. It helps TA to download all the submitted homeworks, and use CSV to record the score and homework comment, and upload CSV score directly to the eeclass platform with just 2 Enter key!

## How to use this tool
1. Login from your browser.
2. Open console (press F12, if you use Chrome)
3. Click `Application` tab in the developer tool.
4. In the left side bar, find `Storage > Cookies`, Copy the name and value of **`PHPSESSID`**
5. Paste them into `env` file as :
```
COOKIE=PHPSESSID={yourphpsessidhere};
```
example :
```
COOKIE=PHPSESSID=12341234123412341234123412341234;
```
6. Click `console` tab in the developer tool. Type in `console.log(document.cookie);`, and copy the output.
7.  Paste them into `env` file as :
```
COOKIE=PHPSESSID={yourphpsessidhere}; {your cookie}
```
example :
```
COOKIE=PHPSESSID=12341234123412341234123412341234; _login_token_=asdfasdef5asdfasdfasdf75asdf0e6; timezone=%2B0800; noteFontSize=100; noteExpand=0; TS01e4fe74=0140e1c48e3f8cad7563959dc63c589fd4dacf00f7e771606b3dbf2a115easdfasdfasdfasdf75a1easdfasdfasdfasdfasdfasdf2c96270a62ea1b1f1d624ca09a6d7asdfasdfasdf57a67474d57f03f643537207bbfef588d12f28a12ac1c57bc50322bb04e8500d0467d65c85d34d3a95f9
```
8. Find your course id in homework page in the browser. Example :
```
https://eeclass.nthu.edu.tw/course/homework/2046
```
2046 is the course id. So fill in the `env` file.
```
COURSE_ID=2046
```

## Run this tool from this repository
Run the below command, and it follow the CLI-interactive prompt to download the submitted file in the root directory of this project, or upload the score.
```
go run .
```

## Run this tool from release
Download the binary from release page, and then run it directly with `env` file in the same directory.

## Goal
- [x] Input score from a csv file
- [ ] Auto Login