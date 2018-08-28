package gofaquest

import (
	"net/http"
	"time"
)

/*
* client is maintained by GoFaquest and only one client is created
 */
var faquestClient *http.Client

func init() {
	faquestClient = &http.Client{}
}

/*
* Create a instance for gofaquest
 */
func NewGoFaquest() *GoFaquest {
	return &GoFaquest{
		params:  make(map[string]string, 0),
		headers: make(map[string]string, 0),
	}
}

/*
* Setup maxRetry time and default is 1
 */
func (self *GoFaquest) SetRetryTimes(maxRetryTimes int) {
	if maxRetryTimes <= 0 {
		maxRetryTimes = 1
	}
	self.retryTimes = maxRetryTimes
}

/*
* Setup params for request
 */
func (self *GoFaquest) SetParams(key string, value interface{}) {
	if key == "" {
		return
	}
	self.params[key] = value.(string)
}

/*
* Setup headers for request
 */
func (self *GoFaquest) SetHeaders(key string, value interface{}) {
	if key == "" {
		return
	}
	self.headers[key] = value.(string)
}

/*
* Setup cookies for request
 */
func (self *GoFaquest) SetCookies(key string, value string, duration time.Duration) {

}

func (self *GoFaquest) SetProxy() {

}
