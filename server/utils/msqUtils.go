package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"strconv"

	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {

	fmt.Println("InitMysql....")
	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	//orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/copyright?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	//dbConn = "root:wxm19990516@tcp(127.0.0.1:3306)/copyright?charset=utf8"
	fmt.Println("dbConn:", dbConn)
	//driverName = "mysql"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
		CreateTableWithCopyright()
		CreateTableWithDocument()
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		publicKey VARCHAR(1024),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}
//创建版权表
func CreateTableWithCopyright() {
	sql := `create table if not exists copyright(
			id int(4) primary key auto_increment not null,
			name varchar(30),
			author varchar(20),
			press varchar(30),
			hash varchar(255),
			publicKey varchar(1024),
			signature varchar(1024),
			timestamp int(10),
			copyrightNum int(20)
			);`
	ModifyDB(sql)
}
//创建文件表
func CreateTableWithDocument() {
	sql := `create table if not exists document(
			id int(4) auto_increment not null,
			name varchar(30),
			path varchar(255),
			hash varchar(255),
			owner varchar(20),
			signature varchar(1024),
			timestamp int(10),
			primary key (id,hash)
			);`
	ModifyDB(sql)
}
//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

//传入的数据不一样，那么MD5后的32位长度的数据肯定会不一样
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

func SwitchTimeStampToData(timestamp int64) string {
	return strconv.FormatInt(timestamp, 10)
}

// 获取hash
func GetHash(path string) (hash string) {
	file, err := os.Open(path)
	defer file.Close()
	if err == nil {
		h_ob := sha256.New()
		_, err := io.Copy(h_ob, file)
		if err == nil {
			hash := h_ob.Sum(nil)
			hashvalue := hex.EncodeToString(hash)
			return hashvalue
		} else {
			return "something wrong when use sha256 interface..."
		}
	} else {
		fmt.Printf("failed to open %s\n", path)
	}
	return
}
