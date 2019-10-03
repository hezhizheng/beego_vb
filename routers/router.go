package routers

import (
	"beego_vb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/:page", &controllers.PageController{})
	beego.Router("/category/:category_id", &controllers.MainController{},"get:Category")
	beego.Router("/pagination", &controllers.MainController{},"get:Pagination")
	beego.Router("/article/:slug", &controllers.VBController{})
}
