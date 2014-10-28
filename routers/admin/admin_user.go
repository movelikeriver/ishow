package admin

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/movelikeriver/ishow/modules/auth"
	"github.com/movelikeriver/ishow/modules/models"
	"github.com/movelikeriver/ishow/modules/utils"
)

type UserAdminRouter struct {
	ModelAdminRouter
	object models.User
}

func (this *UserAdminRouter) Object() interface{} {
	return &this.object
}

func (this *UserAdminRouter) ObjectQs() orm.QuerySeter {
	return models.Users().RelatedSel()
}

// view for list model data
func (this *UserAdminRouter) List() {
	var users []models.User
	qs := models.Users()
	if err := this.SetObjects(qs, &users); err != nil {
		this.Data["Error"] = err
		beego.Error(err)
	}
}

// view for create object
func (this *UserAdminRouter) Create() {
	beego.Error("(this *UserAdminRouter) Create()")
	form := auth.UserAdminForm{Create: true}
	this.SetFormSets(&form)
}

// view for new object save
func (this *UserAdminRouter) Save() {
	beego.Error("(this *UserAdminRouter) Save()")
	form := auth.UserAdminForm{Create: true}
	if this.ValidFormSets(&form) == false {
		return
	}

	var user models.User
	form.SetToUser(&user)
	if err := user.Insert(); err == nil {
		this.FlashRedirect(fmt.Sprintf("/admin/user/%d", user.Id), 302, "CreateSuccess")
		return
	} else {
		beego.Error(err)
		this.Data["Error"] = err
	}
}

// view for edit object
func (this *UserAdminRouter) Edit() {
	beego.Error("(this *UserAdminRouter) Edit(")
	form := auth.UserAdminForm{}
	form.SetFromUser(&this.object)
	this.SetFormSets(&form)
}

// view for update object
func (this *UserAdminRouter) Update() {
	beego.Error("(this *UserAdminRouter) Update()")
	form := auth.UserAdminForm{Id: this.object.Id}
	if this.ValidFormSets(&form) == false {
		return
	}

	// get changed field names
	changes := utils.FormChanges(&this.object, &form)

	url := fmt.Sprintf("/admin/user/%d", this.object.Id)

	// update changed fields only
	if len(changes) > 0 {
		form.SetToUser(&this.object)
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
func (this *UserAdminRouter) Confirm() {
}

// view for delete object
func (this *UserAdminRouter) Delete() {
	if this.FormOnceNotMatch() {
		return
	}

	// delete object
	if err := this.object.Delete(); err == nil {
		this.FlashRedirect("/admin/user", 302, "DeleteSuccess")
		return
	} else {
		beego.Error(err)
		this.Data["Error"] = err
	}
}
