package services

import (
    "github.com/russross/blackfriday"
)

func PageFormatSaveData(data map[string]interface{}) map[string]interface{} {

    res := make(map[string]interface{})
    content := data["content"]
    //sortOrder := data["sort_order"]


    //sortOrderInt := 0
    //if sortOrder != "" && sortOrder != nil{
    //    sortOrderInt , _ = strconv.Atoi(sortOrder.(string))
    //    res["sort_order"] = sortOrderInt
    //    // 删除已经转换赋值的key
    //    delete(data,"sort_order")
    //}

    if content != "" && content != nil{
        output := blackfriday.Run([]byte(content.(string)))
        htmlContent := string(output)
        res["html_content"] = htmlContent
        delete(data,"html_content")
    }

    for k,_ := range data {
        res[k] = data[k]
    }

    return res
}
