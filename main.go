package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "be_capstone/handlers"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    r := gin.Default()

    chatHandler := handlers.NewChatHandler()
    r.POST("/chat", chatHandler.HandleChat)

    r.Run(":8080")
}
