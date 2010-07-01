package db

import (
    "testing"
    "log"
)

type queryTest struct {
    in  string
    out string
}

var queryTests = []queryTest{
    queryTest{"select a, b from t", "a,b|c,d|"},
}

func TestQuery(t *testing.T) {
    log.Stdoutf("TestQuery")
    var sql MySQL
    for _, d := range queryTests {
        v := sql.Connect("localhost", "webapi", "itbuwebapi", "webapi", 3306)
        if v!= 0 {
            t.Errorf("Connect error.")
        }

        v = sql.Execute(d.in)
        if v!=0 {
            t.Errorf("Execute(%v)=%v, want 0.", d.in, v)
        }

        var result string 
        for r,_:=sql.NextRow(); r!=nil; r,_=sql.NextRow() {
            result += r["a"]+","+r["b"]+"|"
        }
        if result != d.out {
            t.Errorf("Query(%v)=%v, want %v.", d.in, result, d.out )
        }
        
    }
}

