package entity

import (
    //"log"
    "os"
    "bufio"
    "encoding/json"
)

func Logout() bool {
    dat, err := os.Open(CurUser)
    check(err)
    line := bufio.NewScanner(dat)
    line.Scan()
    if len(line.Text()) == 0 {
        dat.Close()
        return false
    }
    dat.Close()
    os.Create(CurUser)
    return true
}

func Login(username string, password string) int {
    dat2, err2 := os.Open(CurUser)
    check(err2)
    line2 := bufio.NewScanner(dat2)
    line2.Scan()
    if len(line2.Text()) != 0 {
        dat2.Close()
        //log.Fatal("login failed. Already logged in")
        return 1
    }

    var user User
    dat, err := os.Open(UserInfo)
    check(err)
    line := bufio.NewScanner(dat)
    for line.Scan() {
        json.Unmarshal([]byte(line.Text()), &user);
        if user.Name == username && user.Passwd == password {
    	    dat.Close()

            dat, err := os.OpenFile(CurUser, os.O_WRONLY|os.O_CREATE, 0666)
            check(err)
            dat.WriteString(username)
            dat.Close()

            return 0
        }
    }
    dat.Close()
    //log.Fatal("login failed. The username or password incorrect")
    return 2
}

func check(e error) {
    if e != nil {
        os.Create("entity/curUser.txt")
    }
}
