package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func httpDo()  {
	v := url.Values{}
	v.Set("name", "反倒是")
	v.Set("phone", "13323456123")
	v.Set("source", "a")

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码

	req, _ := http.NewRequest("POST", "http://12zhuangxiu.com/server/join/", body)

	proxy, err := url.Parse("http://111.40.84.73:9797")
	if err != nil {
		panic(err)
	}
	client := &http.Client{
		Transport:&http.Transport{
			Proxy:http.ProxyURL(proxy),
		},
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;") //这个一定要加，不加form的值post不过去，被坑了两小时

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}


func testGo(i int){
	fmt.Println(i)
}
func main() {
	//httpDo()

	//for i :=0 ; i<100 ; i++  {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	//
	//
	//for i :=0 ; i<100 ; i++  {
	//	go func(i int) {
	//		fmt.Println(i)
	//	}(i)
	//}
	for i := 0; i< 2; i++  {
		for k := 0; k<10 ; k++  {
			go testGo(k)
		}
	}

}
