package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Userlist struct {
	List []User `json:"list"`
	ID   string `json:"sid"`
}

var Ul Userlist
var users []User
var currentUser User

//|os.O_APPEND tian jia
func adduser(s User) {
	file, _ := os.OpenFile("database.txt", os.O_RDWR|os.O_CREATE, 0664)
	defer file.Close()

	// ReadAll接收一个io.Reader的参数 返回字节切片
	bytes, _ := ioutil.ReadAll(file)
	// fmt.Println(string(bytes))
	var u Userlist
	json.Unmarshal(bytes, &u)
	u.ID = "Base"
	u.List = append(u.List, User{s.Name, s.Password, s.Email, s.Phone})
	fmt.Println("Write to database")
	// 写入字符串
	st, err := json.Marshal(u)
	ioutil.WriteFile("database.txt", st, 0664)
	err = err

	// 确保写入到磁盘
	file.Sync()
}

func queryUser(s User) bool {
	file, _ := os.OpenFile("database.txt", os.O_RDWR|os.O_CREATE, 0664)
	defer file.Close()

	// ReadAll接收一个io.Reader的参数 返回字节切片
	bytes, _ := ioutil.ReadAll(file)
	// fmt.Println(string(bytes))
	var u Userlist
	json.Unmarshal(bytes, &u)

	for i, v := range u.List {
		i = i
		v = v
		if v.Name == s.Name {
			return true
		}
	}
	return false
}

func queryUserandPassword(s User) bool {
	file, _ := os.OpenFile("database.txt", os.O_RDWR|os.O_CREATE, 0664)
	defer file.Close()

	// ReadAll接收一个io.Reader的参数 返回字节切片
	bytes, _ := ioutil.ReadAll(file)
	// fmt.Println(string(bytes))
	var u Userlist
	json.Unmarshal(bytes, &u)

	for i, v := range u.List {
		i = i
		v = v
		if v.Name == s.Name && v.Password == s.Password {
			return true
		}
	}
	return false
}
