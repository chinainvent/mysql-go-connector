package db

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <mysql/mysql.h>

MYSQL_FIELD* next_field_name(MYSQL_FIELD* f){ return ++f; }
MYSQL_ROW    next_field_value(MYSQL_ROW r){ return ++r; }

*/
import "C"

import (
	"unsafe"
)


type MySQL struct {
    my     *C.MYSQL
    rs     *C.MYSQL_RES
    fi      []string
    errno   int
}

//建立一个到数据库服务器的连接, 
//成功返回0, 否则返回errnum
func (mysql *MySQL) Connect(host, usr, pwd, db string, port uint) (errnum int){
    mysql.my = C.mysql_init(nil)
    if mysql.my==nil {
        mysql.errno = 1
        return 1
    }

	h := C.CString(host)
	u := C.CString(usr)
	p := C.CString(pwd)
	d := C.CString(db)
	defer C.free(unsafe.Pointer(h))
	defer C.free(unsafe.Pointer(u))
	defer C.free(unsafe.Pointer(p))
	defer C.free(unsafe.Pointer(d))

	mysql.my = C.mysql_real_connect(mysql.my, h, u, p, d, (C.unsigned)(port), nil, 0)
    if mysql.my==nil {
        mysql.errno = 2
        return 2
    }
    mysql.errno = 0
    return 0
} 

//关闭一个连接
func (mysql MySQL) Close() {
	C.mysql_close(mysql.my)
    mysql.my = nil
    if mysql.rs!=nil {
        C.mysql_free_result(mysql.rs)
    }
    mysql.fi = nil
	return
}

//检查一个连接是否已关闭
func (mysql *MySQL) IsClosed() bool {
	f := C.mysql_ping(mysql.my)
    mysql.errno = int(f)
	if f == 0 {
        return true
	}
	return false
}

//执行SQL语句(select, insert, update, delete, ...), 
//成功返回0, 否则返回errnum
func (mysql *MySQL) Execute(stmt string) (errnum int) {
	s := C.CString(stmt)
	defer C.free(unsafe.Pointer(s))

	rc := C.mysql_real_query(mysql.my, s, (C.ulong)(len(stmt)))
    if rc!=0 {
        mysql.errno = 1
        return 1 
    }

    if mysql.rs!=nil {
        C.mysql_free_result(mysql.rs)
        mysql.fi = nil
    }
	mysql.rs = C.mysql_store_result(mysql.my)

    mysql.errno = 0
	return 0
    
}

//获取下一行数据(只用于select语句), 
//成功返回errnum=0且row有意义, 否则返回errnum
func (mysql *MySQL) NextRow() (row map[string]string, errnum int) {

    if mysql.rs==nil {
        if C.mysql_field_count(mysql.my)!=0 {
            return nil, 4
        }
        return nil, 5
    }

    if mysql.fi==nil {
        n := uint(C.mysql_num_fields(mysql.rs))
        mysql.fi = make([]string, n)

        f := C.mysql_fetch_fields(mysql.rs)

        for i, _ := range mysql.fi{
            mysql.fi[i] = C.GoString(f.name)
            f = C.next_field_name(f)
        }
    }
    
	r := C.mysql_fetch_row(mysql.rs)
	if r == nil {
        mysql.errno = 0
		return nil, 0
	}

	row = make(map[string]string)
	n := uint(len(mysql.fi))

	for i := uint(0); i < n; i++ {
		row[mysql.fi[i]] = C.GoString(*r)
		r = C.next_field_value(r)
	}

    mysql.errno = 0
    errnum = 0
    return

}

//获取影响的行数(插入的行数、更新的数行、删除的行数、查询的行数等), 
//成功返回0, 否则返回errnum
func (mysql *MySQL) AffectedRows() (rownum uint64) {
	rownum = uint64(C.mysql_affected_rows(mysql.my))
    return
}

//设置本连接的字符集, 成功返回true
func (mysql *MySQL) SetCharacterSet(cs string) bool {
    s := C.CString(cs)
    defer C.free(unsafe.Pointer(s))

    f := C.mysql_set_character_set(mysql.my, s)
    mysql.errno = 5
    if f == 0 {
        mysql.errno = 0
        return true
    }
    return false
}

//获取本连接的字符集
func (mysql *MySQL) GetCharacterSet() (cs string) {
    s := C.mysql_character_set_name(mysql.my)
    cs = C.GoString(s)
    return
}

//设置SQL语句执行完后，是否自动提交事务, 
//成功返回0, 否则返回errnum
func (mysql *MySQL) SetAutoCommit(mode bool) (errnum int) {
	f := C.my_bool(0)
	if mode {
		f = C.my_bool(1)
	}
	f = C.mysql_autocommit(mysql.my, f)
	mysql.errno = 5
	if f == C.my_bool(0) {
		mysql.errno = 0
	}

    return mysql.errno
}

//提交一个事务, 
//成功返回0, 否则返回errnum
func (mysql *MySQL) Commit() (errnum int) {
	f := C.mysql_commit(mysql.my)
	mysql.errno = 5
	if f == C.my_bool(0) {
        mysql.errno = 0
	}
	return mysql.errno
}

//返回一个errnum的描述串
func (mysql *MySQL)StrError() (desc string) {
    return errMsg[mysql.errno]
}
