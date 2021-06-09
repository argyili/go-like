package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/test")
    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()
    // make sure connection is available
    err = db.Ping()
    if err != nil {
        fmt.Print(err.Error())
    }

    stmt, err := db.Prepare("CREATE TABLE user (id int NOT NULL AUTO_INCREMENT, name int(11), age varchar(40), PRIMARY KEY (id));")
    if err != nil {
        fmt.Println(err.Error())
    }
    _, err = stmt.Exec()
    if err != nil {
        fmt.Print(err.Error())
    } else {
        fmt.Printf("User Table successfully ....")
    }
}