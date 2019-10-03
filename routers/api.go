package routers

import (
	"beego_vb/controllers/api"
	"beego_vb/util"
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"strings"
)

var FilterToken = func(ctx *context.Context) {
	// logs.Info("current router path is ", ctx.Request.RequestURI)
	ctx.ResponseWriter.Header().Set("Content-Type","application/json; charset=utf-8")
	Authorization := ctx.Input.Header("Authorization")
	if ctx.Request.RequestURI != "/api/admin/auth/login" && Authorization == "" {
		ctx.ResponseWriter.WriteHeader(401)
		var Response = api.Response(nil,401,"no permission")
		jsonBytes, _ := json.Marshal(Response)
		ctx.ResponseWriter.Write(jsonBytes)

	}
	if ctx.Request.RequestURI != "/api/admin/auth/login" && Authorization != "" {
		token := strings.Split(Authorization, "Bearer ")[1]
		validate , _ := util.ValidateToken(token)
		if !validate {
			var Response = api.Response(nil,401,"no permission")
			jsonBytes, _ := json.Marshal(Response)
			ctx.ResponseWriter.Write(jsonBytes)
		}
	}
}

func init() {

	beego.InsertFilter("/api/*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		//AllowOrigins:      []string{"https://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("api/admin/auth/login", &api.AuthController{},"post:Login")

	// 需要通过授权token的接口
	beego.InsertFilter("/api/admin/*",beego.BeforeRouter,FilterToken)

	beego.Router("api/admin/auth/me", &api.AuthController{},"get:Me")

	beego.Router("api/admin/dashboards", &api.IndexController{},"get:Dashboards")

	beego.Router("api/admin/posts",&api.PostController{},"get:Index;post:Store")
	beego.Router("api/admin/posts/:id",&api.PostController{},"put:Update")
	beego.Router("api/admin/posts/:id",&api.PostController{},"delete:Destroy")

	beego.Router("api/admin/pages",&api.PageController{},"get:Index;post:Store")
	beego.Router("api/admin/pages/:id",&api.PageController{},"put:Update")
	beego.Router("api/admin/pages/:id",&api.PageController{},"delete:Destroy")

	beego.Router("api/admin/categories-all",&api.CategoryController{},"get:AllCategories")
	beego.Router("api/admin/categories",&api.CategoryController{},"get:Index;post:Store")
	beego.Router("api/admin/categories/:id",&api.CategoryController{},"put:Update")
	beego.Router("api/admin/categories/:id",&api.CategoryController{},"delete:Destroy")
}
