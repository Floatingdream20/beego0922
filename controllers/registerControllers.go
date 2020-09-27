package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"helloBeego922/db_mysql"
	"helloBeego922/models"
	"io/ioutil"
)
//该结构体用于处理/register请求
type RegisterControllers struct {
	beego.Controller
}
//该方法用于处理Post请求
func (r *RegisterControllers) Post() {
	//fmt.Println(r == nil)
	//fmt.Println(r.Ctx == nil)
	//fmt.Println(r.Ctx.Request == nil)
//接收前段传递的请求
	/*body,err := r.Ctx.Request.GetBody
	if err !=nil {
		r.Ctx.WriteString("数据接收错误。")
		return
	}*/
	//var user models.User
	//bodyBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
	//if err !=nil {
	//	r.Ctx.WriteString("数据接收错误。")
	//	return
	//}
	bodyBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
	//fmt.Println(err)
	if err != nil {
		r.Ctx.WriteString("数据接收错误,请重试")
		return
	}
	var user models.User
	err = json.Unmarshal(bodyBytes,&user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析错误")
		return
	}
	id,err := db_mysql.InsertUser(user)
	//fmt.Println(err)
	if err !=nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败。")
		return
	}
	fmt.Println(id)
	result := models.RespoonseResult{
		Code: 0,
		Message:"保存成功。",
		Data: nil,
	}
	r.Data["json"]=&result
	r.ServeJSONP()
	//fmt.Println(id)
    //fmt.Println(user.User)
    //fmt.Println(user.Birthday)
    //fmt.Println(user.Address)
    //fmt.Println(user.Nick)
    r.Ctx.WriteString("数据接收成功，并解析成功。")
}
