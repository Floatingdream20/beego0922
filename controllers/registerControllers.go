package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"helloBeego922/models"
)
//该结构体用于处理/register请求
type RegisterControllers struct {
	beego.Controller
}
//该方法用于处理Post请求
func (r *RegisterControllers) Post() {

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
	//err =json.Unmarshal(bodyBytes,&user)
	//if err !=nil {
	//	r.Ctx.WriteString("数据解析错误。")
	//	return
	//}
	var user models.User
	err :=r.ParseForm(&user)
	if err !=nil {
		r.Ctx.WriteString("数据解析错误。")
	}

	fmt.Println(user.User)
	fmt.Println(user.Birthday)
	fmt.Println(user.Address)
	fmt.Println(user.Nick)
	r.Ctx.WriteString("数据接收成功，并解析成功。")
}
