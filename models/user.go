package models

/*type User struct {
	Name string   //User string `json:"name"`
	Birthday string
	Address string
	Nick string
}*/

type User struct {
	User string `json:"user"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	Nick string `json:"nick"`
}