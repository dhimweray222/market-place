package config

import "market_place.com/models"

type Model struct{
	Model interface{}
}

func RegisterModels()[]Model  {
	return []Model{
		{Model: models.User{}},
		{Model: models.Cart{}},
		{Model: models.Product{}},
	}
}
