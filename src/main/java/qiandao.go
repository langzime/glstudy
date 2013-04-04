// qiandao project main.go
package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//"regexp"
	//"strings"
	"flag"
)

var (
	name = flag.String("name", "", "The command that will be executed.")
	pwd  = flag.String("pwd", "", "The command that will be executed.")
)

func main() {
	flag.Parse()
	if len(*name) == 0 || len(*pwd) == 0 {
		fmt.Print("missing parm...")
		return
	}
	cookies := Login("wangyanqing", "wangyq")
	qiandao(cookies, "onDuty")
	qiandao(cookies, "offDuty")
	fmt.Println(*name)
}
func Login(name, pwd string) (cookies []*http.Cookie) {
	data := make(url.Values)
	data.Set("UserCode", name)
	data.Set("PassWord", pwd)
	resp, err := http.DefaultClient.PostForm("http://11.201.0.40:7008/attendance/logon.do", data)
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(string(robots))
	}
	for _, cookie := range resp.Cookies() {
		cookies = append(cookies, cookie)
	}
	return
}
func qiandao(cookies []*http.Cookie, tp string) {
	req, _ := http.NewRequest("GET", "http://11.201.0.40:7008/attendance/attendance.do?actionType="+tp, nil)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(string(robots))
	}
	return
}
