package controllers

import (
	"beego_vb/models"
	"github.com/astaxie/beego"
)

type PageController struct {
	beego.Controller
}

func (c *PageController) Get() {

	// 获取路由参数 :page
	name := c.Ctx.Input.Param(":page")

	page , _ := models.FindPageByName(name)

	if page.Id == 0 {
		c.Abort("404")
	}

	PageAll, _ := models.PageAll()

	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footer"] = "footer.tpl"
	c.LayoutSections["nav"] = "nav.tpl"
	c.LayoutSections["content"] = "page.tpl"

	c.TplName = "page.tpl"

	c.Data["Page"] = page
	c.Data["PageAll"] = PageAll

}
