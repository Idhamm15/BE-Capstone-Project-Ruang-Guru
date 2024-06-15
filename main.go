package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "be_capstone/handlers"
)

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

    r.Run(":8080")
}
