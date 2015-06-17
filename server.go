package main

import (
    "database/sql"
    "github.com/go-martini/martini"
    _ "github.com/go-sql-driver/mysql"
)


func main() {
    m := martini.Classic()

    db, err := sql.Open("mysql", "root:@/capublic")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    m.Get("/", func() string {
        return "hello world"
    })

    m.Run()
}

