package services

import (
    "beego_vb/models/filter"
    "golang.org/x/crypto/bcrypt"
    "strconv"
)

func Login(email string, pwd string) (bool, map[string]string) {
	user, num := filter.FindUserByEmail(email)
    bytePwd := []byte(pwd)

    data := make(map[string]string)

	if num > 0 {

	    userPwd := []byte(user[0]["Password"].(string))

        err := bcrypt.CompareHashAndPassword(userPwd, bytePwd)

        if err != nil{
            return false, data
        }

       data["id"] = strconv.FormatInt(user[0]["Id"].(int64),10) // 取值需要确定每一个值的类型！！！
       data["name"] = user[0]["Name"].(string)
       return true, data
   }else{
       // 抛错
       return false , data
   }
}
