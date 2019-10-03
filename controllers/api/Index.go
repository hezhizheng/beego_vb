package api

import (
    "beego_vb/models"
    "github.com/astaxie/beego"
)

type IndexController struct {
    beego.Controller
}

func (c *IndexController) Dashboards() {
    data := make(map[string]interface{})

    _,categoryCount,_ := models.AllCategory()

    postCount ,_ := models.PostCount()

    pageCount ,_ := models.PageCount()

    data["post_count"] = postCount
    data["category_count"] = categoryCount
    data["page_count"] = pageCount
    var _json = Response(data,Success,"")
    c.Data["json"] = _json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
    c.ServeJSON() // 声明返回json
}
