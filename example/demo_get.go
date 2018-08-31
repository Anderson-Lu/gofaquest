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
