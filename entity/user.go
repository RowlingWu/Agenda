package entity

import "os/user"

type User struct {
	Name string
	Passwd string
	Email string
	Tel string
}

var (
	UserInfo string
	CurUser string
	UserTemp string
)

func init() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	var homeDir string = u.HomeDir
	UserInfo = homeDir + "/userInfo.json"
	CurUser = homeDir + "/curUser.txt"
	UserTemp = homeDir + "/temp.json"
}