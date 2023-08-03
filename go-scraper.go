package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("Hello, world")

	_, file := os.Create("ranking.json")

	println(file)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting crystalline conflicts ranking")
	})

	c.Visit("https://na.finalfantasyxiv.com/lodestone/ranking/crystallineconflict/")

}
