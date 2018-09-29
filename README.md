Gofaquest reusable high-performance HTTP downloader for downloading web pages, interface interactions, etc.

### Dependency

To use gofaquest, just need to do this:

```golang
go get github.com/Anderson-Lu/gofaquest
```

### Setup Proxy

Sometimes, we need to request certain interfaces and pages through the proxy. Here, I encapsulate the native proxy settings, making it easier and more convenient to set up the proxy. You can set the proxy information directly in the following ways:

```golang
func ProxyDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetProxy("127.0.0.1", "1080", "anderson", "helloworld")
	request.SetUserAgent(gofaquest.Chrome.Latest())
	result := request.Get()
	fmt.Println(result)
}
```

Note that if the proxy does not need to set a username and password, the corresponding argument can be passed an empty string.

### Setup UserAgent

In some crawler scenarios, we need to switch proxy strings from time to time. For example, using IE proxy or Google proxy, Gofaquest provides common user proxy strings for `IE` and `Chrome`, and supports user-defined `User-Agent`. ,code show as below:

```golang
func UserAgentDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetUserAgent(gofaquest.Chrome.Latest())
	result := request.Get()
	fmt.Println(result)
}
```

We provide two constants `gofaquest.Chrome` and `gofaquest.IE` to maintain common proxy strings, and provide two methods `Latest()` and `Random()` to get the latest proxy string. And random proxy strings.

### Setup Cookie

`Gofaquest` provides two methods: `SetCookie(key,value string)` and `SetCookies(keyValues ​​...string)` to set the cookie information of the request. It is very convenient:

```golang
func CookieDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetCookie("key","value")
	request.SetCookies("key1","value1","key2","value2","key3","value3")
	result := request.Get()
	fmt.Println(result)
}
```

### Setup Header

`SetHeaders(key string, value interface{})` allows you to set the header information of the request very conveniently, and supports direct setting of numbers, strings and even object type data, avoiding repeated data type conversion operations.

```golang
func HeaderDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetHeaders("name", 123456)
	request.SetHeaders("token", "abcde")
	request.SetHeaders("user-info", struct {Name string}{Name: "Anderson"})
	result := request.Get()
	fmt.Println(result)
}
```

### SetUp Params

To set the parameters, we usually set them by `url.Values`. In `gofaquest`, we can set them directly by `SetParams(key,value)`, which supports setting different types of data directly. We will automatically set non-characters. The string data is converted to a string, and the object is serialized and converted to the corresponding JSON data.

```golang
func ParamsDemo() {
	request := gofaquest.NewGoFaquest()
	request.SetUrl("http://www.baidu.com")
	request.SetParams("param1","value1")
	request.SetParams("param2",1000)
	request.SetParams("param3",struct {Name string}{Name: "Anderson"})
	result := request.Get()
	fmt.Println(result)
}
```

### Setup retry times

Sometimes, we may need to retry the request because of some network problems.

```golang
request.SetRetryTimes(30)
```

### Setup Method

```golang
func (self *GoFaquest) Get() *Result
func (self *GoFaquest) Post() *Result
func (self *GoFaquest) Put() *Result
func (self *GoFaquest) Patch() *Result
func (self *GoFaquest) Delete() *Result
func (self *GoFaquest) Copy() *Result
func (self *GoFaquest) Head() *Result
func (self *GoFaquest) Options() *Result
func (self *GoFaquest) Link() *Result
func (self *GoFaquest) Unlink() *Result
func (self *GoFaquest) Purge() *Result
func (self *GoFaquest) Lock() *Result
func (self *GoFaquest) Unlock() *Result
func (self *GoFaquest) Propfind() *Result
func (self *GoFaquest) View() *Result
```

### Response

The result of all requests will be returned by `*Result`, which has the following structure:

```golang
type Result struct {
	Value []byte
	Error error
	Cost  int64
}
```

Here `Value` is the returned byte array data, `Cost` is time consuming.

