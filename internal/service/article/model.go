package article

import (
	"fmt"
	"web-app-distributor/internal/dao/articleDAO"
)

type Article struct {
	*articleDAO.Data
}

func (c *Article) SetRestaurant() error {
	return nil
}

func (c *Article) GetContent() string {
	var content string

	// Get title from c.restaurant.info.string()
	// Get comment
	content = content + c.Comment
	// Get dishes from c.restaurant.dishes.getDishesStrings()
	// Get stars
	content = content + fmt.Sprintf("%f", c.Stars)

	return content
}

func Get(id string) Article {
	return Article{
		Data: articleDAO.Get(id),
	}
}
