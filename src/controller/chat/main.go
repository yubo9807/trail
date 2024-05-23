package chat

import (
	"encoding/json"
	"fmt"
	"server/src/service"
	"server/src/spider"
	"server/src/utils"
	"strconv"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type messageType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type userInfoType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type userColumnType struct {
	spider.UserColumn
	Online bool `json:"online"`
}

func Chat() *gosocketio.Server {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	// 授权
	server.On("auth", func(c *gosocketio.Channel, token string) {
		if msg := recover(); msg != nil {
			c.Emit("error", messageType{500, fmt.Sprintf("%v", msg)})
		}

		tokenInfo, err := service.Jwt.Verify(token)
		if err != nil {
			c.Emit("error", messageType{401, err.Error()})
			return
		}
		userInfo := userInfoType{
			Name: fmt.Sprintf("%v", tokenInfo["username"]),
			Id:   fmt.Sprintf("%v", tokenInfo["userId"]),
		}

		// 获取房间的聊天记录
		server.On("record_list", func(c *gosocketio.Channel, roomId string) {
			records := recordList(roomId)
			c.Emit("records_"+roomId, records)
		})

		// 添加聊天数据
		server.On("record_add", func(c *gosocketio.Channel, data string) {
			obj := map[string]string{}
			json.Unmarshal([]byte(data), &obj)
			recordAdd(obj["roomId"], userInfo.Id, obj["body"])
			server.BroadcastToAll("records_add_"+userInfo.Id, obj) // 推送给在线所有用户
		})

		// 删除聊天数据
		server.On("record_del", func(c *gosocketio.Channel, roomId, recordId string) {
			recordInfo := recordDetail(recordId)
			if recordInfo.UserId != userInfo.Id {
				c.Emit("error", messageType{400, "forbidden"})
				return
			}
			recordDel(recordId)
			records := recordList(roomId)
			server.BroadcastToAll("records_"+roomId, records)
		})

		// 清空房间聊天数据
		server.On("record_clear", func(c *gosocketio.Channel, roomId string) {
			// 权限校验
			recordClear(roomId)
			records := recordList(roomId)
			server.BroadcastToAll("records_"+roomId, records)
		})

		// 创建房间
		server.On("room_create", func(c *gosocketio.Channel, roomName string) {
			id := utils.CreateID()
			spider.Room.Create(id, roomName, userInfo.Id)
			roomInfo := spider.Room.Detail(strconv.FormatInt(id, 10))
			c.Emit("room_create_info", roomInfo)
		})

		// 修改房间信息
		server.On("room_edit", func(c *gosocketio.Channel, data struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}) {
			spider.Room.Update(data.Id, data.Name)
		})

		// 退出/删除房间
		server.On("room_quit", func(c *gosocketio.Channel, roomId string) {
			room := spider.Room.Detail(roomId)
			if room != nil && room.UserId == userInfo.Id {
				spider.Room.Delete(roomId)
			} else {
				spider.RoomUser.QuitRoom(roomId, userInfo.Id)
			}
		})

		// 获取房间详情
		server.On("room_detail", func(c *gosocketio.Channel, roomId string) {
			roomInfo := spider.Room.Detail(roomId)
			if roomInfo == nil {
				c.Emit("error", messageType{500, "room not found"})
			}
			c.Emit("room_"+roomId, roomInfo)
		})

		// 加入房间
		server.On("room_join", func(c *gosocketio.Channel, roomId string) {
			spider.RoomUser.JoinRoom(roomId, userInfo.Id)
		})

		// 踢出房间
		server.On("room_kick", func(c *gosocketio.Channel, data struct {
			RoomId  string   `json:"roomId"`
			UserIds []string `json:"userIds"`
		}) {
			room := spider.Room.Detail(data.RoomId)
			if room.UserId != userInfo.Id {
				c.Emit("error", messageType{403, "forbidden"})
				return
			}
			for _, val := range data.UserIds {
				spider.RoomUser.QuitRoom(data.RoomId, val)
			}
		})

		// 所有用户
		onlineUsers := []userColumnType{}
		users := spider.User.List(1, 10, "")
		for _, user := range users {
			user := userColumnType{
				UserColumn: user,
				Online:     false,
			}
			onlineUsers = append(onlineUsers, user)
		}

		server.BroadcastToAll("online", userInfo) // 上线通知
		for i, val := range onlineUsers {
			if val.Id == userInfo.Id {
				onlineUsers[i].Online = true
			}
		}
		server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) userInfoType {
			server.BroadcastToAll("offline", userInfo) // 离线通知
			for i, val := range onlineUsers {
				if val.Id == userInfo.Id {
					onlineUsers[i].Online = false
				}
			}
			return userInfo
		})

		c.Emit("users_online", onlineUsers) // 在线用户

		// 用户拥有的房间
		type RoomColumn2 struct {
			spider.RoomColumn
			LastMessage string `json:"lastMessage"`
		}
		rooms := spider.Room.List(userInfo.Id)
		newRooms := []RoomColumn2{}
		for _, val := range rooms {
			newRooms = append(newRooms, RoomColumn2{
				RoomColumn:  val,
				LastMessage: recordLast(val.Id),
			})
		}
		c.Emit("rooms_"+userInfo.Name, newRooms) // 用户拥有的房间

	})

	return server
}
