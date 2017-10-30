package entity

import (
    "testing"
    "os"
    "encoding/json"
)

func TestReadCur(t *testing.T) {
    var name string
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

    name = ReadCur()
    t.Log(name)
}

func TestClearCurUsr(t *testing.T) {
    ClearCurUsr()
}

func TestSeekUsr(t *testing.T) {
    SeekUsr("bob")
}
