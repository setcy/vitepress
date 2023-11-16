package api

import (
	"github.com/setcy/wiki/kernel/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/setcy/wiki/kernel/util"
)

func getMeta(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": []gin.H{
			{
				"text": "Home",
				"items": []gin.H{
					{
						"text": "Child",
						"link": "/child",
					},
				},
			},
			{
				"text": "About",
				"link": "/about",
			},
			{
				"text": "Contact",
				"link": "/contact",
			},
		},
	})
}

const contentPathPrefixLen = 9 // len("/content/")

func getContent(c *gin.Context) {
	ret := util.NewResult()
	defer c.JSON(http.StatusOK, ret)

	path := c.Request.URL.Path
	if len(path) <= contentPathPrefixLen {
		ret.SetError(-1, "invalid path")
		return
	}
	path = path[contentPathPrefixLen:]

	content, err := sql.QueryContentByPath(path)
	if err != nil {
		ret.SetError(-2, err.Error())
		return
	}

	ret.Data = content
	return
}
