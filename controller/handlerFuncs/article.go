package handlerFuncs

import (
	"github.com/gin-gonic/gin"
	"iFei/forever"
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
