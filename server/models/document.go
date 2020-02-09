package models

import (
	"btcu-final/server/utils"
	"strconv"
)

type Document struct {
	Id        int
	Name      string
	Path      string
	Hash      string
	Owner     string
	Timestamp int64
	Signature string
}

func InsertDocument(document Document) (int64, error) {
	return utils.ModifyDB("insert into document(name, path, hash, owner, timestamp, signature)values(?,?,?,?,?,?)",
		document.Name, document.Path, document.Hash, document.Owner, document.Timestamp, document.Signature)
}

func FindDocument() ([]Document, error) {
	rows, err := utils.QueryDB("select name, path, hash, owner, timestamp, signature from document")
	if err != nil {
		return nil, err
	}
	var documents []Document
	for rows.Next() {
		id := 0
		name := ""
		path := ""
		hash := ""
		owner := ""
		signature := ""
		var timestamp int64 = 0
		rows.Scan(&name, &path, &hash, &owner, &timestamp, &signature)
		document := Document{id, name, path, hash, owner, timestamp, signature}
		documents = append(documents, document)
	}
	return documents, nil
}

// 通过id查找文件
func QueryDocumentWithId(id int) Document {
	row := utils.QueryRowDB("select id,name, path, hash,owner, timestamp, signature from document where id=" + strconv.Itoa(id))
	name := ""
	path := ""
	hash := ""
	owner := ""
	signature := ""
	var timestamp int64 = 0
	row.Scan(&name, &path, &hash, &owner, &timestamp, &signature)
	document := Document{id, name, path, hash, owner, timestamp, signature}
	return document
}

// 通过hash查找文件
func QueryDocumentWithHash(hash string) Document {
	sql := "select id,name, path, hash, owner, timestamp, signature from document where hash= " + `'` + hash + `'`
	row := utils.QueryRowDB(sql)
	id := 0
	name := ""
	path := ""
	owner := ""
	signature := ""
	var timestamp int64 = 0
	row.Scan(&id, &name, &path, &hash, &owner, &timestamp, &signature)
	document := Document{id, name, path, hash, owner, timestamp, signature}
	return document
}

// 用于修改文件
func UpdateDocument(document Document) (int64, error) {
	return utils.ModifyDB("update document set name=?,path=?,hash=?,owner=?,signature=? where id=?",
		document.Name, document.Path, document.Hash, document.Owner, document.Signature, document.Id)
}
