package main

import (
	"github.com/astaxie/beego"
	_"helloBeego922/db_mysql"
	_ "helloBeego922/routers"
)

func main() {
	//appName := config.String("appname")
	//fmt.Println("应用名称：",appName)
	//port,err :=config.Int("httpport")
	//if err != nil {
	//	panic("项目配置文件解析失败，请检查配置文件。")
	//}
	//fmt.Println(post)
	beego.Run()
}

