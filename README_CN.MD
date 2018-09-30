Gofaquest可复用的高性能HTTP下载器,用于下载网页,接口交互等.

[英文文档](https://github.com/Anderson-Lu/gofaquest/blob/master/README.MD)

### 依赖


要使用`gofaquest`,只需要添加以下依赖即可:

```golang
go get github.com/Anderson-Lu/gofaquest
```

### 设置代理

有时候,我们需要通过代理去请求某些特定的接口和页面,在这里,我封装了原生的代理设置,使得设置代理更加简单方便,通过以下方式可以直接设置代理信息:

```go
func ProxyDemo() {
  request := gofaquest.NewGoFaquest()
  request.SetUrl("http://www.baidu.com")
  request.SetProxy("127.0.0.1", "1080", "username", "password")
  request.SetUserAgent(gofaquest.Chrome.Latest())
  result := request.Get()
  fmt.Println(result)
}
```

注意,如果代理不需要设置用户名和密码,则相应参数传递空字符串即可.

### 设置Content-Type

有时候,我们需要设置请求的Content-Type属性,比如`application/x-www-form-urlencoded`以及`application/json;charset=UTF-8`等.

```golang
request.SetContentType("application/x-www-form-urlencoded")
```

### 设置User-Agent

在一些爬虫场景下,我们需要时常切换代理字符串,比如使用IE代理或者谷歌代理等,Gofaquest提供了`IE`和`Chrome`常见的用户代理字符串,同时支持用户自定义`User-Agent`,代码如下:

```golang
func UserAgentDemo() {
  request := gofaquest.NewGoFaquest()
  request.SetUrl("http://www.baidu.com")
  request.SetUserAgent(gofaquest.Chrome.Latest())
  result := request.Get()
  fmt.Println(result)
}
```

我们提供了`gofaquest.Chrome`和`gofaquest.IE`两个常量来维护常用的代理字符串,并提供了`Latest()`和`Random()`两个方法,来获取最新的代理字符串和随机代理字符串.

### 设置Cookie

`Gofaquest`提供了`SetCookie(key,value string)`和`SetCookies(keyValues ...string)`两个方法来设置请求的Cookie信息,很方便:

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

### 设置Header

`SetHeaders(key string,value interface{})`允许你非常便捷地设置请求的头部信息,支持直接设置数字,字符串甚至是对象类型的数据,避免了重复的数据类型转换操作.

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

### Disable TLSVerify

```golang
request.DisableTLSVerify()
```


### 设置请求参数

要设置参数,平时我们通过`url.Values`来进行设置,在`gofaquest`中,可以通过`SetParams(key,value)`来直接设置,支持直接设置不同类型的数据,我们会自动将非字符串数据转化为字符串,对象则执行序列化转化为对应的JSON数据.

```golang
func ParamsDemo() {
  request := gofaquest.NewGoFaquest()
  request.SetUrl("http://www.baidu.com")
  request.SetFormParams("param1","value1")
  request.SetFormParams("param2",1000)
  request.SetFormParams("param3",struct {Name string}{Name: "Anderson"})
  result := request.Get()
  fmt.Println(result)
}
```

此外,如果要在`body`中添加`json`数据,请使用`request.SetBodyJson()`方法

###  设置重试次数

有时候,可能因为一些网络问题,我们需要进行重试请求

```golang
request.SetRetryTimes(30)
```

### 设置请求超时时间

```golang
request.SetTimeout(time.Second*10)
```

### 设置请求方法

我们提供了常用的HTTP请求方法

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

### 结果解析

所有的请求的返回结果,都会通过`*Result`来返回,其结构如下:

```golang
type Result struct {
  Value []byte
  Error error
  Cost  int64
}
```

这里的`Value`即为返回的字节数组数据,`Cost`为耗时. 对于数据解析,可以直接使用`json.Unmarshal()`来解析数据,也可以使用[Gofasion](https://github.com/Anderson-Lu/gofasion)来更加便捷地解析数据.