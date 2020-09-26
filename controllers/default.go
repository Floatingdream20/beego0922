package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"helloBeego922/models"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}
//默认请求为get，如果是默认请求，则会调用响应的方法，用Get来处理
func (c *MainController) Get() {
	//1、获取name、age、sex
	//2、做数据对比
	//c.GetString("user") 返回字符串
	//c.GetInt("psd") 返回整数
	UserName := c.Ctx.Input.Query("User")
	Password :=c.Ctx.Input.Query("Psd")
	if UserName !="Floatingdream20" && Password !="123456" {
		//代表
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误。"))
		return
	}
	//校验正确的情况
	c.Data["Website"] = "https://www.baidu.com"
	c.Data["Email"] = "astaxie@2161979329@qq.com"
	c.TplName = "index.tpl"
}
//编写一个post方法,用于处理post类型的请求
/*func (c *MainController) Post() {
	name :=c.Ctx.Request.FormValue("name")
	age :=c.Ctx.Request.FormValue("age")
	sex :=c.Ctx.Request.FormValue("sex")
	fmt.Println(name,age,sex)
	if name !="hua" && age != "666" && sex != "mail"{
		c.Ctx.WriteString("对不起，数据校验错误。")
		return
	}
	c.Ctx.WriteString("校验成功。")
}*/
func (c *MainController) Post(){
	var person models.Person
	dataBytes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil{
		c.Ctx.ResponseWriter.Write([]byte("错误"))
		return
	}
	err =json.Unmarshal(dataBytes,&person)
	if err !=nil {
		c.Ctx.WriteString("解析失败")
		return
	}
	fmt.Println("年龄",person.Name)
	fmt.Println("年龄",person.Age)
	fmt.Println("性别",person.Sex)
	c.Ctx.WriteString("成功")
}


