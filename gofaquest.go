package gofaquest

import (
	"net/http"
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

func (self *GoFaquest) SetUrl(url string) {
	self.url = url
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
* Disable CLS verify
 */
func (self *GoFaquest) DisableTLSVerify() {
	self.skipTLSVerify = true
}

/*
* Setup cookies for request
 */
func (self *GoFaquest) SetCookie(key string, value string) {
	self.cookies[key] = value
}

/*
* Setup cookies for request
 */
func (self *GoFaquest) SetCookies(kvs ...string) {
	if len(kvs)%2 != 0 {
		return
	}
	for i := 0; i < len(kvs); i += 2 {
		self.cookies[kvs[i]] = kvs[i+1]
	}
}

/*
*  Setup proxy
 */
func (self *GoFaquest) SetProxy(host string, port string, username string, password string) {
	self.proxy = Proxy{
		Enable:   true,
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

//Retry
func (self *GoFaquest) retry(retryTimes int, job func() error) error {
	var err error
	for i := 0; i < retryTimes; i++ {
		if e := job(); e != nil {
			err = e
			continue
		}
		return nil
	}
	return err
}
