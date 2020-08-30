package main

import (
	log "github.com/cihub/seelog"
	"net/http"
	"tools/Mysql"
)

func Deafault(w http.ResponseWriter,r *http.Request)  {
	_, _ = w.Write([]byte("请输入正确的api"))
}

func Index(w http.ResponseWriter,r *http.Request)  {
	_, _ = w.Write([]byte("hello world"))
}

func GetUserInfo(w http.ResponseWriter,r *http.Request)  {
	conn := Mysql.Mysql{StrDb:"root:123456@tcp(127.0.0.1:3306)/jifen_db?charset=utf8"}

	conn.OpenMysql()//打开数据库

	resultList := conn.Select("select * from user_tb")
	var strRsp string
	for _,v := range resultList{
		strRsp += v
	}
	_,_ = w.Write([]byte(strRsp))
	conn.CloseMysql()
}

func init(){
	initLog()
	initRouter()
}


func initLog(){
	logger, err := log.LoggerFromConfigAsFile("C:\\Go\\gopath\\src\\webserver\\seelog.xml")

	if err != nil {
		log.Error("parse seelog.xml error")
	}

	log.ReplaceLogger(logger)

	defer log.Flush()
	log.Info("Hello from Seelog!")

}

func initRouter(){
	http.HandleFunc("/",Deafault)
	http.HandleFunc("/index",Index)
	http.HandleFunc("/query",GetUserInfo)
}

func main()  {
	defer log.Flush()

	err := http.ListenAndServe("127.0.0.1:9000",nil)
	if err != nil {
		log.Errorf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}

	return
}
