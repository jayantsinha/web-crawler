package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return router
}

func main() {

	router := setupRouter()
	if Environment == PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Enable gin logging if enabled in config
	if Logging {
		router.Use(gin.Logger())
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Web Crawler using Go",
		})
	})

	router.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Web Crawler " + Version,
		})
	})

	srv := &http.Server{
		Addr:    ListenPort,
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting Down Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
