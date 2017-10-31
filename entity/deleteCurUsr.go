package entity

import (
    "os"
    "bufio"
    "encoding/json"
)

var curUser string = "entity/curUser.txt"
var userInfo string = "entity/userInfo.json"
var userTemp string = "entity/temp.json"

func ReadCur() (string, int) {  // 0:成功. 1:需要登录. 2:失败
    var name string
    var cur *os.File
    var err error
    if cur, err = os.Open(curUser); err != nil {
        return string(""), 2
    }
    line := bufio.NewScanner(cur)
    line.Scan()
    name = line.Text()
    cur.Close()
    if (len(name) == 0) {
        return string(""), 1
    }

    cur, err = os.Create(curUser)
    if err == nil {
        cur.Close()
    } else {
        return err.Error(), 2
    }

    return name, 0
}

func SeekUsr(name string) bool {
    file, err := os.Open(userInfo)
    if err != nil {
        return false
    }

    // begin searching
    // ftemp copies userInfo.txt, deletes curUser and is renamed as userInfo.txt
    ftemp, err := os.OpenFile(userTemp, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        return false
    }
    line := bufio.NewScanner(file)
    var user User
    for line.Scan() {
        if err = json.Unmarshal([]byte(line.Text()), &user); err == nil {
            if user.Name != name {
                ftemp.WriteString(line.Text() + "\n")
            }
        }
    }
    file.Close()
    ftemp.Close()
    if err = os.Remove(userInfo); err != nil {
        return false
    }
    if err = os.Rename(userTemp, userInfo); err != nil {
        return false
    }
    return true
}
