package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

type JsonResponse struct {
	Urls []struct {
		Loc        string   `json:"loc"`
		Title      string   `json:"title"`
		Image      string   `json:"image"`
		LinkedUrls []string `json:"linked_urls"`
	} `json:"urls"`
}

var urlset Set


func ScrapingController(ctx *gin.Context) {
	urlset.New()
	c := colly.NewCollector(
		colly.UserAgent("GoWebCrawlerBot/1.0"),
		colly.URLFilters(
			regexp.MustCompile("https?://([a-z0-9]+[.])*wiprodigital[.].*"),
		),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		//Delay:      1 * time.Second,
	})
	urls := JsonResponse{}
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		match, _ := regexp.MatchString("https?://([a-z0-9]+[.])*wiprodigital[.].*", e.Request.AbsoluteURL(link))
		if match {

			urls.Urls = append(urls.Urls, struct {
				Loc        string   `json:"loc"`
				Title      string   `json:"title"`
				Image      string   `json:"image"`
				LinkedUrls []string `json:"linked_urls"`
			}{e.Request.AbsoluteURL(link), formatTitle(e.Text), "", []string{}})
		}

		c.Visit(e.Request.AbsoluteURL(link))
	})


	c.OnRequest(func(r *colly.Request) {


		urlset.Add(r.URL.String())
	})
	c.Visit("https://wiprodigital.com/")
	c.Wait()


	//for e, _ := range urlset.set {
	//
	//}

	ctx.JSON(200, gin.H{"urlset": urls})
}


func formatTitle(str string) string {
	str = strings.TrimSpace(strings.ReplaceAll(str, "\n", ""))
	//str = strings.ReplaceAll(str, "  ", "")
	if len(str) == 0 {
		return ""
	}
	title := strings.Builder{}
	title.WriteByte(str[0])
	separated := false
	for i:= 1; i< len(str); i++ {
		if str[i] == ' ' && str[i+1] == ' ' && !separated {
			title.WriteByte('-')
			separated = true
		}
		if !(str[i] == ' ' && str[i+1] == ' ') {
			title.WriteByte(str[i])
		}
	}
	return title.String()
}