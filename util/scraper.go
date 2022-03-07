package util

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type Product struct {
	Name []map[string]interface{}
}

func Scrap(url string) Product {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	item := Product{}
	var locations []map[string]interface{}

	c.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		locations = append(locations, map[string]interface{}{
			"description": e.Attr("content"),
		})

	})
	c.OnHTML("meta[name=Description]", func(e *colly.HTMLElement) {
		locations = append(locations, map[string]interface{}{
			"Description": e.Attr("content"),
		})

	})
	c.OnHTML("meta[name=Keywords]", func(e *colly.HTMLElement) {
		locations = append(locations, map[string]interface{}{
			"Keywords": e.Attr("content"),
		})

	})
	c.OnHTML("meta[name=keywords]", func(e *colly.HTMLElement) {
		locations = append(locations, map[string]interface{}{
			"keywords": e.Attr("content"),
		})

	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.Visit(url)
	item.Name = append(item.Name, locations...)

	return item

}
