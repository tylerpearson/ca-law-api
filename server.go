package main

import (
    "database/sql"
    "github.com/go-martini/martini"
    _ "github.com/go-sql-driver/mysql"
    "github.com/martini-contrib/render"
    "github.com/coopernurse/gorp"
    "log"
)

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}

type Code struct {
    Code        string `db:"code" json:"code"`
    Title       string `db:"title" json:"title"`
}

var dbmap *gorp.DbMap

func initDb() *gorp.DbMap {

    db, err := sql.Open("mysql", "root:@/capublic")
    checkErr(err, "sql.Open failed")

    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

    return dbmap
}


func main() {
    m := martini.Classic()

    m.Use(render.Renderer())

    dbmap = initDb()

    defer dbmap.Db.Close()

    m.Get("/", func() string {
        return "hello world"
    })

    m.Get("/codes", func(r render.Render) {
        var codes []Code
        _, err := dbmap.Select(&codes, "SELECT * FROM codes_tbl")

        if err == nil {
            r.JSON(200, codes)
        } else {
            r.JSON(404, map[string]interface{}{"error": true})
        }
    })

    m.Run()
}

