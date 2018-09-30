package gofaquest

import "fmt"

/*
* Gofaquest(Golang Fast Request) is a high-performance http request pool service for golang.
* retryTimes is used to setup retry time for specific request and min times is 0
 */
type GoFaquest struct {
	err           error
	retryTimes    int
	params        map[string]string
	headers       map[string]string
	cookies       map[string]string
	body          string
	proxy         Proxy
	method        int
	timeout       int
	url           string
	skipTLSVerify bool
	userAgent     string
}

/*
* Proxyinfo for Gofaquest
 */
type Proxy struct {
	Enable   bool
	Host     string
	Port     string
	Username string
	Password string
}

/*
* Result stores request result bytes or error if occours
 */
type Result struct {
	Value []byte
	Error error
	Cost  int64
}

const (
	FAQUEST_METHOD_UNSET = iota
	FAQUEST_METHOD_GET
	FAQUEST_METHOD_POST
	FAQUEST_METHOD_PUT
	FAQUEST_METHOD_PATCH
	FAQUEST_METHOD_DELETE
	FAQUEST_METHOD_COPY
	FAQUEST_METHOD_HEAD
	FAQUEST_METHOD_OPTIONS
	FAQUEST_METHOD_LINK
	FAQUEST_METHOD_UNLINK
	FAQUEST_METHOD_PURGE
	FAQUEST_METHOD_LOCK
	FAQUEST_METHOD_UNLOCK
	FAQUEST_METHOD_PROPFIND
	FAQUEST_METHOD_VIEW
)

var FAQUEST_METHOD_MAP = map[int]string{
	1:  "GET",
	2:  "POST",
	3:  "PUT",
	4:  "PATCH",
	5:  "DELETE",
	6:  "COPY",
	7:  "HEAD",
	8:  "OPTIONS",
	9:  "LINK",
	10: "UNLINK",
	11: "PURGE",
	12: "LOCK",
	13: "UNLOCK",
	14: "PROPFIND",
	15: "VIEW",
}

var (
	ERROR_METHOD_NOT_FOUND = fmt.Errorf("gofaquest error: invalid request method,please setup first")
)
