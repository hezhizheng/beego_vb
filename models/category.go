package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Categories struct {
	Id       int `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func init() {
	orm.RegisterModel(new(Categories))
}


func AllCategory() ([]orm.Params,  int64,  error)  {
	o := orm.NewOrm()

	var lists []orm.Params

	sql := "SELECT id,name, ( SELECT count(*) FROM posts WHERE posts.category_id = categories.id ) AS product_count FROM categories Order By product_count DESC"

	num, err  :=  o.Raw(sql).Values(&lists)
	return  lists,  num,  err
}

func CategoryPagination(page int, limit int , filter map[string]interface{} ) ([]orm.Params, int64 , int , int) {
	o  :=  orm.NewOrm()

	var lists []orm.Params

	type CategoryCount struct {
		Count int64
	}

	var count CategoryCount

	keyword := ""
	if _, ok := filter["keyword"]; ok {
		keyword = filter["keyword"].(string)
	}

	countSql := " SELECT COUNT(*) AS count FROM categories "
	countSql += " WHERE 1 = 1 "

	if keyword!= "" {
		countSql += " AND name LIKE '%" + keyword + "%' "
	}

	y := o.Raw(countSql).QueryRow(&count)

	if y !=nil{
	}

	offset := (page - 1) * limit

	sql := " SELECT id , name , created_at "
	sql += " FROM categories "
	sql += " WHERE 1 = 1 "

	if keyword!= "" {
		sql += " AND name LIKE '%" + keyword + "%' "
	}
	sql += " ORDER BY categories.id DESC "
	sql += " LIMIT ?,? "

	num, ValuesErr  :=  o.Raw(sql,offset,limit).Values(&lists)

	if ValuesErr == nil || num > 0{
	}

	return lists, count.Count,  page , limit
}


func CategoryCreate(name string ) int64  {

	o := orm.NewOrm()

	var categories Categories
	categories.Name = name
	id, _ := o.Insert(&categories)

	return id
}

func CategoryUpdate(data map[string]interface{} ) bool  {

	o := orm.NewOrm()

	categories := Categories{Id:data["category_id"].(int)}

	res := false

	if o.Read(&categories) == nil {
		categories.Name = data["name"].(string)
		if _, err := o.Update(&categories); err == nil {
			res = true
		}
	}

	return res
}

func CategoryDelete(id string ) bool  {

	o := orm.NewOrm()

	ret := false

	res, err := o.Raw("delete from categories  WHERE id = ?", id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			ret = true
		}
	}

	return ret
}


