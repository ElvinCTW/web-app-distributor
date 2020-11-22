package restaurant

import (
	"strconv"
	"web-app-distributor/internal/dao/restaurantDAO"
)

type Data struct {
	*restaurantDAO.Data
}

func New(address, city, distinct, genre string) Data {
	r := Data{
		&restaurantDAO.Data{
			Info: restaurantDAO.Info{
				Address:  address,
				City:     city,
				Distinct: distinct,
				Genre:    genre,
			},
			Dishes:   []restaurantDAO.Dish{},
			Pictures: []string{},
		},
	}

	return r
}

func (c *Data) Save() string {
	return restaurantDAO.Create(c.Data)
}

//todo
func (c *Data) AddRestaurantPicture(pictureId string) {
	c.Pictures = append(c.Pictures, pictureId)
}

//todo
func (c *Data) AddDishPicture(pictureId, dishName string, dishPrice int) error {
	return nil
}

// todo
func (c *Data) addDishes(pictureId, dishName string, dishPrice int) error {
	return nil
}

// todo
func addPicutures() error {
	return nil
}

func (c *Data) StringRestaurantInfo() string {
	return "[" + c.Info.City + "。" + c.Info.Distinct + "。" + c.Info.Genre + "]\n"
}

func (c *Data) StringDishes() string {
	var s string
	for _, d := range c.Dishes {
		s = s + stringDish(d)
	}
	return s
}
func stringDish(d restaurantDAO.Dish) string {
	return "- " + d.Name + " $" + strconv.Itoa(d.Price) + "\n"
}
