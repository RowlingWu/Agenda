package entity

import (
    "testing"
    "os"
    "encoding/json"
)

var name string

func TestReadCur_succs(t *testing.T) {
    f, _ := os.OpenFile("entity/curUser.txt", os.O_WRONLY, 0666)
    f.WriteString("bob")
    f.Close()

    users := [3]User{
        {"rowling","a","@","13"},
        {"bob","b","#","4"},
        {"alice","a","e","5"},
    }
    f, _ = os.OpenFile("entity/userInfo.txt", os.O_WRONLY, 0666)
    for _, u := range users {
        b, _ := json.Marshal(u)
        f.WriteString(string(b) + "\n")
    }
    f.Close()

    var flag int
    if name, flag = ReadCur(); flag == 0 {
        t.Log("read and clear current user info success")
    } else {
        t.Fatal("failed to read or clear current user info")
    }
}

func TestSeekUsr(t *testing.T) {
    ok := SeekUsr(name)
    if ok == true {
        t.Log("delete current user info success")
    } else {
        t.Fatal("failed to delete current user info")
    }
}

func TestReadCur_fail(t *testing.T) {
    if _, f := ReadCur(); f == 1 {
        t.Fatal("fatal: please log in first")
    } else {
        t.Log(f)
    }
}