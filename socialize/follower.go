package socialize

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func followerlist() {

	url := "/douyin/relation/follower/list/?user_id=&token="
	method := "GET"

	client := &http.Client{}
	// 创建新请求
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	// 发送api请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
