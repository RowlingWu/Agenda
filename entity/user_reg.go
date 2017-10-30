package entity
import (
  "os"
  "io"
  "encoding/json"
  "log"
)
var userPath = "entity/userInfo.json"

var dirty bool

var userList []User

func init() {
  dirty = false
  MyRead();
}
func MyRead() error {
   file, err := os.Open(userPath);
   if err != nil {
     log.Fatal(err.Error())
     return err
   }

   defer file.Close()
   js := json.NewDecoder(file)
   switch err := js.Decode(&userList); err {
   case nil, io.EOF:
     return nil
   default:
     log.Fatal("Decode userinfo file failed:", err)
     return err
   }
}
func MyWrite() error {
  file,err := os.Create("userPath");
  if err != nil {
    return err
  }
  defer file.Close()
  js :=json.NewEncoder(file)
  if err:=js.Encode(&userList);err!=nil {
    log.Fatal("writeJSON:",err)
    return err
  }
  return nil
}

func MyCreate(newusr *User) {
  userList = append(userList, *newusr)
  dirty = true
}
func GetAllUsers() ([]User) {
   return userList
}
func MyUsrQuery(u string) (bool) {
  for _, i := range userList {
    if u == i.Name {
      log.Fatal("has been used")
      return false
    }
  }
  return true
}
func MyRegister(u string, pw string, em string, ph string) (bool,error) {
    flag := MyUsrQuery(u)
   if flag == false {
     log.Fatal("This username has been registered already, please choose annother one")
     return false, nil
   }
   MyCreate(&User{u,pw,em,ph})
   MyWrite();
   return true, nil
}
