package main

import (
	"database/sql"
	// "fmt"
	"log"
	"net/http"
	// "strconv"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    user_id int
    name string
    create_time time.Time
    update_time time.Time
}

type Post struct {
    post_id int
    title string
    body string
    author_id int
    create_time time.Time
    update_time time.Time
    
}
type Like struct {
    Like_id int
    Post_id int
    User_id int
    Create_time time.Time
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "pass"
    // dbName := "tcp(127.0.0.1:3306)/likebase"
    dbName := "likebase"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    
    return db
}

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT like_id, post_id, user_id FROM likes ORDER BY like_id ASC")
    if err != nil {
        panic(err.Error())
    }

    emp := Like{}
    res := []Like{}

    for selDB.Next() {
        var like_id, post_id, user_id int
        err = selDB.Scan(&like_id, &post_id, &user_id)
        if err != nil {
            panic(err.Error())
        }
        emp.Like_id = like_id
        emp.Post_id = post_id
        emp.User_id = user_id        
        res = append(res, emp)
    }
    var tmpl = template.Must(template.ParseFiles("form/Index.html"))
    tmpl.Execute(w, res)
    defer db.Close()
}

func countLike(post_id int, user_id int)(int) {
    db := dbConn()
    ins, err := db.Query("select count(like_id) as count from likes where post_id=? and user_id=?", post_id, user_id)
    if err != nil {
        panic(err.Error())
    }
    var count int
    for ins.Next() {
        ins.Scan(&count)
    }
    // fmt.Printf(strconv.Itoa(count))
    defer db.Close()
    return count
}

func insertLike(post_id int, user_id int) {
    if (countLike(post_id, user_id) != 0) {
        log.Println("Like Count Out of Range")
        return
    }
    db := dbConn()
    // var date = time.Now()
    ins, err := db.Prepare("insert likes set post_id=?, user_id=?, create_at=Now()")
    if err != nil {
        panic(err.Error())
    }
    ins.Exec(post_id, user_id)
    defer db.Close()
}

func cancelLike(post_id int, user_id int) {
    if (countLike(post_id, user_id) != 1) {
        log.Println("Like Count Out of Range")
        return
    }
    db := dbConn()
    ins, err := db.Prepare("delete from likes where post_id=? and user_id=?")
    if err != nil {
        panic(err.Error())
    }
    ins.Exec(post_id, user_id)
    defer db.Close()
}

func main() {
    http.HandleFunc("/", Index)
    // Click once
    insertLike(1, 1)
    // Click again
    cancelLike(1, 1)
    insertLike(1, 2)
    
    http.ListenAndServe(":8085", nil)

    
}