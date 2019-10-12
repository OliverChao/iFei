package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouterMap() *gin.Engine {
	engine := gin.Default()

	//engine := gin.New()
	//engine.Use(gin.Recovery())
	//engine.Use(gin.Logger())
	//sessionRedis.NewStore()
	//store := cookie.NewStore([]byte("secret"))
	//store.Options(sessions.Options{
	//	Path:     "/",
	//	MaxAge:   21600,
	//	Secure:   strings.HasPrefix("http://127.0.0.1", "https"),
	//	HttpOnly: true,
	//})
	//engine.Use(sessions.Sessions("ifei", store))
	engine.GET("/engine_test", func(c *gin.Context) {
		//var m map[string]string
		m := make(map[string]interface{})
		defer c.JSON(http.StatusOK, m)
		m["test"] = "engine_test"
		m["name"] = "oliver"
		m["girlfriend"] = "annabelle"
		m["age"] = 9999
		m["s"] = "oliver loves annabelle~"

	})
	api := engine.Group("/api")

	api.GET("/test_api", func(c *gin.Context) {
		m := make(map[string]interface{})
		defer c.JSON(http.StatusOK, m)
		m["test"] = "api test"
		m["name"] = "oliver"
		m["girlfriend"] = "annabelle"
		m["age"] = 9999
		m["s"] = "oliver loves annabelle~"
		m["zero data"] = map[string]interface{}{
			"test_data": "annabelle is a lovely girl",
		}
	})

	//engine.StaticFile()

	return engine
}
