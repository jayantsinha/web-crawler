### Web Crawler using Go

This is a simple implementation of a web crawler in Go. It scans for all the URLs in the given domain and returns a JSON response.

[![Go Report Card](https://goreportcard.com/badge/github.com/jayantsinha/web-crawler)](https://goreportcard.com/report/github.com/jayantsinha/web-crawler)  [![CircleCI](https://circleci.com/gh/jayantsinha/web-crawler/tree/master.svg?style=shield)](https://circleci.com/gh/jayantsinha/web-crawler/tree/master)
#### Features:
* Scrapes all the url on the given domain
* Async scraping
* Uses super fast Gin framework for running an HTTP rest server

#### Limitations:
* No support for multi-domain scraping
* No support for web pages behind authentication/forms
* No support for dynamic/Ajax-based web pages
* The scraper does not support invisible scraping
* No caching support

#### Requirement
* Go version 1.12 or above

#### Installation and Running
To build the source, go must be installed with GOROOT and GOPATH set correctly. Read [this](https://golang.org/doc/install) to set up your go environment. Once the setup is complete, clone the repository or put the source directory into `$GOPATH/src/`.

Now run `go get ./...` inside the source directory (web-crawler). This will download all the dependencies into `$GOPATH/pkg/`. If this does not work in rare case; install dependecies individually:<br/>
`go get -u github.com/gin-gonic/gin`<br/>
`go get -u github.com/gocolly/colly/...`<br/>
`go get github.com/gin-contrib/gzip`<br/>

After the dependencies are installed, run `go build` from inside the source directory. This will create an executable file according to the host OS. To run the crawler service use <br/>
LINUX/MAC: `$ ./web-crawler`<br/>
WINDOWS: `> web-crawler.exe`

If you want to build for a different target machine then use the following command:<br/>
WINDOWS(64-bit): `env GOOS=windows GOARCH=amd64 go build`<br/>
LINUX(64-bit):   `env GOOS=linux GOARCH=amd64 go build`<br/>
MAC(64-bit):     `env GOOS=darwin GOARCH=amd64 go build`<br/>

By default, the gin mode is set to 'DebugMode' so that when you run the executable; all the registered endpoints can be seen. This can be changed by changing the `Environment` var from `config.go`.

To run tests, use `go test`

Once the service is running, it will expose `GET /crawl` endpoint. Create a request as follows:<br/> 
**REQUEST:** `GET localhost:8888/crawl`<br/>
**HEADER:** `"Scrape": "https://wiprodigital.com"`<br/>

#### Performance
It takes 8.6 secs to crawl 226 URLs of [https://wiprodigital.com](https://wiprodigital.com) and create JSON response on an 8-core machine running windows 10 with 8GB of memory using Postman client.
There is no guarantee that the first run will return all the 226 URLs but tests show that it takes 2-4 initial runs to produce a consistent result. This depends on many parameters like limits in domain, NS, etc.

#### Milestone
The limitations mentioned will be removed in the next release(s).