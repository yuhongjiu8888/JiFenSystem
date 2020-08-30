package Mysql

import (
	"database/sql"
	"encoding/json"

	//"fmt"
	log "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"webserver/data"
)

type Mysql struct {
	DbHander *sql.DB
	StrDb string
}

func (m *Mysql)OpenMysql()  {
	//操作数据库代码
	//第一个参数是数据库驱动
	//第二个参数是链接数据库字符串
	//u:u@tcp(127.0.0.1:3306)
	var err error
	m.DbHander, err = sql.Open("mysql", m.StrDb)
	if err != nil {
		log.Errorf("链接错误", err)
		return
	}

	log.Debug("mysql connect success!")

}

func (m *Mysql)CloseMysql() {
	m.DbHander.Close()
	log.Debug("mysql is closeing!")
}

func (m *Mysql)Select(query string)([]string){
	//查询多行数据
	var person data.User
	resList := []string{}

	res, err := m.DbHander.Query(query)
	if err != nil{
		log.Errorf("Query error,%v",err)
		return resList
	}
	log.Debug(res)

	for res.Next() {
		res.Scan(&person.Id,&person.UserNo, &person.UserName, &person.UserPasswd,&person.UserCreatetime)
		log.Info(person)
		b,_:=json.Marshal(person)
		str:=string(b)
		resList = append(resList,str)
	}
	return resList
}