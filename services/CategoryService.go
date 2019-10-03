package services

import (
    "strconv"
)

func CategoryFormatFilterData(data map[string]string) map[string]interface{} {

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
