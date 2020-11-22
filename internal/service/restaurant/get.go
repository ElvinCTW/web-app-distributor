package restaurant

import "web-app-distributor/internal/dao/restaurantDAO"

type Getter struct{}

var getter = &Getter{}

func (c *Getter) ById(id string) Data {
	resp := Data{}
	resp.Data = restaurantDAO.GetDataById(id)
	return resp
}
