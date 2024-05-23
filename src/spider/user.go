package spider

import (
	"server/src/service"
	"server/src/utils"
	"strconv"
	"strings"
	"time"
)

type userType struct {
	dbKeys  []string
	keysOwn string
}

var User userType

type UserColumn struct {
	Id         string  `db:"id" json:"id"`
	Name       string  `db:"name" json:"name"`
	Role_id    string  `db:"role_id" json:"roleId"`
	CreateTime string  `db:"create_time" json:"createTime"`
	UpdateTime *string `db:"update_time" json:"updateTime"`
}

type UserColumnAll struct {
	UserColumn
	Pass string `db:"pass" json:"pass"`
}

func init() {
	User.dbKeys = utils.GetStructDBKeys(UserColumn{})
	User.keysOwn = strings.Join(User.dbKeys, ", ")
}

// 用户列表
func (userType) List(pageNumber, pageSize int, name string) []UserColumn {
	db := service.Sql.DBConnect()
	defer db.Close()

	// 查询条件
	query := ""
	if name != "" {
		query = " WHERE name LIKE '%" + name + "%'"
	}

	// 分页
	paging := ""
	if pageNumber > 0 && pageSize > 0 {
		paging = " LIMIT " + strconv.Itoa(pageNumber-1*pageSize) + ", " + strconv.Itoa(pageSize)
	}

	list := []UserColumn{}
	err := db.Select(&list, "SELECT "+User.keysOwn+" FROM user"+query+paging)
	if err != nil {
		panic(err.Error())
	}
	return list
}

// 用户详情 id|name
func (userType) Detail(key string) *UserColumnAll {
	db := service.Sql.DBConnect()
	defer db.Close()
	list := []UserColumnAll{}
	err := db.Select(&list, "SELECT "+User.keysOwn+", pass FROM user WHERE id = ? OR name = ?", key, key)
	if err != nil {
		panic(err.Error())
	}
	if len(list) > 0 {
		return &list[0]
	}
	return nil
}

// 添加用户
func (userType) Add(name, pass, roleId string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO user(id, name, pass, role_id, create_time) VALUES(?, ?, ?, ?, ?)", id, name, pass, roleId, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 修改用户
func (userType) Update(id, name, pass, roleId string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec("UPDATE user SET name = ?, pass = ?, roleId = ?, update_time = ? WHERE id = ?", name, pass, updateTime, roleId, id)
	if err != nil {
		panic(err.Error())
	}
}

// 删除用户
func (userType) Delete(id string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	_, err := db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
}
