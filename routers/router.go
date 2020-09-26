package routers

import (
	"helloBeego922/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register",&controllers.RegisterControllers{})
    beego.Router("/", &controllers.MainController{})
}
