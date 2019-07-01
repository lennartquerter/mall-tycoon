package building

import "lenimal.nl/tycoon/shop"

type Building struct {
	Id         int64
	Name       string
	Shops      []shop.Shop
	Popularity int64
	Level      int32
	MaxFloors  int32
}
