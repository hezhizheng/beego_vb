package api

import (
    "beego_vb/models"
    "beego_vb/services"
    "beego_vb/util"
    "encoding/json"
    "github.com/astaxie/beego"
    "strconv"
)

type PostController struct {
    beego.Controller
}

func (c *PostController) Index() {
    filter := make(map[string]string)

    filter["keyword"] = c.GetString("keyword","")
    filter["category_id"] = c.GetString("category_id","0")

    formatFilter := services.FormatFilterData(filter)

    page , _ := strconv.Atoi(c.GetString("page","1"))
    limit , _ := strconv.Atoi(c.GetString("limit","10"))

    PLists, cnt, page, limit := models.PostPagination(page, limit,formatFilter)

    Posts := util.PageUtil(int(cnt), page, limit, PLists)

    var json = Response(Posts,Success,"")
    c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
    c.ServeJSON() // 声明返回json
}

func (c *PostController) Store() {
    data := make(map[string]interface{})

    json.Unmarshal(c.Ctx.Input.RequestBody, &data)

    //data["video_url"] = c.GetString("video_url","")
    //data["title"] = c.GetString("title","")
    //data["content"] = c.GetString("content","")
    //data["html_content"] = c.GetString("html_content","")
    //data["category_id"] = c.GetString("category_id")
    //data["video_cover"] = c.GetString("video_cover","")
    //data["slug"] = c.GetString("slug","")

    formatData := services.FormatSaveData(data)

    create := models.PostCreate(formatData)

    if create > 0 {
        var json = Response(true,Success,"")
        c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
        c.ServeJSON() // 声明返回json
    }

}

func (c *PostController) Update() {

    id := c.Ctx.Input.Param(":id")

    data := make(map[string]interface{})

    json.Unmarshal(c.Ctx.Input.RequestBody, &data)

    //data["video_url"] = c.GetString("video_url","")
    //data["title"] = c.GetString("title","")
    //data["content"] = c.GetString("content","")
    //data["html_content"] = c.GetString("html_content","")
    //data["category_id"] = c.GetString("category_id")
    //data["video_cover"] = c.GetString("video_cover","")
    //data["slug"] = c.GetString("slug")
    data["post_id"] = id

    formatData := services.FormatSaveData(data)

    update := models.PostUpdate(formatData)

    if update {
        var json = Response(update,Success,"")
        c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
        c.ServeJSON() // 声明返回json
    }

}

func (c *PostController) Destroy() {

    id := c.Ctx.Input.Param(":id")

    destroy := models.PostSoftDelete(id)

    if destroy {
		var json = Response(destroy,Success,"")
		c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
		c.ServeJSON() // 声明返回json
	}


}
