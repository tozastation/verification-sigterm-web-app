package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// reference by https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}