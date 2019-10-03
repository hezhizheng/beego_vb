package filter

import "github.com/astaxie/beego/orm"


func CategoryFilterById(id string ) ([]orm.Params,  int64)  {
    o := orm.NewOrm()

    var lists []orm.Params

    num, _ := o.QueryTable("categories").Filter("id", id).Values(&lists)

    return  lists,  num
}
