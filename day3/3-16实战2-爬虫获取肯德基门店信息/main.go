package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	cists := []string{"北京", "上海"}
	for _, c := range(cists) {
		//http://www.kfc.com.cn/kfccda/storelist/index.aspx
		c_name := fmt.Sprintf("cname=&pid&keyword=%v&pageIndex=1&pageSize=30", c)
		resp, _ := http.Post(
			"http://www.kfc.com.cn/kfccda/ashx/GetStoreList.ashx?op=keyword",
			"application/x-www-form-urlencoded",
			strings.NewReader(c_name),
		)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		defer resp.Body.Close()
}
	}
	