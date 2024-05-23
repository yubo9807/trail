package spider

import (
	"server/src/service"
	"server/src/utils"
	"strings"
	"time"
)

type roomType struct {
	dbKeys  []string
	keysOwn string
}

var Room roomType

type RoomColumn struct {
	Id         string  `db:"id" json:"id"`
	Name       string  `db:"name" json:"name"`
	CreateTime string  `db:"create_time" json:"createTime"`
	UpdateTime *string `db:"update_time" json:"updateTime"`
	UserId     string  `db:"user_id" json:"userId"`
}

func init() {
	Room.dbKeys = utils.GetStructDBKeys(RoomColumn{})
	Room.keysOwn = strings.Join(Room.dbKeys, ", ")
}

// 房间列表
func (roomType) List(userId string) []RoomColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	roomList := []RoomColumn{}
	err := db.Select(&roomList, "select "+Room.keysOwn+" from room where user_id = ?", userId)
	if err != nil {
		panic(err.Error())
	}
	return roomList
}

// 房间详情
func (roomType) Detail(id string) *RoomColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	rooms := []RoomColumn{}
	err := db.Select(&rooms, "select "+Room.keysOwn+" from room where id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	if len(rooms) == 0 {
		return nil
	}
	return &rooms[0]
}

// 创建房间
func (roomType) Create(id int64, name, userId string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	createTime := time.Now().Unix()
	_, err := db.Exec("insert into room (id, name, user_id, create_time) values (?, ?, ?, ?)", id, name, userId, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 修改房间信息
func (roomType) Update(id, name string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec("update room set name = ?, update_time = ? where id = ?", name, updateTime, id)
	if err != nil {
		panic(err.Error())
	}
}

// 删除房间
func (roomType) Delete(id string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	_, err1 := db.Exec("delete from room where id = ?", id)
	if err1 != nil {
		panic(err1)
	}
	_, err2 := db.Exec("delete from room_user where room_id = ?", id)
	if err2 != nil {
		panic(err2)
	}
}
