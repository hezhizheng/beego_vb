package main

import (
	_ "beego_vb/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.Debug = false // 是否开启调试模式 调试模式下会打印出sql语句
	orm.RegisterDriver("mysql",  orm.DRMySQL)

	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPwd := beego.AppConfig.String("mysqlpass")
	mysqlHost := beego.AppConfig.String("mysqlurls")
	mysqlDb := beego.AppConfig.String("mysqldb")
	mysqlPort := beego.AppConfig.String("mysqlport")

	// "admin:admin@tcp(127.0.0.1:3306)/vb?charset=utf8mb4"
	dataSource := mysqlUser + ":" + mysqlPwd + "@tcp" + "(" + mysqlHost + ":" + mysqlPort +  ")/" + mysqlDb + "?charset=utf8mb4"
	orm.RegisterDataBase("default", "mysql", dataSource, 30)
}

func main() {
	// 设置vue静态访问映射路径
	vueStaticDir := beego.AppConfig.String("vueStaticDir")
	beego.SetStaticPath("/vue-admin", vueStaticDir+"index.html")
	beego.SetStaticPath("/vue-admin/css",vueStaticDir+"css")
	beego.SetStaticPath("/vue-admin/fonts",vueStaticDir+"fonts")
	beego.SetStaticPath("/vue-admin/img",vueStaticDir+"img")
	beego.SetStaticPath("/vue-admin/js",vueStaticDir+"js")

	beego.Run()
}

