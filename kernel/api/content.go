package api

import (
	"github.com/setcy/wiki/kernel/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/setcy/wiki/kernel/util"
)

type PageTree struct {
	// The explicit Text for the tree node. Exist in all tree node.
	Text string `json:"text"`

	// The text in url for the tree node, which is saved as the path in database. Exist in all tree node.
	url string

	// The Link for the tree leaf, user can click to jump to the page. Exist in leaf only.
	Link string `json:"link,omitempty"`

	// The children of the tree node. Exist in node only.
	Items []*PageTree `json:"items,omitempty"`

	// Same as Items, but use map to quick find. Exist in node only.
	treeMap map[string]*PageTree
}

func getMeta(c *gin.Context) {
	ret := util.NewResult()
	defer c.JSON(http.StatusOK, ret)

	metas, err := sql.QueryPageMeta()
	if err != nil {
		ret.SetError(-1, err.Error())
		return
	}

	root := &PageTree{
		Items:   []*PageTree{},
		treeMap: map[string]*PageTree{},
	}

	for _, meta := range metas {
		path := strings.Split(meta.Path, "/")
		if len(path) == 0 {
			continue
		}
		fatherNode := root
		for i := 0; i < len(path); i++ {
			if path[i] == "" {
				continue
			}

			if i == len(path)-1 {
				// leaf
				leaf := &PageTree{
					Text: path[i],
					url:  path[i],
					Link: "/" + meta.Path,
				}
				fatherNode.Items = append(fatherNode.Items, leaf)
				fatherNode.treeMap[path[i]] = leaf

			} else {
				// node
				var node *PageTree

				if n, ok := fatherNode.treeMap[path[i]]; ok {
					node = n
				} else {
					node = &PageTree{
						Text:    path[i],
						url:     path[i],
						Items:   []*PageTree{},
						treeMap: map[string]*PageTree{},
					}
					fatherNode.Items = append(fatherNode.Items, node)
					fatherNode.treeMap[path[i]] = node
				}

				fatherNode = node
			}
		}
	}

	ret.Data = root.Items

	return
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

	content, toc, err := sql.QueryContentByPath(path)
	if err != nil {
		ret.SetError(-2, err.Error())
		return
	}

	ret.Data = map[string]interface{}{
		"content": content,
		"toc":     toc,
	}

	return
}
