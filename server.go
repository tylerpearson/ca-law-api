package main

import (
    "database/sql"
    "github.com/coopernurse/gorp"
    "github.com/go-martini/martini"
    _ "github.com/go-sql-driver/mysql"
    "github.com/martini-contrib/render"
    "log"
)

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}

type Code struct {
    Code  string `db:"code" json:"code"`
    Title string `db:"title" json:"title"`
}

type Toc struct {
    LawCode             string  `db:"law_code" json:"law_code"`
    Division            *string `db:"division" json:"division"`
    Title               *string `db:"title" json:"title"`
    Part                *string `db:"part" json:"part"`
    Chapter             *string `db:"chapter" json:"chapter"`
    Article             *string `db:"article" json:"article"`
    Heading             *string `db:"heading" json:"heading"`
    Active              *string `db:"active_flg" json:"active"`
    TransUID            *string `db:"trans_uid" json:"trans_uid"`
    TransUpdate         *string `db:"trans_update" json:"trans_update"`
    NodeSequence        *string `db:"node_sequence" json:"node_sequence"`
    NodeLevel           *string `db:"node_level" json:"node_level"`
    NodePosition        *string `db:"node_position" json:"node_position"`
    NodeTreePath        *string `db:"node_treepath" json:"node_treepath"`
    ContainsLawSections *string `db:"contains_law_sections" json:"contains_law_sections"`
    HistoryNote         *string `db:"history_note" json:"history_note"`
    OpStatues           *string `db:"op_statues" json:"op_statues"`
    OpChapter           *string `db:"op_chapter" json:"op_chapter"`
    OpSection           *string `db:"op_section" json:"op_section"`
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
        return "Welcome to the California Laws API!"
    })

    m.Group("/api", func(r martini.Router) {

        m.Get("/codes", func(r render.Render) {
            var codes []Code
            _, err := dbmap.Select(&codes, "SELECT * FROM codes_tbl")

            if err == nil {
                r.JSON(200, codes)
            } else {
                r.JSON(404, map[string]interface{}{"error": true})
            }
        })

        m.Get("/tocs", func(r render.Render) {
            var tocs []Toc

            _, err := dbmap.Select(&tocs, "SELECT law_code, division, title, part, chapter, article, heading, active_flg, trans_uid, trans_update, node_sequence, node_level, node_position, node_treepath, contains_law_sections, history_note, op_chapter, op_section FROM law_toc_tbl LIMIT 100")

            if err == nil {
                r.JSON(200, tocs)
            } else {
                log.Fatal(err)
                r.JSON(404, map[string]interface{}{"error": true})
            }
        })

    })

    m.Run()
}
