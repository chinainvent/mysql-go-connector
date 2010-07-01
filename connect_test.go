package db

import (
    "testing"
)

type connTestInput struct {
    host, usr, pwd, db string
    port    uint 
}

type connTest struct {
    in      connTestInput
    out     int
} 

var connTests = []connTest {
    //correct parameter
    connTest{ connTestInput{"localhost", "webapi", "itbuwebapi", "webapi", 3306}, 0},
    connTest{ connTestInput{"127.0.0.1", "webapi", "itbuwebapi", "webapi", 3306}, 0},

    //incorrect parameters
    connTest{ connTestInput{"127.0.0.1", "webap", "itbuwebapi", "webapi", 3306}, 2},
    connTest{ connTestInput{"localhost", "webapi", "itbuwebap", "webapi", 3306}, 2},
    connTest{ connTestInput{"localhost", "webapi", "itbuwebapi", "webap", 3306}, 2},

    //empty parameters
    connTest{ connTestInput{"localhost", "webapi", "", "webapi", 3306}, 2},
    connTest{ connTestInput{"localhost", "", "", "webapi", 3306}, 2},
    connTest{ connTestInput{"", "", "", "webapi", 3306}, 2},
}

func TestConnect(t *testing.T) {
    var sql MySQL
    for _, d := range connTests {
        in := d.in
        v := sql.Connect(in.host, in.usr, in.pwd, in.db, in.port )
        if v!= d.out {
            t.Errorf("Connect(%v)=%v, want %v.", d.in, v, d.out)
        }
    }
}

func TestPing(t *testing.T) {
    var sql MySQL
    for _, d := range connTests {
        in := d.in
        v := sql.Connect(in.host, in.usr, in.pwd, in.db, in.port )
        if (sql.IsClosed() && v==0) || (!sql.IsClosed() && v!=0) {
            t.Errorf("IsCloseed(%v)=%v, want %v.", d.in, v, d.out)
        }
    }
}

func TestClose(t *testing.T) {
    var sql MySQL
    for _, d := range connTests {
        in := d.in
        sql.Connect(in.host, in.usr, in.pwd, in.db, in.port )
        sql.Close()
        if !sql.IsClosed() {
            t.Errorf("After Close, IsClosed=true, want false.")
        }
    }
}

