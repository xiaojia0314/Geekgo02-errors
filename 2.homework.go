package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const(
	UserName = "root"
	PassWord = "tiger"
)

func initSql()(db *sql.DB, err error){
	db, err = sql.Open("mysql", UserName+":"+PassWord+"@tcp(127.0.0.1:3306)/GeekGoError?charset=utf8")
	if err != nil{
		err = errors.Wrap(err, "Sql Open failed")
		return
	}
	err = db.Ping()
	if err != nil{
		err = errors.Wrap(err, "Sql link failed")
		return
	}
	return
}

type User struct {
	Id	 	int64
	Name  	string
}

func queryUser(db *sql.DB, sqlStr string)error{
	rows, err := db.Query(sqlStr)
	if err !=nil{
		return errors.Wrap(err, "query failed")
	}
	for rows.Next(){
		user := User{}
		if err = rows.Scan(&user.Id, &user.Name); err != nil{
			switch {
				case err == sql.ErrNoRows:
					err = errors.Wrap(err, "ErrNoRows")
				default:
					err = errors.Wrap(err, "Unknown Errors")
			}
			return err
		}
		fmt.Println(user)
	}
	fmt.Println("query finished")
	return err
}

func main() {
	sqlConn, err := initSql()
	defer sqlConn.Close()
	if err != nil{
		fmt.Println("InitSql err:", err)
		fmt.Printf("Detail error stackTrace message\n%+v\n", err)
		return
	}
	sqlStr := "select * from user where id <= 2"
	err = queryUser(sqlConn, sqlStr)
	if err != nil{
		fmt.Println("Query sql error:")
		fmt.Printf("Detail error stackTrace message\n%+v\n", err)
	}

}

/*
{1 tom}
Query sql error:
Detail error stackTrace message
sql: Scan error on column index 1, name "name": converting NULL to string is unsupported
Unknown Errors
main.queryUser
	E:/Golang Projects/src/Geekgo/2.errors/2.homework.go:46
main.main
	E:/Golang Projects/src/Geekgo/2.errors/2.homework.go:65
runtime.main
	E:/go/src/runtime/proc.go:225
runtime.goexit
	E:/go/src/runtime/asm_amd64.s:1371

*/