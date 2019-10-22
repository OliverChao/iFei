package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"iFei/controller/handlerFuncs"
	"iFei/controller/middleware"
	"iFei/echo"
	"iFei/forever"
	"net/http"
	"strings"
)

func RegisterRouterMap() *gin.Engine {
	engine := gin.Default()
	//engine := gin.New()
	//engine.Use(gin.Recovery())
	//engine.Use(gin.Logger())
	//sessionRedis.NewStore()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   21600,
		Secure:   strings.HasPrefix("http://127.0.0.1", "https"),
		HttpOnly: true,
	})
	engine.Use(sessions.Sessions("ifei", store))

	engine.Any("/", func(c *gin.Context) {
		defer c.JSON(200, "Yoo~~~ Hello~~~ iFei~~~")
		//session := sessions.Default(c)
		//session.Set("engineCookie","test")
		//_ = session.Save()
	})
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

	engine.GET("/get_redis", handlerFuncs.TestRedis)
	engine.POST("/markdown_load", handlerFuncs.TestLuteEngine)

	engine.POST("/ifei/login", handlerFuncs.LoginIn)
	engine.GET("ifei/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Options(sessions.Options{
			Path:   "/",
			MaxAge: -1,
		})
		session.Clear()
		if err := session.Save(); nil != err {
			logrus.Errorf("[logout] saves session failed: " + err.Error())
		}

		c.JSON(http.StatusOK, "clear cookie successfully")
	})

	engine.GET("/ifei/test_login", func(c *gin.Context) {
		ret := echo.NewRetResult()
		ret.Code = -1
		defer c.JSON(http.StatusOK, ret)

		username := c.DefaultQuery("username", "oliver")
		password := c.DefaultQuery("password", "toor")
		logrus.Infof("receive para u:%s p%s", username, password)
		sha := sha256.New()
		sha.Write([]byte(password))
		sum := sha.Sum(nil)
		passwd := hex.EncodeToString(sum)

		user, e := forever.VerifyUser(username, passwd)
		if e != nil {
			ret.Msg = "[test_login] user verify failed..."
			logrus.Info("user verify failed...")
			return
		}

		session := sessions.Default(c)
		session.Set("token", "test token")
		session.Set("va", user.Name)
		e = session.Save()
		if e != nil {
			logrus.Error("cookie save error", e)
			return
		} else {
			logrus.Info("cookie save successfully")
		}

		//todo: write login msg to redis server
		newToken := echo.GenerateToken()
		client := forever.GetGlobalRedisClient()
		e = client.Set(username, newToken, -1).Err()
		if e != nil {
			logrus.Error("[redis] save login msg error")
		} else {
			logrus.Info("[redis] save login msg successfully")
		}
		//

		ret.Msg = "verify successfully"
		ret.Code = 1
		ret.Data = user

	})
	//engine
	api := engine.Group("/api")
	api.Use(middleware.LoginCheck)

	api.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		get := session.Get("cotest")
		logrus.Info("get cotest cookie:", get)

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

	api.POST("/markdown", handlerFuncs.DealToMarkdown)
	//engine.StaticFile()
	api.POST("/post", handlerFuncs.PostArticle)
	//api.GET("/article")

	return engine
}
