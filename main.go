package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"

	"be_capstone/handlers"
	"be_capstone/middleware"
	"be_capstone/database"
)

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}

func main() {
	// Set mode release jika tidak dalam debug
	// gin.SetMode(gin.ReleaseMode)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.Init()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	authHandler := handlers.NewAuthHandler()
	articleHandler := handlers.NewArticleHandler()
	chatHandler := handlers.NewChatHandler()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/logout", authHandler.Logout) // Tambahkan rute logout

	r.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
		role, _ := c.Get("role")
		if role == "admin" {
			c.JSON(http.StatusOK, gin.H{"message": "Kamu admin"})
		} else if role == "user" {
			c.JSON(http.StatusOK, gin.H{"message": "Kamu user"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Kamu belum login"})
		}
	})

	// Public routes
	r.GET("/articles", articleHandler.GetArticles)
	r.GET("/articles/:id", articleHandler.GetArticle)

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/articles", articleHandler.CreateArticle)
		admin.PUT("/articles/:id", articleHandler.UpdateArticle)
		admin.DELETE("/articles/:id", articleHandler.DeleteArticle)
	}
	
    r.POST("/chat", chatHandler.HandleChat)

	// Set trusted proxies
	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	// Use `PORT` provided in environment or default to 3000
	var port = envPortOr("5000")

	// Mulai server
	log.Fatal(http.ListenAndServe(port, r))
}
