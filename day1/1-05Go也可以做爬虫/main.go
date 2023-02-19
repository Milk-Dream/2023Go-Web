package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("开始下载...")
	resp, _ := http.Get("https://www.baidu.com/")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	ioutil.WriteFile("baidu.html", body, 0666)
	fmt.Println("下载完毕...")

	fmt.Println("开始下载图片")
	imgUrls := []string {
		"https://pics1.baidu.com/feed/c8ea15ce36d3d539e745d1965805555b342ab0fd.jpeg@f_auto?token=f7e98a988ce7f6ba3e4104ccfd026aae",
		"https://pics1.baidu.com/feed/cefc1e178a82b901084053d4130f157c3b12eff2.jpeg@f_auto?token=1c5ef5394d85f39b62aeacc128316d84",
	}

	for i, url := range imgUrls {
		imgName := strings.Join([]string{strconv.Itoa(i), "jpg"}, ".")
		fmt.Println("开始下载:", imgName)
		resp, _ := http.Get(url)
		content, _ := ioutil.ReadAll(resp.Body)
		ioutil.WriteFile(imgName, content, 0666)
		fmt.Println("下载完毕:", imgName)
	}
}