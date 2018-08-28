package gofaquest

/*
* Gofaquest(Golang Fast Request) is a high-performance http request pool service for golang.
* retryTimes is used to setup retry time for specific request and min times is 0
 */
type GoFaquest struct {
	retryTimes int
	params     map[string]string
	headers    map[string]string
	cookies    map[string]string
	proxy      Proxy
	method     int
	timeout    int
}

/*
* Proxyinfo for Gofaquest
 */
type Proxy struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Result struct {
	Value []byte
	Error error
	Cost  int64
}

const (
	FAQUEST_METHOD_GET = iota
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
	0: "GET",
	1: "POST",
	2: "PUT",
	3: "PATCH",
	4: "DELETE",
	5: "COPY",
	6: "HEAD",
	7: "OPTIONS",
	8: "LINK",
	9: "UNLINK",
	10: "PURGE",
	11: "LOCK",
	12: "UNLOCK",
	13: "PROPFIND",
	14: "VIEW",
}
