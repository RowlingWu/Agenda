package entity

import (
    "log"
    "os"
    "bufio"
    "encoding/json"
)

func ReadCur() string {
    var name string
    var cur *os.File
    var err error
    if cur, err = os.Open("entity/curUser.txt"); err != nil {
        log.Fatal("failed to read info of the current user")
    }
    line := bufio.NewScanner(cur)
    line.Scan()
    name = line.Text()
    cur.Close()
    if (len(name) == 0) {
        log.Fatal("fatal: Please login first")
    }
    return name
}

func ClearCurUsr() {
    cur, err := os.Create("entity/curUser.txt")
    if err == nil {
        cur.Close()
    } else {
        log.Fatal("failed to delete current user")
    }
}

func SeekUsr(name string) {
    file, err := os.Open("entity/userInfo.txt")
    if err != nil {
        log.Fatal("failed to delete current user")
    }

    // begin searching
    // ftemp copies userInfo.txt, deletes curUser and is renamed as userInfo.txt
    ftemp, err := os.OpenFile("entity/temp.txt", os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        log.Fatal(err.Error())
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
    if err = os.Remove("entity/userInfo.txt"); err != nil {
        log.Fatal("failed to delete current user")
    }
    if err = os.Rename("entity/temp.txt", "entity/userInfo.txt"); err != nil {
        log.Fatal("failed to delete current user")
    }
}
