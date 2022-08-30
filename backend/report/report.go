package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/aecra/covid/email"
	"github.com/aecra/covid/object"
)

func ReportAllClock() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Minute)

	users := object.GetActiveClockUser()
	rand.Seed(time.Now().UnixNano())

	for _, user := range users {
		clock(&user)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func ReportAllHealth() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Minute)

	users := object.GetActiveHealthUser()
	rand.Seed(time.Now().UnixNano())

	for _, user := range users {
		ReportSignal(&user)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func ReportSignal(user *object.User) {
	var reportResult string
	if user.Position == "school" {
		reportResult = clock(user)
	} else {
		reportResult = health(user)
	}

	var noticeResult string
	err := email.EmailNotice(user.Email, user.Username, reportResult)
	if err != nil {
		noticeResult = err.Error()
	} else {
		noticeResult = "success"
	}

	object.AddRecord(&object.Record{
		Username:     user.Username,
		Email:        user.Email,
		Position:     user.Position,
		ReportResult: reportResult,
		NoticeResult: noticeResult,
	})
}

type ReportRes struct {
	E int      `json:"e"`
	M string   `json:"m"`
	D struct{} `json:"d"`
}

func clock(user *object.User) string {
	httpPostUrl := "https://xxcapp.xidian.edu.cn/xisuncov/wap/open-report/save"
	var SendData = []byte(`{"sfzx":1,"tw":1,"area":"陕西省 西安市 长安区","city":"西安市","province":"陕西省","address":"陕西省西安市长安区兴隆街道丁香路西安电子科技大学南校区","geo_api_info":{"type":"complete","position":{"Q":34.123646375869,"R":108.82832438151098,"lng":108.828324,"lat":34.123646},"location_type":"html5","message":"Get ipLocation failed.Get geolocation success.Convert Success.Get address success.","accuracy":79,"isConverted":"true","status":1,"addressComponent":{"citycode":"029","adcode":"610116","businessAreas":[],"neighborhoodType":"","neighborhood":"","building":"","buildingType":"","street":"雷甘路","streetNumber":"266#","country":"中国","province":"陕西省","city":"西安市","district":"长安区","township":"兴隆街道"},"formattedAddress":"陕西省西安市长安区兴隆街道丁香路西安电子科技大学南校区","roads":[],"crosses":[],"pois":[],"info":"SUCCESS"},"sfcyglq":0,"sfyzz":0,"qtqk":"","ymtys":0}`)

	request, err := http.NewRequest("POST", httpPostUrl, bytes.NewBuffer(SendData))
	if err != nil {
		return "server error: request created failed"
	}
	request.Header.Set("Host", "xxcapp.xidian.edu.cn")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Length", fmt.Sprint(len(SendData)))
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 11; POCO F2 Pro Build/RKQ1.200826.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36 MMWEBID/1230 MicroMessenger/8.0.17.2040(0x28001133) Process/toolsmp WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64")
	request.Header.Set("Origin", "https://xxcapp.xidian.edu.cn")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
	request.Header.Set("Sec-Fetch-Mode", "cors")
	request.Header.Set("Sec-Fetch-Dest", "empty")
	request.Header.Set("Referer", "https://xxcapp.xidian.edu.cn/site/ncov/xidiandailyup")
	request.Header.Set("Accept-Encoding", "utf-8")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")

	request.Header.Set("Cookie", "eai-sess="+user.Eaisess+"; UUkey="+user.Uukey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "server error: request sent failed"
	}

	defer response.Body.Close()

	var reportRes ReportRes
	err = json.NewDecoder(response.Body).Decode(&reportRes)
	if err != nil {
		return "server error: response decode failed"
	}

	return reportRes.M
}

func health(user *object.User) string {
	url := "https://xxcapp.xidian.edu.cn/ncov/wap/default/save"

	payload := strings.NewReader(user.Home)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "server error: request created failed"
	}

	req.Header.Add("Host", "xxcapp.xidian.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 11; POCO F2 Pro Build/RKQ1.200826.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36 MMWEBID/1230 MicroMessenger/8.0.17.2040(0x28001133) Process/toolsmp WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding", "utf-8")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("X-KL-Ajax-Request", "Ajax_Request")
	req.Header.Add("Content-Length", fmt.Sprint(len(user.Home)))
	req.Header.Add("Origin", "https://xxcapp.xidian.edu.cn")
	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://xxcapp.xidian.edu.cn/ncov/wap/default/index")
	req.Header.Add("Cookie", "eai-sess="+user.Eaisess+"; UUkey="+user.Uukey)
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "server error: request sent failed"
	}

	defer res.Body.Close()

	var reportRes ReportRes
	err = json.NewDecoder(res.Body).Decode(&reportRes)
	if err != nil {
		return "server error: response decode failed"
	}

	return reportRes.M
}
