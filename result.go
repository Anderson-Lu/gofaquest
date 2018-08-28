package gofaquest

import (
	"encoding/json"
)

func (self *Result) CostMillseconds() int64 {
	return self.Cost
}

func (self *Result) Scan(target interface{}) error {
	return json.Unmarshal(self.Value,&target)
}