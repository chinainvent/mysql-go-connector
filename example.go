package main

import (
    "fmt"
    "db"
)


func Test(sql db.SQL){

    errnum := sql.Connect("localhost", "webapi", "itbuwebapi", "webapi", 3306)
    if errnum!=0 {
        fmt.Println(sql.StrError())
        return
    }

    errnum = sql.Execute("select * from call_log")
    if errnum!=0 {
        fmt.Println(sql.StrError())
        return
    }


    for r,_:=sql.NextRow(); r!=nil; r,_=sql.NextRow() {
        for _,v:= range r {
            fmt.Printf(" %s |", v)
        }
        fmt.Println()
    }
    fmt.Println()
    
}


func main() {

    var sql db.MySQL
    Test(&sql)
    (db.SQL)(&sql).Close()

    var sq db.SQL = new(db.MySQL)
    sq.Close()

    return
}

