package main

import (
	"database/sql"
	DbModel "gowork/model/db"
	Config "gowork/utils"
	"log"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	_ "github.com/lib/pq"
)

//Home o
func Home(app *iris.Application, engine *xorm.Engine) {

	app.Get("/sql1", func(ctx iris.Context) {
		var (
			id       int
			username string
		)

		config := Config.ReadConfig("config.ini")
		connStr := config["connStr"]
		db, err := sql.Open("postgres", connStr)
		log.Print("conted")
		if err != nil {
			log.Fatal(err)
			ctx.Text("err1" + err.Error())
			return
		}
		defer db.Close()
		rows, err := db.Query("select id,username from member;")
		if rows == nil {
			log.Fatal(err)
			ctx.Text("err2" + err.Error())
			return
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
			ctx.Text("t1err:" + err.Error())
			return
		}
		defer rows.Close()
		strs := ""
		for rows.Next() {
			err := rows.Scan(&id, &username)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, username)
			strs += username
		}
		ctx.Text("t1:" + strs)
	})

	app.Get("/list", func(ctx iris.Context) {
		//xorm reverse postgres "user=lfo dbname=lfo password=lfo sslmode=disable host=lfoteam.ddns.net" $GOPATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm ./db
		mlist := make([]DbModel.Member, 0)
		err := engine.Find(&mlist)
		if err != nil {
			log.Fatal(err)
			ctx.Text("err:" + err.Error())
			return
		}
		// strs := "xorm"
		// for _, entity := range mlist {
		// 	strs += entity.Username
		// 	log.Print(strs)
		// }
		ctx.JSON(mlist)
	})
	app.Post("/add", func(ctx iris.Context) {
		var member DbModel.Member
		ctx.ReadJSON(&member)
		member.Time = time.Now()
		affected, err := engine.Insert(&member)
		log.Print(affected)
		if err != nil {
			log.Fatal(err)
			res := BaseResponse{}
			res.Sus = false
			ctx.JSON(res)
			return
		}
		res := BaseResponse{}
		res.Sus = true
		ctx.JSON(res)
	})

}
