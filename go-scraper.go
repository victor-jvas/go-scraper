package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type player struct {
	Name     string
	Position int
	World    string
}

func main() {
	//_, file := os.Create("ranking.json")

	//println(file)

	//var ranking []player

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML(".cc-ranking__result__name", func(h *colly.HTMLElement) {
		println(h.ChildText("h3"))

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("scraping done!")
	})

	c.Visit("https://na.finalfantasyxiv.com/lodestone/ranking/crystallineconflict/")

}
