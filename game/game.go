package game

import (
	"bufio"
	"fmt"
	"lenimal.nl/tycoon/building"
	"lenimal.nl/tycoon/customer"
	"lenimal.nl/tycoon/neighbourhood"
	"lenimal.nl/tycoon/player"
	"lenimal.nl/tycoon/shop"
	"os"
	"strings"
)

type Game struct {
	Finished       bool
	Player         player.Player
	Building       building.Building
	Neighbourhoods []neighbourhood.Neighbourhood
	Customers      []customer.Customer
	AvailableShops []shop.Shop
}

func (g *Game) Setup() {

}

func (g *Game) Loop() {
	reader := bufio.NewReader(os.Stdin)

	g.getMenu(reader)
}

func (g *Game) getMenu(reader *bufio.Reader) {
	fmt.Printf(`
Welcome to build-a-mall +
Menu:

[b] Buy
[u] Upgrade
[i] Invest
[r] Run Gameday
[s] Get Status
[e] Exit
`)
	menuChoice := getLine(reader)

	switch menuChoice {
	case "b":
		g.buyMenu(reader)
	case "u":
		g.upgradeMenu(reader)
	case "i":
		g.investMenu(reader)
	case "e":
		g.Finished = true
	case "s":
		g.PrintGameState()

	case "r":
		g.RunDay()
	}
}

func (g *Game) upgradeMenu(reader *bufio.Reader) {
	fmt.Printf(`
Upgrade:
[s] Shop
[r] Return

`)
	upgradeChoice := getLine(reader)

	switch upgradeChoice {
	case "s":
		// show shops in mall
		// select a shop
		// check money
		// level shop up
	case "r":
		g.getMenu(reader);
	}

}

func (g *Game) buyMenu(reader *bufio.Reader) {
	fmt.Printf(`
Buy:
[b] Build shop
[s] Buy shop plans
[r] Return
`)
	buyChoice := getLine(reader)

	switch buyChoice {
	case "b":
	// show available shops
	// select shop
	// check money
	// substract money
	// select position
	case "s":
		// show all shops
		// select shop
		// check hearts?
		// subtract hearts
		// add to available shops
	case "r":
		g.getMenu(reader);
	}
}

func (g *Game) investMenu(reader *bufio.Reader) {
	fmt.Printf(`
Invest:
[n] New neighbourhood
[c] Choose neighbourhood
[r] Return
`)
	buyChoice := getLine(reader)

	switch buyChoice {
	case "r":
		g.getMenu(reader);
	}
}

func (g *Game) RunDay() {
	// todo: add ticks?
	g.Player.Money += 1000
	g.Player.Hearts += 4
	g.Building.Popularity++
	for _, v := range g.Building.Shops {
		v.Reputation++
	}

	g.PrintGameState();
}

func (g *Game) PrintGameState() {
	fmt.Printf(`
Money: %d,
Hearts: %d,

Building:
Popularity: %d

`, g.Player.Money, g.Player.Hearts, g.Building.Popularity)
	for _, v := range g.Building.Shops {
		fmt.Printf("Shop %s: Rep: %d", v.Name, v.Reputation)
	}
}

func getLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}
