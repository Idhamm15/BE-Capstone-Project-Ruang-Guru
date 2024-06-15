package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"be_capstone/handlers"
)

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}

func main() {
    // Set mode release jika tidak dalam debug
	// gin.SetMode(gin.ReleaseMode)

    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    r := gin.Default()

    chatHandler := handlers.NewChatHandler()
    r.GET("/", func(c *gin.Context) {
		// Handler logic
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
    r.POST("/chat", chatHandler.HandleChat)

    // Set trusted proxies
	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	// Use `PORT` provided in environment or default to 3000
	var port = envPortOr("3000")

	// Mulai server
	log.Fatal(http.ListenAndServe(port, r))
}
