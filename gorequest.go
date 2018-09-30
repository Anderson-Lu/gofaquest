package gofaquest

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (self *GoFaquest) do() *Result {
	if self.err != nil {
		return &Result{Value: nil, Error: self.err}
	}
	if self.method == FAQUEST_METHOD_UNSET {
		return &Result{Value: nil, Error: ERROR_METHOD_NOT_FOUND}
	}

	params := url.Values{}
	if len(self.params) > 0 {
		for k, v := range self.params {
			params.Set(k, v)
		}
	}

	paramsStr := ""
	if self.body != "" {
		paramsStr = self.body
	} else {
		paramsStr = params.Encode()
	}

	request, e := http.NewRequest(FAQUEST_METHOD_MAP[self.method], self.url, strings.NewReader(paramsStr))

	if e != nil {
		self.err = e
		return &Result{Value: nil, Error: e}
	}

	for k, v := range self.headers {
		request.Header.Set(k, v)
	}

	for k, v := range self.cookies {
		request.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}

	if len(self.userAgent) > 0 {
		request.Header.Set("User-Agent", self.userAgent)
	}

	if self.timeout <= 0 {
		self.timeout = 10
	}

	faquestClient.Timeout = time.Second * time.Duration(self.timeout)

	defer self.reset()

	tr := &http.Transport{}
	if self.skipTLSVerify {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	if self.proxy.Enable {
		proxyUrl := ""
		if self.proxy.Username != "" {
			proxyUrl = fmt.Sprintf("%s:%s@%s:%s", self.proxy.Username, self.proxy.Password, self.proxy.Host, self.proxy.Port)
		} else {
			proxyUrl = fmt.Sprintf("%s:%s", self.proxy.Host, self.proxy.Port)
		}
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		tr.Proxy = proxy
	}

	faquestClient.Transport = tr

	if self.retryTimes <= 0 {
		self.retryTimes = 1
	}

	var resp *http.Response
	var err error
	self.retry(self.retryTimes, func() error {
		resp, err = faquestClient.Do(request)
		return err
	})

	if err != nil {
		self.err = err
		return &Result{Value: nil, Error: err}
	}

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		self.err = err
		return &Result{Value: nil, Error: err}
	}

	return &Result{Value: ret, Error: nil}
}

func (self *GoFaquest) reset() {
	faquestClient.Timeout = time.Second * 10
	faquestClient.Transport = &http.Transport{}
}

func (self *GoFaquest) Get() *Result {
	self.method = FAQUEST_METHOD_GET
	return self.do()
}

func (self *GoFaquest) Post() *Result {
	self.method = FAQUEST_METHOD_POST
	return self.do()
}

func (self *GoFaquest) Put() *Result {
	self.method = FAQUEST_METHOD_PUT
	return self.do()
}

func (self *GoFaquest) Patch() *Result {
	self.method = FAQUEST_METHOD_PATCH
	return self.do()
}

func (self *GoFaquest) Delete() *Result {
	self.method = FAQUEST_METHOD_DELETE
	return self.do()
}

func (self *GoFaquest) Copy() *Result {
	self.method = FAQUEST_METHOD_COPY
	return self.do()
}

func (self *GoFaquest) Head() *Result {
	self.method = FAQUEST_METHOD_HEAD
	return self.do()
}

func (self *GoFaquest) Options() *Result {
	self.method = FAQUEST_METHOD_OPTIONS
	return self.do()
}

func (self *GoFaquest) Link() *Result {
	self.method = FAQUEST_METHOD_LINK
	return self.do()
}

func (self *GoFaquest) Unlink() *Result {
	self.method = FAQUEST_METHOD_UNLINK
	return self.do()
}

func (self *GoFaquest) Purge() *Result {
	self.method = FAQUEST_METHOD_PURGE
	return self.do()
}

func (self *GoFaquest) Lock() *Result {
	self.method = FAQUEST_METHOD_LOCK
	return self.do()
}

func (self *GoFaquest) Unlock() *Result {
	self.method = FAQUEST_METHOD_UNLOCK
	return self.do()
}

func (self *GoFaquest) Propfind() *Result {
	self.method = FAQUEST_METHOD_PROPFIND
	return self.do()
}

func (self *GoFaquest) View() *Result {
	self.method = FAQUEST_METHOD_VIEW
	return self.do()
}
