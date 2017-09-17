...
package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"           //用户名
	dbpassword = "dyk123"         //密码
	dbname     = "item_table"     //表名
)

type UserBean struct {
	userName string;
	userAge  int;
	userSex  int;
}

func main() {
	//bean := UserBean{"张三", 100, 1}
	//goInsert(bean)
	//goUpdata(bean)
	//remove(bean)
	goSelect()
}

func goSelect()  {
	db, err := sql.Open("mysql", "root:dyk123@tcp(127.0.0.1:3306)/Yikexin")
	checkErr(err);

	rows, err := db.Query("SELECT * FROM user_table")
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

/**
插入数据
 */
func goInsert(bean UserBean) {
	db, err := sql.Open("mysql", "root:dyk123@tcp(127.0.0.1:3306)/Yikexin")
	checkErr(err);
	stmt, err := db.Prepare("INSERT user_table SET user_name=?,user_age=?,user_sex=?")
	checkErr(err)

	res, err := stmt.Exec(bean.userName, bean.userAge, bean.userSex)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

/**
更新数据
 */
func goUpdata(bean UserBean) {

	db, err := sql.Open("mysql", "root:dyk123@tcp(127.0.0.1:3306)/Yikexin")
	checkErr(err)

	stmt, err := db.Prepare(`UPDATE user_table SET user_age=?,user_sex=? WHERE user_name=?`)
	checkErr(err)
	res, err := stmt.Exec(bean.userAge, bean.userSex, bean.userName)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

/**
删除
 */
func goRemove(bean UserBean) {
	db, err := sql.Open("mysql", "root:dyk123@tcp(127.0.0.1:3306)/Yikexin")
	checkErr(err)

	stmt, err := db.Prepare(`DELETE FROM user_table WHERE user_name=?`)
	checkErr(err)
	res, err := stmt.Exec(bean.userName)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func checkErr(i error) {
	if i != nil {
		panic(i)
	}
}

...
