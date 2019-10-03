package controllers

import (
	"beego_vb/models"
	"beego_vb/services"
	"beego_vb/util"
	"fmt"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	// 1、获取用户传来的pwd,使用字节切片转换
	originPwd := []byte("666666") // 用户传的实际密码值

	// 2、调用 bcrypt.GenerateFromPassword 生成加密字符串
	hashPwd, _ := bcrypt.GenerateFromPassword(originPwd, bcrypt.DefaultCost)

	// 3、此时 hashPwd 为字节切片，实际加密字符串需使用string转换
	fmt.Println(string(hashPwd)) // 可以将此加密串保存到数据库，可作为密码匹配验证

	RecommendList, _, _ := models.RecommendList()

	PageAll, _ := models.PageAll()

	PLists, cnt, page, limit := models.PostPagination(1, 8,nil)

	Posts := util.PageUtil(int(cnt), page, limit, PLists)

	cList, _, _ := models.AllCategory()

	c.Layout = "layout.tpl"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footer"] = "footer.tpl"
	c.LayoutSections["nav"] = "nav.tpl"
	c.LayoutSections["content"] = "index.tpl"


	c.TplName = "index.tpl"

	c.Data["recommend"] = RecommendList
	c.Data["categories"] = cList
	c.Data["Posts"] = Posts
	c.Data["PageAll"] = PageAll

}

func (c *MainController) Pagination() {
	p := c.GetString("p","1")
	l := c.GetString("l","8")
	categoryId := c.GetString("category_id","0")
	intP, _ := strconv.Atoi(p)
	intL, _ := strconv.Atoi(l)

	filter := make(map[string]string)
	filter["category_id"] = categoryId
	formatFilter := services.FormatFilterData(filter)

	PLists, cnt,  _, limit := models.PostPagination(intP, intL, formatFilter)

	Posts := util.PageUtil(int(cnt), intP, limit, PLists)


	c.Data["json"] = Posts // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
	c.ServeJSON() // 声明返回json
}

func (c *MainController) Category() {
	// 获取路由参数 :slug
	CategoryId := c.Ctx.Input.Param(":category_id")

	RecommendList, _, _ := models.RecommendList()

	PageAll, _ := models.PageAll()

	filter := make(map[string]string)
	filter["category_id"] = CategoryId
	formatFilter := services.FormatFilterData(filter)

	PLists, cnt, page, limit := models.PostPagination(1, 8,formatFilter)

	Posts := util.PageUtil(int(cnt), page, limit, PLists)

	cList, _, _ := models.AllCategory()

	c.Layout = "layout.tpl"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footer"] = "footer.tpl"
	c.LayoutSections["nav"] = "nav.tpl"
	c.LayoutSections["content"] = "index.tpl"

	c.TplName = "index.tpl"

	c.Data["recommend"] = RecommendList
	c.Data["categories"] = cList
	c.Data["Posts"] = Posts
	c.Data["PageAll"] = PageAll
}
