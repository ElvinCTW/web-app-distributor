package article

import (
	"web-app-distributor/internal/dao/articleDAO"
)

type Getter struct{}

var getter = &Getter{}

func Get() *Getter {
	return getter
}

func (c *Getter) ById(id string) *Data {
	resp := &Data{}
	resp.Data = articleDAO.GetDataById(id)
	return resp
}
