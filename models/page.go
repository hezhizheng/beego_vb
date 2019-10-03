package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Pages struct {
	Id       int `json:"id"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	Content string `json:"content"`
	HtmlContent string `json:"html_content"`
	CreatedAt string `json:"created_at"`
	SortOrder float64 `json:"sort_order"`
}

func init() {
	orm.RegisterModel(new(Pages))
}

func  PageCount() ( int64,  error)  {
	o  :=  orm.NewOrm()

	num, err := o.QueryTable("pages").Count() // SELECT COUNT(*) FROM USER

	return  num,  err
}

func PageAll() ([]orm.Params, int64) {
	o  :=  orm.NewOrm()

	var lists []orm.Params

	num , _:=  o.QueryTable(new(Pages)).OrderBy("sort_order").Values(&lists)

	return lists ,num
}

func FindPageByName(name string) (Pages, error) {
	o  :=  orm.NewOrm()

	var lists Pages

	 err  :=  o.QueryTable(new(Pages)).Filter("name",name).One(&lists)

	return lists ,err
}


func PagePagination(page int, limit int ) ([]orm.Params, int64 , int , int) {
	o  :=  orm.NewOrm()

	var lists []orm.Params

	type PageCount struct {
		Count int64
	}

	var count PageCount

	countSql := " SELECT COUNT(*) AS count FROM pages "

	y := o.Raw(countSql).QueryRow(&count)

	if y !=nil{
	}

	offset := (page - 1) * limit

	sql := " SELECT * "
	sql += " FROM pages "
	sql += " ORDER BY pages.id DESC "
	sql += " LIMIT ?,? "

	num, ValuesErr  :=  o.Raw(sql,offset,limit).Values(&lists)

	if ValuesErr == nil || num > 0{
	}

	return lists, count.Count,  page , limit
}


func PageCreate(data map[string]interface{} ) int64  {

	o := orm.NewOrm()

	var page Pages
	page.DisplayName = data["display_name"].(string)
	page.Content = data["content"].(string)
	page.HtmlContent = data["html_content"].(string)
	page.Name = data["name"].(string)
	page.SortOrder = data["sort_order"].(float64)
	page.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	id, _ := o.Insert(&page)

	return id
}

func PageUpdate(data map[string]interface{} ) bool  {

	o := orm.NewOrm()

	page := Pages{Id:data["page_id"].(int)}

	res := false

	if o.Read(&page) == nil {
		page.DisplayName = data["display_name"].(string)
		page.Name = data["name"].(string)
		page.Content = data["content"].(string)
		page.HtmlContent = data["html_content"].(string)
		page.SortOrder = data["sort_order"].(float64)
		if _, err := o.Update(&page); err == nil {
			res = true
		}
	}

	return res
}


func PageDelete(id string ) bool  {

	o := orm.NewOrm()

	ret := false

	res, err := o.Raw("delete from pages  WHERE id = ?", id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			ret = true
		}
	}

	return ret
}
