package forum

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func Start(args []string) {
	// Selecting port
	port := "8000"
	if len(args) == 3 {
		port = args[2]
	}

	// Current path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// Info
	fmt.Println("---------------")
	fmt.Println("Selected port: " + port)
	fmt.Println("---------------")

	// Creating router
	router := gin.Default()
	router.LoadHTMLGlob(dir + "/front/forum/index.html")

	router.Static("/src", dir+"/front/forum/static/")
	router.StaticFile("/favicon.ico", dir+"/favicon/favicon.ico")
	getApiRouter(router)

	router.NoRoute(func(c *gin.Context) {
		path := strings.Split(c.Request.URL.Path, "/")
		if (path[1] != "") && (path[1] == "api") {
			c.JSON(404, gin.H{
				"status":  10,
				"message": "Not found",
				"body":    nil,
			})
		} else {
			c.HTML(http.StatusOK, "index.html", "")
		}
	})

	// Starting
	router.Run(":" + port)
}

func getApiRouter(baseRouter *gin.Engine) {
	GenerateCategory()

	api := baseRouter.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/check", CheckToken)
			auth.GET("/user", GetUser)
			auth.POST("/registration", Registration)
			auth.POST("/login", Login)
			auth.POST("/logout", Logout)
		}

		get := api.Group("/get")
		{
			get.GET("/categories", GetAllCategories)
			get.GET("/threads", GetAllThreads)
			get.GET("/thread", GetThread)
			get.GET("/user", GetTargetUser)
		}

		forum := api.Group("/forum")
		{
			forum.POST("/create", CreateThread)
			forum.POST("/send", SendPost)
			forum.POST("/reputation", ChangeReputation)
		}
	}
}
