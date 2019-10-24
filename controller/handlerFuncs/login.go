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

	//todo :add RSA decrypt if encrypt by client

	var arg echo.BindLogin
	if err := c.BindJSON(&arg); err != nil {
		ret.Msg = "post data format error"
		return
	}
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
	ret.Data = user.RetData()
	//dataMap := user.RetData()
	//dataMap["gen_time"] = time.Now().Unix()
	////dataMap["expire"] =
	//
	////, _ := json.Marshal(dataMap)
	//
	//echo.TokenData{
	//	Data: dataMap,
	//	Sign: "",
	//}

	//todo : token
	session := sessions.Default(c)
	session.Set("token", "test token")
	session.Set("va", user.Name)
	e = session.Save()

	if e != nil {
		logrus.Error("[LoginIn] cookie save error")
		return
	} else {
		logrus.Info("[LoginIn] cookie save successfully")
	}

	ret.Msg = "verify successfully"
	ret.Code = 1

	//extraData := make(map[string]interface{})
	//extraData["info"] =
}
