package controllers

import (
	"beego_vb/models"
	"github.com/astaxie/beego"
)

type VBController struct {
	beego.Controller
}

func (c *VBController) Get() {

	// 获取路由参数 :slug
	slug := c.Ctx.Input.Param(":slug")

	post , _ := models.FindPostBySlug(slug)

	if post.Id == 0 {
		c.Abort("404")
	}


	PageAll, _ := models.PageAll()

	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footer"] = "footer.tpl"
	c.LayoutSections["nav"] = "nav.tpl"
	c.LayoutSections["content"] = "article.tpl"

	c.TplName = "article.tpl"

	c.Data["PostDetail"] = post
	c.Data["PageAll"] = PageAll

}
