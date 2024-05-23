package spider

import (
	"server/src/service"
	"time"
)

type roomUserType struct{}

var RoomUser roomUserType

type RoomUserColumn struct {
	Id         string `db:"id"`
	RoomId     string `db:"room_id"`
	UserId     string `db:"user_id"`
	CreateTime string `db:"create_time"`
}

// 加入房间
func (roomUserType) JoinRoom(roomId, userId string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	createTime := time.Now().Unix()
	_, err := db.Exec("insert into room_user(room_id, user_id, create_time) values(?, ?, ?)", roomId, userId, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 退出房间
func (roomUserType) QuitRoom(roomId, userId string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	_, err := db.Exec("delete from room_user where room_id = ? and user_id = ?", roomId, userId)
	if err != nil {
		panic(err.Error())
	}
}

// 用户拥有的房间
func (roomUserType) GetUserRooms(userId string) []RoomColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	rooms := []RoomColumn{}
	keysOwn := ""
	for i, val := range Room.dbKeys {
		keysOwn += "t2." + val
		if i < len(Room.dbKeys)-1 {
			keysOwn += ", "
		}
	}
	db.Select(&rooms, "select "+keysOwn+" from room_user as t1 left join room as t2 on t1.room_id = t2.id where t2.user_id = ?", userId)
	return rooms
}

// 获取该房间的用户
func (roomUserType) GetRoomUsers(roomId string) []UserColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	users := []UserColumn{}
	db.Select(&users, "select t2.* from room_user as t1 left join user as t2 on t1.user_id = t2.id")
	return users
}
