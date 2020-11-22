package restaurant

import (
	"web-app-distributor/internal/dao/restaurantDAO"
)

type Data struct {
	*restaurantDAO.Data
}

func New(Address, City, Distinct, Genre string) Data {
	r := Data{
		&restaurantDAO.Data{
			Info: restaurantDAO.Info{
				Address:  Address,
				City:     City,
				Distinct: Distinct,
				Genre:    Genre,
			},
			Dishes:    []restaurantDAO.Dish{},
			PictureId: []string{},
		},
	}

	return r
}

func Get() *Getter {
	return getter
}

func (c *Data) Save() string {
	return restaurantDAO.Create(c.Data)
}

// todo
func (c *Data) AddDishPicture() error {
	if err := addDishes(); err != nil {
		return err
	} else if err = addPicutures(); err != nil {
		return err
	} else {
		return nil
	}
}

// todo
func addDishes() error {
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
		s = s + getDishString(d)
	}
	return s
}
func getDishString(d restaurantDAO.Dish) string {
	return "- " + d.Name + " $" + d.Price + "\n"
}
