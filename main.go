package main

import (
	Config "gowork/utils"
	"log"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	config := Config.ReadConfig("config.ini")
	connStr := config["connStr"]
	var engine *xorm.Engine
	var err error
	log.Print("connStr:" + connStr)
	engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	engine.ShowSQL(true)
	Home(app, engine)
	app.Run(iris.Addr(":9001"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}
