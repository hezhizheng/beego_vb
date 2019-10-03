package api

import (
	"beego_vb/services"
	"beego_vb/util"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {

	reqData := make(map[string]interface{})

	json.Unmarshal(c.Ctx.Input.RequestBody, &reqData)

	login , data := services.Login(reqData["email"].(string),reqData["password"].(string))

	if login{
		token := util.GenerateToken(data,0)
		data["token"] = token
		var json = Response(data,Success,"")
		c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
		c.ServeJSON() // 声明返回json
	}else{
		var json = Response(data,6000,"账号或密码错误！")
		c.Data["json"] = json // 格式一定为 c.Data["json"] ！！！ 必须指定为 json
		c.ServeJSON() // 声明返回json
	}
}

func (c *AuthController) Me() {
	Authorization := c.Ctx.Request.Header["Authorization"]
	token := strings.Split(Authorization[0], "Bearer ")[1]
	user := util.GetUser(token)

	data := make(map[string]string)
	data["id"] = strconv.FormatInt(user[0]["Id"].(int64),10)
	data["name"] = user[0]["Name"].(string)
	data["email"] = user[0]["Email"].(string)
	data["avatar"] = user[0]["Avatar"].(string)

	var json = Response(data,Success,"")
	c.Data["json"] = json
	c.ServeJSON()
}
