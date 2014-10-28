package post

import (
	"strings"

	"github.com/astaxie/beego"

	"github.com/movelikeriver/ishow/modules/post"
	"github.com/movelikeriver/ishow/modules/utils"
	"github.com/movelikeriver/ishow/routers/base"
)

type SearchRouter struct {
	base.BaseRouter
}

func (this *SearchRouter) Get() {
	this.TplNames = "search/posts.html"

	pers := 25

	q := strings.TrimSpace(this.GetString("q"))

	this.Data["Q"] = q

	if len(q) == 0 {
		return
	}

	page, _ := utils.StrTo(this.GetString("p")).Int()

	posts, meta, err := post.SearchPost(q, page)
	if err != nil {
		this.Data["SearchError"] = true
		beego.Error("SearchPosts: ", err)
		return
	}

	if len(posts) > 0 {
		this.SetPaginator(pers, meta.TotalFound)
	}

	this.Data["Posts"] = posts
	this.Data["Meta"] = meta
}
