package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//var (
//	IFconBase = baseCon.LoadBaseConfig()
//	redirectUrl = IFconBase.Host+":"+string(IFconBase.Port)+"/ifei"
//)

func LoginCheck(c *gin.Context) {

	session := sessions.Default(c)
	//session.Set("cotest","test~")
	//_ = session.Save()
	token, ok := session.Get("token").(string)
	if !ok {
		//c.Redirect(http.StatusFound, redirectUrl)
		c.AbortWithStatus(302)
	} else {
		logrus.Info("[LoginCheck] get token...", token)
	}
	get := session.Get("va")
	logrus.Info("get user info ", get)
	//s := get.(string)
	c.Next()
}
