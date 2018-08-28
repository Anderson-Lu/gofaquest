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
