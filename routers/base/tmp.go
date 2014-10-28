package base

import (
	"github.com/movelikeriver/ishow/modules/mailer"
)

type TestRouter struct {
	BaseRouter
}

func (this *TestRouter) Get() {
	this.TplNames = this.GetString(":tmpl")
	this.Data = mailer.GetMailTmplData(this.Locale.Lang, &this.User)
	this.Data["Code"] = "CODE"
}
