package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Posts struct {
	Id       int
	VideoUrl string
	Title string
	Content string
	VideoCover string
	Slug string
	HtmlContent string
	PreviewImg string
	CreatedAt string
	IsDeleted int
	CategoryId int
	IsRecommend int
}

func init() {
	orm.RegisterModel(new(Posts))
}


func  PostCount() ( int64,  error)  {
	o  :=  orm.NewOrm()

	num, err := o.QueryTable("posts").Filter("is_deleted",0).Count() // SELECT COUNT(*) FROM USER

	return  num,  err
}


func  RecommendList() ([]orm.Params,  int64,  error)  {
	o  :=  orm.NewOrm()

	var lists []orm.Params

	// Values 返回 key => value
	num, err  :=  o.QueryTable(new(Posts)).Filter("is_recommend",1).Filter("is_deleted",0).Values(&lists)

	return  lists,  num,  err
}

func FindPostBySlug(slug string) (Posts, error) {
	o  :=  orm.NewOrm()

	var lists Posts

	 err  :=  o.QueryTable(new(Posts)).Filter("slug",slug).Filter("is_deleted",0).One(&lists)

	return lists ,err
}

func PostCountByCategoryId(categoryId string) int64 {

	o  :=  orm.NewOrm()

	type PostCount struct {
		Count int64
	}

	var count PostCount

	countSql := " SELECT COUNT(*) AS count FROM posts where category_id = " + categoryId

	y := o.Raw(countSql).QueryRow(&count)

	if y !=nil{
	}

	return count.Count
}

func PostPagination(page int, limit int , filter map[string]interface{} ) ([]orm.Params, int64 , int , int) {
	o  :=  orm.NewOrm()

	var lists []orm.Params

	type PostCount struct {
		Count int64
	}

	var count PostCount

	categoryId := 0
	if _, ok := filter["category_id"]; ok {
		categoryId = filter["category_id"].(int)
	}

	keyword := ""
	if _, ok := filter["keyword"]; ok {
		keyword = filter["keyword"].(string)
	}

	countSql := " SELECT COUNT(*) AS count FROM posts INNER JOIN categories ON categories.id = posts.category_id"
	countSql += " WHERE 1 = 1 "
	countSql += " AND is_deleted = 0 "

	if categoryId > 0 {
		countSql += " AND posts.category_id = " + strconv.Itoa(categoryId)
	}

	if keyword!= "" {
		countSql += " AND title LIKE '%" + keyword + "%' "
	}

	y := o.Raw(countSql).QueryRow(&count)

	if y !=nil{
	}

	offset := (page - 1) * limit

	sql := " SELECT category_id AS CategoryId, content AS Content, posts.id AS Id, posts.is_recommend AS IsRecommend, "
	sql += " slug AS Slug, title AS Title, video_cover AS VideoCover, video_url AS VideoUrl, categories.NAME AS CategoryName , posts.created_at AS CreatedAt, "
	sql += " html_content AS HtmlContent "
	sql += " FROM posts INNER JOIN categories ON categories.id = posts.category_id "
	sql += " WHERE 1 = 1 "
	sql += " AND is_deleted = 0 "
	if categoryId > 0 {
		sql += " AND posts.category_id = " + strconv.Itoa(categoryId)
	}
	if keyword!= "" {
		sql += " AND title LIKE '%" + keyword + "%' "
	}
	sql += " ORDER BY posts.id DESC "
	sql += " LIMIT ?,? "

	num, ValuesErr  :=  o.Raw(sql,offset,limit).Values(&lists)

	if ValuesErr == nil || num > 0{
	}

	return lists, count.Count,  page , limit
}


func PostCreate(data map[string]interface{} ) int64  {

	o := orm.NewOrm()

	var post Posts
	post.VideoUrl = data["video_url"].(string)
	post.VideoCover = data["video_cover"].(string)
	post.Content = data["content"].(string)
	post.HtmlContent = data["html_content"].(string)
	post.Title = data["title"].(string)
	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	post.CategoryId = data["category_id"].(int)
	post.Slug = data["slug"].(string)
	post.IsRecommend = data["is_recommend"].(int)
	id, _ := o.Insert(&post)

	return id
}

func PostUpdate(data map[string]interface{} ) bool  {

	o := orm.NewOrm()

	post := Posts{Id:data["post_id"].(int)}

	res := false

	if o.Read(&post) == nil {
		post.VideoUrl = data["video_url"].(string)
		post.VideoCover = data["video_cover"].(string)
		post.Content = data["content"].(string)
		post.HtmlContent = data["html_content"].(string)
		post.Title = data["title"].(string)
		post.CategoryId = data["category_id"].(int)
		post.Slug = data["slug"].(string)
		post.IsRecommend = data["is_recommend"].(int)
		if _, err := o.Update(&post); err == nil {
			res = true
		}
	}

	return res
}

func PostSoftDelete(id string ) bool  {

	o := orm.NewOrm()

	ret := false

	res, err := o.Raw("UPDATE posts SET is_deleted = ? WHERE id = ?", "1", id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			ret = true
		}
	}

	return ret
}
