package entity

import (
    "log"
    "os"
    "bufio"
    "encoding/json"
)

func Logout() bool {
    dat, err := os.Open("entity/curUser.txt")
    check(err)
    line := bufio.NewScanner(dat)
    line.Scan()
    if len(line.Text()) == 0 {
        dat.Close()
        return false
    }
    dat.Close()
    os.Create("entity/curUser.txt")
    return true
}

func Login(username string, password string) bool {
    dat2, err2 := os.Open("entity/curUser.txt")
    check(err2)
    line2 := bufio.NewScanner(dat2)
    line2.Scan()
    if len(line2.Text()) != 0 {
        dat2.Close()
        log.Fatal("login failed. Already logged in")
        return false
    }

    var user User
    dat, err := os.Open("entity/userInfo.json")
    check(err)
    line := bufio.NewScanner(dat)
    for line.Scan() {
        json.Unmarshal([]byte(line.Text()), &user);
        if user.Name == username && user.Passwd == password {
    	    dat.Close()

            dat, err := os.OpenFile("entity/curUser.txt", os.O_WRONLY|os.O_CREATE, 0666)
            check(err)
            dat.WriteString(username)
            dat.Close()

            return true
        }
    }
    dat.Close()
    log.Fatal("login failed. The username or password incorrect")
    return false
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}
