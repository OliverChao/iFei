package handlerFuncs

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"iFei/echo"
	"iFei/forever"
	"net/http"
)

func LoginIn(c *gin.Context) {
	//ret := echo.BindPostArticle{}
	ret := echo.NewRetResult()
	ret.Code = -1
	defer c.JSON(http.StatusOK, ret)

	//todo :add RSA decrypt if encrypt by js

	var arg echo.BindLogin
	if err := c.BindJSON(&arg); err != nil {
		ret.Msg = "post data format error"
		return
	}
	//把 sha注册到全局会变快一点点吗
	//
	sha := sha256.New()
	sha.Write([]byte(arg.Password))
	sum := sha.Sum(nil)
	passwd := hex.EncodeToString(sum)

	user, e := forever.VerifyUser(arg.Username, passwd)
	if e != nil {
		//logrus.Error()
		ret.Code = -2
		ret.Msg = "verify failed"
		return
	}

	// return some information about this user
	// todo ://change it
	ret.Data = user

	//todo : token
	session := sessions.Default(c)
	session.Set("token", "test token")
	session.Set("va", user)
	e = session.Save()

	if e != nil {
		logrus.Error("cookie save error")
		return
	} else {
		logrus.Info("cookie save successfully")
	}

	//extraData := make(map[string]interface{})
	//extraData["info"] =
}
