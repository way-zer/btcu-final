package models

import (
	"btcu-final/server/utils"
	"fmt"
)

type User struct {
	Id         int
	Username   string
	Password   string
	PublicKey  string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

//--------------数据库操作-----------------

//插入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,publicKey,status,createtime) values (?,?,?,?,?)",
		user.Username, user.Password, user.PublicKey, user.Status, user.Createtime)
}

//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}

func QueryUserWithPublicKey(publicKey string) int {
	sql := fmt.Sprintf("where publicKey='%s'", publicKey)
	return QueryUserWightCon(sql)
}

func GetPublicKeyWithUsername(username string) (publickey string) {
	sql := fmt.Sprintf("select publicKey from users where username='%s'", username)
	row := utils.QueryRowDB(sql)
	row.Scan(&publickey)
	return publickey
}
