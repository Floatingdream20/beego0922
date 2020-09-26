package models

/*type User struct {
	Name string   //User string `json:"name"`
	Birthday string
	Address string
	Nick string
}*/

type User struct {
	User string `form:"name"`
	Birthday string `form:"birthday"`
	Address string `form:"address"`
	Nick string `form:"nick"`
}