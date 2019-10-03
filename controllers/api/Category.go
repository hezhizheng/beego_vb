package api

import (
    "beego_vb/models"
    "beego_vb/util"
    "encoding/json"
    "github.com/astaxie/beego"
    "strconv"
)

type CategoryController struct {
    beego.Controller
}

// AllCategories
func (c *CategoryController) AllCategories() {


    categories ,_ ,_ := models.AllCategory()

    var _json = Response(categories,Success,"")
    c.Data["json"] = _json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
    c.ServeJSON() // 声明返回json

}

func (c *CategoryController) Index() {

    filter := make(map[string]interface{})

    filter["keyword"] = c.GetString("keyword","")

    page , _ := strconv.Atoi(c.GetString("page","1"))
    limit , _ := strconv.Atoi(c.GetString("limit","10"))

    PLists, cnt, page, limit := models.CategoryPagination(page, limit,filter)

    Categories := util.PageUtil(int(cnt), page, limit, PLists)

    var json = Response(Categories,Success,"")
    c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
    c.ServeJSON() // 声明返回json

}

func (c *CategoryController) Store() {

    reqData := make(map[string]interface{})

    json.Unmarshal(c.Ctx.Input.RequestBody, &reqData)

    create := models.CategoryCreate(reqData["name"].(string))

    if create > 0 {
        var json = Response(true,Success,"")
        c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
        c.ServeJSON() // 声明返回json
    }

}

func (c *CategoryController) Update() {

    id := c.Ctx.Input.Param(":id")

    data := make(map[string]interface{})

    json.Unmarshal(c.Ctx.Input.RequestBody, &data)

    intId ,_ := strconv.Atoi(id)

    data["category_id"] = intId

    update := models.CategoryUpdate(data)

    if update {
        var json = Response(update,Success,"")
        c.Data["json"] = json
        c.ServeJSON()
    }

}

func (c *CategoryController) Destroy() {

    id := c.Ctx.Input.Param(":id")

    postCount := models.PostCountByCategoryId(id)

    if postCount > 0 {
        var json = Response(false,7000,"存在文章，不能删除")
        c.Data["json"] = json
        c.ServeJSON()
    }

    if postCount == 0 {
        destroy := models.CategoryDelete(id)

        if destroy {
            var json = Response(destroy,Success,"")
            c.Data["json"] = json
            c.ServeJSON()
        }
    }



}
