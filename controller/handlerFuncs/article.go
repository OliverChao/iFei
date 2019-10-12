package handlerFuncs

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"iFei/forever"
	"io/ioutil"
	"net/http"
)

func TestRedis(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"name":"oliver",
	//})
	s := forever.GetArticleInfo()
	if s != "" {
		c.JSON(http.StatusOK, s)
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"err": "errorororororororororo",
			"msg": "oliver loves annabelle",
		})
	}

}

func TestLuteEngine(c *gin.Context) {
	markdownEngine, e := forever.GetMarkdownEngine()
	if e != nil {
		c.JSON(404, nil)

		return
	}
	bytes, err := ioutil.ReadAll(c.Request.Body)
	logrus.Info(string(bytes))
	if err != nil {
		logrus.Error("[ioutil.ReadAll] " + err.Error())
	}
	html, e := markdownEngine.Markdown("", bytes)
	//logrus.Debug("html")
	c.JSON(200, string(html))
}
