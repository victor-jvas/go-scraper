package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type player struct {
	Name     string
	Position string
	World    string
	Tier     string
	Points   string
	Wins     string
}

func printPlayer(p player) {
	fmt.Println("Rank: ", p.Position)
	fmt.Println("Name: ", p.Name)
	fmt.Println("World: ", p.World)
	fmt.Println("Tier: ", p.Tier)
	fmt.Println("Points: ", p.Points)
	fmt.Println("Wins: ", p.Wins)
}

func main() {
	//_, file := os.Create("ranking.json")

	//println(file)

	var ranking []player

	c := colly.NewCollector(
		colly.AllowedDomains("na.finalfantasyxiv.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("div.ranking_set", func(h *colly.HTMLElement) {

		var p player
		p.Name = h.ChildText("div.name h3")
		p.Position = h.ChildText("div.order")
		p.World = h.ChildText("span.world ")
		p.Tier = h.ChildAttr("img.js--wolvesden-tooltip", "data-tooltip")
		p.Points = strings.Replace(h.ChildText("div.points"), "\t", "", -1)
		p.Wins = strings.Replace(h.ChildText("div.wins"), "\t", "", -1)

		//t := h.ChildAttr("img.js--wolvesden-tooltip", "data-tooltip")
		//fmt.Println(t)
		//fmt.Println(strings.TrimSpace(p.Points))

		ranking = append(ranking, p)

		//printPlayer(p)

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("scraping done!")
	})

	c.Visit("https://na.finalfantasyxiv.com/lodestone/ranking/crystallineconflict/")

	toJSON(ranking)

}

func toJSON(p []player) {
	file, err := json.MarshalIndent(p, "", " ")

	if err != nil {
		log.Println("Unable to create file")
		return
	}

	_ = os.WriteFile("crystaline-ranking.json", file, 0644)
}
