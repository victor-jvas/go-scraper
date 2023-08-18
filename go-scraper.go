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

func main() {

	var ranking []player

	c := colly.NewCollector(
		colly.AllowedDomains("na.finalfantasyxiv.com"),
		colly.CacheDir("./ranking_cache"),
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

		ranking = append(ranking, p)

	})
	/*this should crawl and get all players from all servers
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		if !strings.Contains(link, "/crystallineconflict") {
			return
		}
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})
	*/

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
