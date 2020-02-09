package controllers

type ExitController struct {
	BaseController
}

func (this *ExitController) Get() {
	this.DelSession("loginuser")
	this.Redirect("/", 302)
}
