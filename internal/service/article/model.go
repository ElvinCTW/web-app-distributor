package article

import (
	"fmt"
	"web-app-distributor/internal/dao/articleDAO"
)

const (
	starStringPrefix  = "主觀評價: "
	starStringPostfix = "顆星"
)

type Data struct {
	*articleDAO.Data
}

func (c *Data) StringStar() string {
	return fmt.Sprintf(starStringPrefix+"%f"+starStringPostfix, c.Stars)
}

func (c *Data) SetRestaurant() error {
	return nil
}
