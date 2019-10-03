package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id  int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
}

func init() {
	orm.RegisterModel(new(Users))
}
