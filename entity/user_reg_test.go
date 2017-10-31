package entity

import (
	"testing"

)

var users = []User{
	{"u1", "up", "u1@q", "123"},
	{"u2", "ua", "u2@q", "456"},
	{"u3", "ua", "u3@q", "789"},
}

func init() {
	userList = nil
}

func TestMyCreate(t *testing.T) {

	var user1 User
	user1 = users[0]
	MyCreate(&user1)

	length := len(userList)
	if userList[length-1].Name != user1.Name || userList[length-1].Tel != user1.Tel ||userList[length-1].Email != user1.Email ||userList[length-1].Passwd != user1.Passwd {
		t.Error("some errors happen in create\n")
	} else {
		t.Log("check for MyCreate pass\n")
	}
}
