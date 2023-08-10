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

	c := colly.NewCollector(
		colly.AllowedDomains("reddit.com", "www.reddit.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML(".cc-ranking", func(h *colly.HTMLElement) {
		println("matched")
		//println(h.ChildAttr("section", "cc-ranking__select clearfix"))
		item := h.ChildAttr("section", "cc-ranking__select clearfix")
		println(item)
		h.ForEach("section.cc-ranking__select clearfix", func(i int, e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("scraping done!")
	})

	c.Visit("https://na.finalfantasyxiv.com/lodestone/ranking/crystallineconflict/")

}
