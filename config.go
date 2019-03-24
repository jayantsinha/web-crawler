package main

const (
	DEVELOPMENT = iota
	PRODUCTION
	TEST
)

var (
	// Domain name of the http server
	Domain = "localhost"
	// ListenPort for gin http server
	ListenPort = ":8888"
	// Environment sets the gin mode
	Environment = TEST
	// Logging is a flag which will tell the server whether to enable logging or not
	// enabled with log path, the log will be saved to the file specified in LoggingPath
	Logging = false
	// LoggingPath is the absolute path of the file to use for logging.
	// Please set proper owner and permissions for the file
	LoggingPath = "/var/log/web-crawler.log"
	// Version of the web crawler
	Version = "v1.0"
)
