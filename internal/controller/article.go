package controller

import (
	"web-app-distributor/internal/service/article"
	"web-app-distributor/internal/service/restaurant"
)

func GetArticle(id string) string {
	if a := article.Get().ById(id); a == nil {
		return ""
	} else if r := restaurant.Get().ById(a.RestaurantId); r == nil {
		return ""
	} else {
		content := r.StringRestaurantInfo() + a.Comment + r.StringDishes() + a.StringStar()
		return content
	}
}
