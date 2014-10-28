package api

import (
	"github.com/astaxie/beego/orm"

	"github.com/movelikeriver/ishow/modules/auth"
	"github.com/movelikeriver/ishow/modules/models"
	"github.com/movelikeriver/ishow/modules/utils"
	"github.com/movelikeriver/ishow/routers/base"
)

type ApiRouter struct {
	base.BaseRouter
}

func (this *ApiRouter) Users() {
	result := map[string]interface{}{
		"success": false,
	}

	defer func() {
		this.Data["json"] = result
		this.ServeJson()
	}()

	if !this.IsAjax() {
		return
	}

	action := this.GetString("action")

	if this.IsLogin {

		switch action {
		case "get-follows":
			var data []orm.ParamsList
			this.User.FollowingUsers().ValuesList(&data, "FollowUser__NickName", "FollowUser__UserName")
			result["success"] = true
			result["data"] = data

		case "follow", "unfollow":
			id, err := utils.StrTo(this.GetString("user")).Int()
			if err == nil && id != this.User.Id {
				fuser := models.User{Id: id}
				if action == "follow" {
					auth.UserFollow(&this.User, &fuser)
				} else {
					auth.UserUnFollow(&this.User, &fuser)
				}
				result["success"] = true
			}
		}
	}
}
