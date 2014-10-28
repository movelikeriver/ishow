package article

import (
	"github.com/movelikeriver/ishow/modules/models"
	"github.com/movelikeriver/ishow/routers/base"
)

type ArticleRouter struct {
	base.BaseRouter
}

func (this *ArticleRouter) loadArticle(article *models.Article) bool {
	uri := this.Ctx.Request.RequestURI
	err := models.Articles().RelatedSel("User").Filter("IsPublish", true).Filter("Uri", uri).One(article)
	if err == nil {
		this.Data["Article"] = article
	} else {
		this.Abort("404")
	}
	return err != nil
}

func (this *ArticleRouter) Show() {
	this.TplNames = "article/show.html"
	article := models.Article{}
	if this.loadArticle(&article) {
		return
	}
}
