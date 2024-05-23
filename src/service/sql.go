package service

import (
	"database/sql"
	"fmt"
	"regexp"
	"server/configs"
	"server/src/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sqlType struct{}

var Sql sqlType

// 连接数据库
func (sql *sqlType) DBConnect() *LoggedDB {
	db, err := sqlx.Open("mysql", configs.Config.SqlSecret)
	if err != nil {
		panic(err.Error())
	}

	loggedDB := &LoggedDB{db}

	return loggedDB
}

type LoggedDB struct {
	*sqlx.DB
}

func (ldb *LoggedDB) Select(dest interface{}, query string, args ...interface{}) error {
	record(query, args...)
	return ldb.DB.Select(dest, query, args...)
}

func (ldb *LoggedDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	record(query, args...)
	return ldb.DB.Exec(query, args...)
}

// 格式化 sql
func formatSql(str string) string {
	regex1 := regexp.MustCompile(`\t`)
	str = regex1.ReplaceAllString(str, "")
	regex2 := regexp.MustCompile(`\n`)
	str = regex2.ReplaceAllString(str, " ")
	return str
}

var SqlStrs [][2]string

func record(query string, args ...interface{}) {
	sqlStr := formatSql(query)
	argsStr := ""
	lastIndex := len(args) - 1
	for i, val := range args {
		partition := utils.If(i == lastIndex, "", ", ")
		argsStr += fmt.Sprintf("%v", val) + partition
	}
	SqlStrs = append(SqlStrs, [2]string{sqlStr, argsStr})
}
