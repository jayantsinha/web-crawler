package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// JSONResponse is the struct for json response for /crawl endpoint
type JSONResponse struct {
	Urls []struct {
		Loc        string        `json:"loc"`
		Title      string        `json:"title"`
		Image      string        `json:"image"`
		LinkedUrls []interface{} `json:"linked_urls"`
	} `json:"urls"`
}

// URL holds the X-Scrape header value
var URL string

// Pattern is the regex pattern to filter out urls outside the domain specified in URL variable
var Pattern string

// MaxDepth holds the X-Max-Depth header value. It limits the recursion depth of the visited URLs.
// If set to 0, it sets the scraper for infinite recursion
var MaxDepth string

// ScrapingController is the handler for /crawl endpoint
func ScrapingController(ctx *gin.Context) {
	URL = ctx.GetHeader("X-Scrape")
	MaxDepth = ctx.GetHeader("X-Max-Depth")
	if URL == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid header"})
		return
	}

	Pattern, err := createRegexPattern(URL)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid URL in request"})
		return
	}

	urls := JSONResponse{}
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Web Crawler WiproTest/v1.0"),
	)

	if MaxDepth != "" {
		c.MaxDepth, err = strconv.Atoi(MaxDepth)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "X-Max-Depth should be >= 0"})
			return
		}
	}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		match, _ := regexp.MatchString(Pattern, e.Request.AbsoluteURL(e.Attr("href")))
		if match {
			c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
		}
	})

	c.OnHTML("html", func(el *colly.HTMLElement) {
		unqlinks := make(map[string]bool)
		links := el.ChildAttrs("body a[href]", "href")
		for _, link := range links {
			if len(link) > 1 {
				unqlinks[link] = true
			}
		}
		images := el.ChildAttrs("body img[src]", "src")
		s := struct {
			Loc        string        `json:"loc"`
			Title      string        `json:"title"`
			Image      string        `json:"image"`
			LinkedUrls []interface{} `json:"linked_urls"`
		}{}
		ll := make([]interface{}, 0, 1)
		if len(unqlinks) > 0 {
			for k, _ := range unqlinks {
				if k[0] != '#' {
					ll = append(ll, el.Request.AbsoluteURL(k))
				}
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

	err = c.Visit(URL)
	if err != nil {
		log.Println("Invalid URL: ", err)
		ctx.JSON(400, gin.H{"message": "Invalid URL: " + URL})
		return
	}
	c.Wait()

	log.Println("Found ", len(urls.Urls), " URLs")

	ctx.JSON(200, gin.H{"urlset": urls})
}

func createRegexPattern(u string) (string, error) {
	pu, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	pat := "https?://([a-z0-9]+[.])*"
	pat += pu.Host
	pat += ".*"
	return pat, nil
}
