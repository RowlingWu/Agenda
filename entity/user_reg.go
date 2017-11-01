package entity
import (
  "os"
//"io"
  "encoding/json"
 //"log"
 "bufio"

)

//var UserInfo string = "entity/UserInfo.json"
var dirty bool

var userList []User

func init() {
  dirty = false
  MyRead();
}
/*func MyRead() error {
   file, err := os.Open("entity/UserInfo.json");
   if err != nil {
     log.Fatal(err.Error() + "reg")
     return err
   }

   defer file.Close()
   js := json.NewDecoder(file)
   switch err := js.Decode(&userList); err {
   case nil, io.EOF:
     return nil
   default:
    log.Fatal("Decode UserInfo file failed:", err)
     return err
   }
}*/
/*func MyWrite() error {
  file,err := os.Create("entity/UserInfo.json");
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
}*/
func MyRead() error {
  file, err := os.Open(UserInfo);
  if err != nil {
    //log.Fatal(err.Error() + "reg")
    return err
   }

   defer file.Close()

   line := bufio.NewScanner(file)
   count := 0
   //length :=len(userList)
   //fmt.Printf("size of list: %d\n",length)
   var temp User
   for line.Scan() {
    //fmt.Println("right read")
    if err = json.Unmarshal([]byte(line.Text()),&temp); err == nil {
      userList = append(userList, temp)
      //fmt.Printf("name is: %s and read name: %s\n",temp.Name,userList[count].Name)
      count = count +1
    }
   }
   return err
}
func MyWrite() error {
  file,err := os.Create(UserInfo);
  if err != nil {
    return err
  }
  defer file.Close()
  for _,i := range userList {
      js, err := json.Marshal(i)
      if err != nil {
        //log.Fatal("encoding failed")
      } else {
        file.Write(js)
file.WriteString("\r\n")
      }
  }
  return err
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
      //log.Fatal("The name: %s has been used.",u)
      return false
    }
  }
  return true
}
func MyRegister(u string, pw string, em string, ph string) (bool,error) {
    flag := MyUsrQuery(u)
   if flag == false {
    //log.Fatal("This username has been registered already, please choose annother one")
     return false, nil
   }
   MyCreate(&User{u,pw,em,ph})
   MyWrite();
   return true, nil
}
