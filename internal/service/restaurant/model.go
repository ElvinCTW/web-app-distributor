package restaurant

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Info      Info
	Dishes    []Dish
	PictureId []string
}

func New(Address, City, Distinct, Genre string) Restaurant {
	r := Restaurant{
		Info: Info{
			Address:  Address,
			City:     City,
			Distinct: Distinct,
			Genre:    Genre,
		},
		Dishes:    []Dish{},
		PictureId: []string{},
	}

	return r
}

// todo
func (c *Restaurant) AddDishPicture() error {
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

func (c *Restaurant) GetDishesString() string {
	var s string
	for _, d := range c.Dishes {
		s = s + d.String()
	}
	return s
}

type Info struct {
	Address  string
	City     string
	Distinct string
	Genre    string
}

func (c *Info) String() string {
	return "[" + c.City + "。" + c.Distinct + "。" + c.Genre + "]\n"
}

type Dish struct {
	Name      string
	Price     string
	PictureId []string
}

func (c *Dish) String() string {
	return "- " + c.Name + " $" + c.Price + "\n"
}
