package main

import (
	"fmt"
	"lenimal.nl/tycoon/building"
	"lenimal.nl/tycoon/customer"
	"lenimal.nl/tycoon/game"
	"lenimal.nl/tycoon/neighbourhood"
	"lenimal.nl/tycoon/player"
	"lenimal.nl/tycoon/shop"
)

func main() {
	fmt.Printf("started\n")

	p := player.Player{
		Id:     1,
		Name:   "lenn",
		Money:  5000,
		Hearts: 10,
	}

	b := building.Building{
		Id:         1,
		Name:       "building",
		Popularity: 0,
		Shops:      []shop.Shop{},
		Level:      0,
		MaxFloors:  4,
	}
	as := []shop.Shop{
		shop.Shop{
			Id:            1,
			Name:          "shop1",
			ShopTypeId:    1,
			ShopUpgrades:  []shop.ShopUpgrade{},
			Level:         0,
			BuildPrice:    1000,
			IsAvailable:   true,
			OperationCost: 100,
			Stock:         10,
			Reputation:    10,
			Popularity:    10,
			Floor:         0,
			FloorPosition: 0,
		},
		shop.Shop{
			Id:           2,
			Name:         "shop2",
			ShopTypeId:   1,
			ShopUpgrades: []shop.ShopUpgrade{},
			Level:        0,
			BuildPrice:   2000,
			IsAvailable:  false,
			HeartPrice:   10,

			OperationCost: 300,
			Stock:         20,
			Reputation:    30,
			Popularity:    30,
			Floor:         0,
			FloorPosition: 0,
		},
	}
	c := []customer.Customer{}
	n := []neighbourhood.Neighbourhood{}

	g := game.Game{
		Finished:       false,
		Player:         p,
		Building:       b,
		Customers:      c,
		AvailableShops: as,
		Neighbourhoods: n,
	}

	g.Setup()

	for g.Finished != true {
		fmt.Printf("loop\n")
		g.Loop()
	}

	fmt.Printf("finished\n")

}
