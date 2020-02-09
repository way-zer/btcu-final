package models

import (
	"btcu-final/server/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type Copyright struct {
	Id           int
	Name         string
	Author       string
	Press        string
	Hash         string
	PublicKey    string
	Signature    string
	Timestamp    int64
	CopyrightNum int
}

// 添加版权信息
func AddCopyright(copyright Copyright) (int64, error) {
	i, err := insertCopyright(copyright)
	return i, err
}

func insertCopyright(copyright Copyright) (int64, error) {
	return utils.ModifyDB("insert into copyright(name, author, press, hash, publicKey, signature, timestamp, copyrightNum) values(?,?,?,?,?,?,?,?)",
		copyright.Name, copyright.Author, copyright.Press, copyright.Hash, copyright.PublicKey, copyright.Signature, copyright.Timestamp, copyright.CopyrightNum)
}

// 条件查询
func QueryCopyrightWithCon(sql string) ([]Copyright, error) {
	sql = "select id,name,author,press,hash,publicKey,signature,timestamp,copyrightNum from copyright " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var copyrightList []Copyright
	for rows.Next() {
		id := 0
		name := ""
		author := ""
		press := ""
		hash := ""
		publicKey := ""
		signature := ""
		copyrightNum := 0
		var timestamp int64
		timestamp = 0
		rows.Scan(&id, &name, &author, &press, &hash, &publicKey, &signature, &timestamp, &copyrightNum)
		copyright := Copyright{id, name, author, press, hash, publicKey, signature, timestamp, copyrightNum}
		copyrightList = append(copyrightList, copyright)
	}
	return copyrightList, nil
}

// 按照id查询
func QueryCopyrightWithId(id int) Copyright {
	row := utils.QueryRowDB("select id,name,author,press,hash,publicKey,signature,timestamp,copyrightNum from copyright " + strconv.Itoa(id))
	name := ""
	author := ""
	press := ""
	hash := ""
	publicKey := ""
	signature := ""
	copyrightNum := 0
	var timestamp int64
	timestamp = 0
	row.Scan(&id, &name, &author, &press, &hash, &publicKey, &signature, &timestamp, &copyrightNum)
	copyright := Copyright{id, name, author, press, hash, publicKey, signature, timestamp, copyrightNum}
	return copyright
}

// 通过hash查找文件
func QueryCopyrightWithHash(hash string) Copyright {
	sql := "select id,name,author,press,hash,publicKey,signature,timestamp,copyrightNum from copyright where hash= " + `'` + hash + `'`
	row := utils.QueryRowDB(sql)
	id := 0
	name := ""
	author := ""
	press := ""
	publicKey := ""
	signature := ""
	copyrightNum := 0
	var timestamp int64
	timestamp = 0
	row.Scan(&id, &name, &author, &press, &hash, &publicKey, &signature, &timestamp, &copyrightNum)
	copyright := Copyright{id, name, author, press, hash, publicKey, signature, timestamp, copyrightNum}
	return copyright
}

func QueryCopyrightWithName(name, author string) Copyright {
	sql := fmt.Sprintf("where name='%s' and author='%s' ", name, author)
	fmt.Println(sql)
	copyright, err := QueryCopyrightWithCon(sql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(copyright[0])
	return copyright[0]
}

// 更新版权信息
func UpdateCopyright(copyright Copyright) (int64, error) {
	return utils.ModifyDB("update articles set name=?,author=?,press=?,hash=?,publicKey=?,signature=? where id=?",
		copyright.Name, copyright.Author, copyright.Press, copyright.Hash, copyright.PublicKey, copyright.Signature)
}

func FindCopyrightWithPage(page int) ([]Copyright, error) {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryCopyrightWithPage(page, num)
}

func QueryCopyrightWithPage(page, num int) ([]Copyright, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryCopyrightWithCon(sql)
}

var copyrightRowsNum = 0

func GetCopyrightRowsNum() int {
	if copyrightRowsNum == 0 {
		copyrightRowsNum = QueryCopyrightRowNum()
	}
	return copyrightRowsNum
}

func QueryCopyrightRowNum() int {
	row := utils.QueryRowDB("select count(id) from copyright")
	num := 0
	row.Scan(&num)
	return num
}

func SetCopyrightRowsNum() {
	copyrightRowsNum = QueryCopyrightRowNum()
}
