package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/aecra/covid/object"
)

type ClockRes struct {
	M string `json:"m"`
}

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "", log.LstdFlags)
)

func report() {
	// sleep a random time between 0 and 30 mins
	time.Sleep(time.Duration(rand.Intn(30)) * time.Minute)
	defer fmt.Print(&buf)

	users := object.GetActiveUser()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})

	wg := sync.WaitGroup{}
	for _, user := range users {
		if user.Position == "school" {
			wg.Add(1)
			reportSignal(user, &wg, "school")
		} else {
			wg.Add(1)
			reportSignal(user, &wg, "home")
		}
	}
	wg.Wait()
}

func reportSignal(user object.User, wg *sync.WaitGroup, position string) {
	logger.Println("clock in:", user.Name)
	err, res := clock(user, position)
	if err != nil {
		logger.Println("clock error: ", err, res)
	} else {
		logger.Println("clock in result:", res)
	}

	if strings.Contains(res, "您已上报过") {
		logger.Println("already reported")
		wg.Done()
		return
	}

	logger.Println("send email:", user.Email)
	if err := notice(user.Email, user.Name, res); err != nil {
		logger.Println("send email error:", err)
	}
	object.AddRecord(&object.Record{
		Name:     user.Name,
		Email:    user.Email,
		Position: position,
		Content:  res,
		Result:   res,
	})

	wg.Done()
}

func clock(user object.User, position string) (err error, res string) {
	httpPostUrl := "https://xxcapp.xidian.edu.cn/xisuncov/wap/open-report/save"
	var jsonData = []byte(`{"sfzx":1,"tw":1,"area":"陕西省 西安市 长安区","city":"西安市","province":"陕西省","address":"陕西省西安市长安区兴隆街道丁香路西安电子科技大学南校区","geo_api_info":{"type":"complete","position":{"Q":34.123646375869,"R":108.82832438151098,"lng":108.828324,"lat":34.123646},"location_type":"html5","message":"Get ipLocation failed.Get geolocation success.Convert Success.Get address success.","accuracy":79,"isConverted":"true","status":1,"addressComponent":{"citycode":"029","adcode":"610116","businessAreas":[],"neighborhoodType":"","neighborhood":"","building":"","buildingType":"","street":"雷甘路","streetNumber":"266#","country":"中国","province":"陕西省","city":"西安市","district":"长安区","township":"兴隆街道"},"formattedAddress":"陕西省西安市长安区兴隆街道丁香路西安电子科技大学南校区","roads":[],"crosses":[],"pois":[],"info":"SUCCESS"},"sfcyglq":0,"sfyzz":0,"qtqk":"","ymtys":0}`)

	var SendData []byte
	if position == "school" {
		SendData = jsonData
	} else {
		SendData = []byte(user.Home)
	}
	request, err := http.NewRequest("POST", httpPostUrl, bytes.NewBuffer(SendData))
	if err != nil {
		return err, ""
	}
	request.Header.Set("Host", "xxcapp.xidian.edu.cn")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Length", fmt.Sprint(len(SendData)))
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("User-Agent", "Mozilla/5.0(Linux;Android7.0;wv litebaiduboxapp)AppleWebKit/ 537.36(KHTML,likeGecko)Version/ 4.0Chrome/78.0.3904.96MobileSafari/ 537.36T7/10.3SearchCraft/2.6.2(Baidu;P17.0)")
	request.Header.Set("Origin", "https://xxcapp.xidian.edu.cn")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
	request.Header.Set("Sec-Fetch-Mode", "cors")
	request.Header.Set("Sec-Fetch-Dest", "empty")
	request.Header.Set("Referer", "https://xxcapp.xidian.edu.cn/site/ncov/xidiandailyup")
	request.Header.Set("Accept-Encoding", "utf-8")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")

	request.Header.Set("Cookie", "eai-sess="+user.Eaisess+"; UUkey="+user.Uukey)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err, "系统故障，请联系管理员！"
	}
	defer response.Body.Close()
	var clockRes ClockRes
	err = json.NewDecoder(response.Body).Decode(&clockRes)
	if err != nil || strings.Contains(clockRes.M, "用户信息已失效") {
		return err, "用户信息已失效, 请更新 Cookies, 如若无效，请联系管理员！"
	}
	return nil, string(clockRes.M)
}

func notice(email, greet, res string) (err error) {
	httpPostUrl := "https://api3.aecra.cn/email/"
	jsonData := []byte(`{
		"token": "TidtdD3c4DrM6aGM",
		"title": "疫情打卡",
		"email": "` + email + `",
		"greet": "` + greet + `",
		"content": "` + res + `"
}`)
	request, err := http.NewRequest("POST", httpPostUrl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}
