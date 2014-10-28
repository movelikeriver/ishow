package admin

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/movelikeriver/ishow/modules/models"
	"github.com/movelikeriver/ishow/modules/post"
	"github.com/movelikeriver/ishow/modules/utils"
)

type TopicAdminRouter struct {
	ModelAdminRouter
	object models.Topic
}

func (this *TopicAdminRouter) Object() interface{} {
	return &this.object
}

func (this *TopicAdminRouter) ObjectQs() orm.QuerySeter {
	return models.Topics().RelatedSel()
}

// view for list model data
func (this *TopicAdminRouter) List() {
	var topics []models.Topic
	qs := models.Topics().RelatedSel()
	if err := this.SetObjects(qs, &topics); err != nil {
		this.Data["Error"] = err
		beego.Error(err)
	}
}

// view for create object
func (this *TopicAdminRouter) Create() {
	form := post.TopicAdminForm{Create: true}
	this.SetFormSets(&form)
}

// view for new object save
func (this *TopicAdminRouter) Save() {
	form := post.TopicAdminForm{Create: true}
	if this.ValidFormSets(&form) == false {
		return
	}

	var topic models.Topic
	form.SetToTopic(&topic)
	if err := topic.Insert(); err == nil {
		this.FlashRedirect(fmt.Sprintf("/admin/topic/%d", topic.Id), 302, "CreateSuccess")
		return
	} else {
		beego.Error(err)
		this.Data["Error"] = err
	}
}

// view for edit object
func (this *TopicAdminRouter) Edit() {
	form := post.TopicAdminForm{}
	form.SetFromTopic(&this.object)
	this.SetFormSets(&form)
}

// view for update object
func (this *TopicAdminRouter) Update() {
	form := post.TopicAdminForm{Id: this.object.Id}
	if this.ValidFormSets(&form) == false {
		return
	}

	// get changed field names
	changes := utils.FormChanges(&this.object, &form)

	url := fmt.Sprintf("/admin/topic/%d", this.object.Id)

	// update changed fields only
	if len(changes) > 0 {
		form.SetToTopic(&this.object)
		if err := this.object.Update(changes...); err == nil {
			this.FlashRedirect(url, 302, "UpdateSuccess")
			return
		} else {
			beego.Error(err)
			this.Data["Error"] = err
		}
	} else {
		this.Redirect(url, 302)
	}
}

// view for confirm delete object
func (this *TopicAdminRouter) Confirm() {
}

// view for delete object
func (this *TopicAdminRouter) Delete() {
	if this.FormOnceNotMatch() {
		return
	}

	// delete object
	if err := this.object.Delete(); err == nil {
		this.FlashRedirect("/admin/topic", 302, "DeleteSuccess")
		return
	} else {
		beego.Error(err)
		this.Data["Error"] = err
	}
}
