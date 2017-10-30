package entity

import (
    "testing"
)

func Test_Login_1(t *testing.T) {
    if Login("annie","123") {
        t.Error("Login test 1 failed")
    } else {
        t.Log("Login test 1 accected")
    }
}



func Test_Login_2(t *testing.T) {
    if Login("rowling","a") {
        t.Log("Login test 2 accected")
    } else {
        t.Error("Login test 2 failed")
    }
}



func Test_Login_3(t *testing.T) {
    if Login("rowling","a") {
        t.Error("Login test 3 failed")
    } else {
        t.Log("Login test 3 accected")
    }
}



func Test_Logout_1(t *testing.T) {
    if Logout() {
        t.Log("Logout test 1 accected")
    } else {
        t.Error("Login test 1 failed")
    }
}



func Test_Logout_2(t *testing.T) {
    if Logout() {
        t.Error("Login test 2 failed")
    } else {
        t.Log("Logout test 2 accected")
    }
}
