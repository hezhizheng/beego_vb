package api

import (
    "beego_vb/models"
    "beego_vb/services"
    "beego_vb/util"
    "encoding/json"
    "github.com/astaxie/beego"
    "strconv"
)

type PageController struct {
    beego.Controller
}

func (c *PageController) Index() {

    page , _ := strconv.Atoi(c.GetString("page","1"))
    limit , _ := strconv.Atoi(c.GetString("limit","10"))

    PLists, cnt, page, limit := models.PagePagination(page, limit)

    Pages := util.PageUtil(int(cnt), page, limit, PLists)

    var _json = Response(Pages,Success,"")
    c.Data["json"] = _json
    c.ServeJSON()
}

func (c *PageController) Store() {
    data := make(map[string]interface{})

    json.Unmarshal(c.Ctx.Input.RequestBody, &data)

    saveData := services.PageFormatSaveData(data)

    create := models.PageCreate(saveData)

    if create > 0 {
        var json = Response(true,Success,"")
        c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
        c.ServeJSON() // 声明返回json
    }

}

func (c *PageController) Update() {

    id := c.Ctx.Input.Param(":id")

    idInt , _ := strconv.Atoi(id)

    data := make(map[string]interface{})

    json.Unmarshal(c.Ctx.Input.RequestBody, &data)

    data["page_id"] = idInt

    saveData := services.PageFormatSaveData(data)

    update := models.PageUpdate(saveData)

    if update {
        var json = Response(update,Success,"")
        c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
        c.ServeJSON() // 声明返回json
    }

}

func (c *PageController) Destroy() {

    id := c.Ctx.Input.Param(":id")

    destroy := models.PageDelete(id)

    if destroy {
		var json = Response(destroy,Success,"")
		c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
		c.ServeJSON() // 声明返回json
	}


}
