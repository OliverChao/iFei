package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoginCheck(c *gin.Context) {
	session := sessions.Default(c)
	//session.Set("cotest","test~")
	//_ = session.Save()
	get := session.Get("va")
	logrus.Info("get user info ", get)
	//s := get.(string)
	c.Next()
}
