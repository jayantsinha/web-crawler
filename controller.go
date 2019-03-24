package main

import (
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type JsonResponse struct {
	Urls []struct {
		Loc        string        `json:"loc"`
		Title      string        `json:"title"`
		Image      string        `json:"image"`
		LinkedUrls []interface{} `json:"linked_urls"`
	} `json:"urls"`
}

const URL = "https://wiprodigital.com"
const PATTERN = "https?://([a-z0-9]+[.])*wiprodigital[.].*"

// ScrapingController is the handler for /crawl endpoint
func ScrapingController(ctx *gin.Context) {
	urls := JsonResponse{}
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Web Crawler WiproTest/v1.0"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		match, _ := regexp.MatchString(PATTERN, e.Request.AbsoluteURL(e.Attr("href")))
		if match {
			c.Visit(e.Attr("href"))
		}
	})

	c.OnHTML("html", func(el *colly.HTMLElement) {
		links := make([]string, 0, 1)
		links = el.ChildAttrs("body a[href]", "href")
		images := el.ChildAttrs("body img[src]", "src")
		s := struct {
			Loc        string        `json:"loc"`
			Title      string        `json:"title"`
			Image      string        `json:"image"`
			LinkedUrls []interface{} `json:"linked_urls"`
		}{}
		ll := make([]interface{}, 0, 1)
		for _, v := range links {
			match, _ := regexp.MatchString(PATTERN, v)
			if match {
				ll = append(ll, v)
			}
		}
		s.LinkedUrls = ll
		s.Title = el.DOM.Find("title").Text()
		s.Loc = el.Request.URL.String()
		for _, v := range images {
			if v != "" {
				s.Image = v
				break
			}
		}
		urls.Urls = append(urls.Urls, s)
	})

	err := c.Visit(URL)
	if err != nil {
		log.Println("Invalid URL: ", err)
		ctx.JSON(400, gin.H{"message": "Invalid URL: " + URL})
		return
	}
	c.Wait()

	log.Println("Found ", len(urls.Urls), " URLs")

	ctx.JSON(200, gin.H{"urlset": urls})
}