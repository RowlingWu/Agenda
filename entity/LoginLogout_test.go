package entity

import (
    "testing"
)

func Test_Login_1(t *testing.T) {
    if Login("hhh","123") == 2 {
        t.Log("Login test 1 accected")
    } else {
        t.Error("Login test 1 failed")
    }
}



func Test_Login_2(t *testing.T) {
    MyRegister("ww", "111", "a", "111")
    if Login("ww","111") == 0 {
        t.Log("Login test 2 accected")
    } else {
        t.Error("Login test 2 failed")
    }
}



func Test_Login_3(t *testing.T) {
    if Login("ww","111") ==1 {
        t.Log("Login test 3 accected")
    } else {
        t.Error("Login test 3 failed")
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
