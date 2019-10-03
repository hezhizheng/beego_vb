package services

import (
    "github.com/teris-io/shortid"
    "github.com/russross/blackfriday"
    "strconv"
)

func FormatFilterData(data map[string]string) map[string]interface{} {

    res := make(map[string]interface{})
    categoryId := data["category_id"]

    categoryIdInt := 0
    if categoryId != ""{
        categoryIdInt , _ = strconv.Atoi(categoryId)
    }

    res["category_id"] = categoryIdInt
    res["keyword"] = data["keyword"]

    return res
}

func FormatSaveData(data map[string]interface{}) map[string]interface{} {

    res := make(map[string]interface{})
    categoryId := data["category_id"]
    postId := data["post_id"]
    content := data["content"]
    isRecommend := data["is_recommend"]
    slug := data["slug"]


    categoryIdInt := 0
    if categoryId != "" && categoryId != nil{
        categoryIdInt , _ = strconv.Atoi(categoryId.(string))
        res["category_id"] = categoryIdInt
        // 删除已经转换赋值的key
        delete(data,"category_id")
    }

    postIdInt := 0

    if postId != "" && postId != nil{
        postIdInt , _ = strconv.Atoi(postId.(string))
        res["post_id"] = postIdInt
        delete(data,"post_id")
    }

    isRecommendInt := 0
    if isRecommend != "" && isRecommend != nil{
        isRecommendInt , _ = strconv.Atoi(isRecommend.(string))
        res["is_recommend"] = isRecommendInt
        delete(data,"is_recommend")
    }

    if content != "" && content != nil{
        output := blackfriday.Run([]byte(content.(string)))
        htmlContent := string(output)
        res["html_content"] = htmlContent
        delete(data,"html_content")
    }


    if slug == "" || slug == nil {
        shortid , _ := shortid.Generate()
        res["slug"] = shortid
        delete(data,"slug")
    }

    for k,_ := range data {
        res[k] = data[k]
    }

    return res
}
