package gofaquest

func (self *Proxy) Valid() bool {
	return self.Host != "" && self.Port != ""
}
