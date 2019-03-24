### Web Crawler using Go

This is a simple implementation of a web crawler in Go. It scans for all the URLs in the given domain.

Features:
* Scrapes all the url on the same domain
* Async scraping
* Uses super fast Gin framework for running an HTTP rest server

Limitations:
* No support for multi-domain scraping
* No support for web pages behind authentication
* No support for web pages behind forms
* No support for dynamic Ajax web pages
* The scraper does not support invisible scraping 

Requirement
* Go version 1.12 or above
