package controller

import "web-app-distributor/internal/service/restaurant"

func ShowRestaurant(id string) restaurant.Data {
	return restaurant.Get().ById(id)
}
