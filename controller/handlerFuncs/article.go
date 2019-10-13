package handlerFuncs

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"iFei/forever"
	"iFei/model"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func DealToMarkdown(c *gin.Context) {
	//todo: rewirite this return data
	retMap := make(map[string]interface{})
	defer c.JSON(http.StatusOK, retMap)
	markdownEngine, e := forever.GetMarkdownEngine()
	if e != nil {
		retMap["code"] = -1
		retMap["msg"] = "cannot get markdown engine"
		//c.JSON(404, nil)
		return
	}

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		retMap["code"] = -1
		retMap["msg"] = "post data format error"
		return
	}

	mtext := arg["text"].(string)
	html, err := markdownEngine.MarkdownStr("", mtext)
	if err != nil {
		logrus.Error("[ioutil.ReadAll] " + err.Error())
	}
	retMap["data"] = html
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//html, e := markdownEngine.Markdown("", bytes)
}

func PostArticle(c *gin.Context) {
	//todo: rewirite this return data
	retMap := make(map[string]interface{})
	defer c.JSON(http.StatusOK, retMap)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		retMap["code"] = -1
		retMap["msg"] = "post data format error"
		return
	}

	article := model.Article{
		UUID:        strings.Replace(uuid.NewV4().String(), "-", "", -1),
		PushedAt:    time.Now(),
		Title:       arg["title"].(string),
		Content:     arg["content"].(string),
		Tags:        arg["tags"].(string),
		Commentable: arg["commentable"].(bool),
		Stared:      arg["stared"].(bool),
		Topped:      arg["topped"].(bool),

		// Todo: 需要从session 的信息中获取 userid, 这里需要改
		UserID: 1,
	}
	forever.AddArticle(&article)
	retMap["code"] = 1
	retMap["msg"] = "add successfully"
}

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
	//logrus.Info(string(bytes))
	if err != nil {
		logrus.Error("[ioutil.ReadAll] " + err.Error())
	}
	html, e := markdownEngine.Markdown("", bytes)
	//logrus.Debug("html")
	c.JSON(200, string(html))
}
