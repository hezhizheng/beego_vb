package filter

import "github.com/astaxie/beego/orm"

func FindUserByEmail(email string ) ([]orm.Params,  int64)  {
    o := orm.NewOrm()

    var lists []orm.Params

    num, _ := o.QueryTable("users").Filter("email", email).Values(&lists)

    return  lists,  num
}

func FindUserByEmailAndPwd(email string , pwd string) ([]orm.Params,  int64)  {
    o := orm.NewOrm()

    var lists []orm.Params

    num, _ := o.QueryTable("users").Filter("email", email).Filter("password", pwd).Values(&lists)

    return  lists,  num
}

func FindUserById(id string ) ([]orm.Params,  int64)  {
    o := orm.NewOrm()

    var lists []orm.Params

    num, _ := o.QueryTable("users").Filter("id", id).Values(&lists)

    return  lists,  num
}
