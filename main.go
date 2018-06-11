package main

import (
	_"blog/routers"
	"github.com/astaxie/beego"
	//"blog/models"
	"blog/models"
)

func main() {
	beego.SetStaticPath("/vue", "vue")
	beego.SetStaticPath("/dist", "views/dist")
	connStr := beego.AppConfig.String("mgoconnstr")
	print(connStr)
 	dbName := beego.AppConfig.String("dbname")
	print(dbName)

	session := models.InitDB(connStr, dbName)
	defer session.Close()
 	beego.Run()
}

