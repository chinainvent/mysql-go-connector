package db 

var errMsg = [...]string { 
    0:"success", 
    1:"inner error", 
    2:"connect failed", 
    3:"connection was closed", 
    4:"not select stmt", 
    5:"other error",
}

type SQL interface {
    //建立一个到数据库服务器的连接, 
    //成功返回0, 否则返回errnum
    Connect(host, usr, pwd, db string, port uint) (errnum int) 
    //关闭一个连接
    Close()
    //检查一个连接是否已关闭
    IsClosed() bool
    //设置本连接的字符集, 成功返回true
    SetCharacterSet(cs string) bool
    //获取本连接的字符集
    GetCharacterSet() (cs string) 
    //执行SQL语句(select, insert, update, delete, ...), 
    //成功返回0, 否则返回errnum
    Execute(stmt string) (errnum int)
    //获取下一行数据(只用于select语句), 
    //成功返回errnum=0且row有意义, 否则返回errnum
    NextRow() (row map[string]string, errnum int)
    //获取影响的行数(插入的行数、更新的数行、删除的行数、查询的行数等), 
    //成功返回0, 否则返回errnum
    AffectedRows() (rownum uint64)
    //设置SQL语句执行完后，是否自动提交事务, 
    //成功返回0, 否则返回errnum
    SetAutoCommit(mode bool) (errnum int)
    //提交一个事务, 
    //成功返回0, 否则返回errnum
    Commit() (errnum int)
    //返回当前的错误描述串
    StrError() (desc string)
}

