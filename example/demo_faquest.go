package main

import (
	"fmt"

	"github.com/Anderson-Lu/gofaquest"
)

func main() {
	HttpGetDemo()
}

func HttpGetDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetRetryTimes(1)
	result := request.Get()
	fmt.Println(string(result.Value))
}

func ProxyDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetProxy("127.0.0.1", "1080", "anderson", "helloworld")
	result := request.Get()
	fmt.Println(result)
}

func UserAgentDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetUserAgent(gofaquest.Chrome.Latest())
	result := request.Get()
	fmt.Println(result)
}

func CookieDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetCookie("key", "value")
	request.SetCookies("key1", "value1", "key2", "value2", "key3", "value3")
	result := request.Get()
	fmt.Println(result)
}

func HeaderDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetHeaders("name", 123456)
	request.SetHeaders("token", "abcde")
	request.SetHeaders("user-info", struct {Name string}{Name: "Anderson"})
	result := request.Get()
	fmt.Println(result)
}

func ParamsDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetParams("param1","value1")
	request.SetParams("param2",1000)
	request.SetParams("param3",struct {Name string}{Name: "Anderson"})
	result := request.Get()
	fmt.Println(result)	
}

