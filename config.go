package main

const (
	DEVELOPMENT = iota
	PRODUCTION
	TEST
)

var (
	Domain = "localhost"
	ListenPort = ":8888"
	Environment = PRODUCTION
	Logging = false
	LoggingPath = ""
	Version = "v1.0"
)
