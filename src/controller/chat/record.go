package chat

import (
	"server/src/utils"
	"strconv"
	"time"
)

type Record struct {
	Id         string `json:"id"`
	CreateTime int64  `json:"createTime"`
	UserId     string `json:"userId"`
	Portrait   string `json:"portrait"`
	Body       string `json:"body"`
	RoomId     string `json:"roomId"`
}

var records = []Record{}

// 添加聊天记录
func recordAdd(roomId, userId, body string) {
	records = append(records, Record{
		Id:         strconv.FormatInt(utils.CreateID(), 10),
		CreateTime: time.Now().Unix(),
		Body:       body,
		RoomId:     roomId,
		UserId:     userId,
	})
}

// 获取当前房间聊天记录
func recordList(roomId string) []Record {
	return utils.Filter(records, func(val Record, i int) bool {
		return val.RoomId == roomId
	})
}

// 聊天详情
func recordDetail(recordId string) Record {
	return utils.Find(records, func(val Record, i int) bool {
		return val.Id == recordId
	})
}

// 删除聊天
func recordDel(recordId string) {
	records = utils.Filter(records, func(val Record, i int) bool {
		return val.Id == recordId
	})
}

// 清空房间聊天数据
func recordClear(roomId string) {
	records = utils.Filter(records, func(val Record, i int) bool {
		return val.RoomId == roomId
	})
}

// 获取最后一条聊天记录
func recordLast(roomId string) string {
	for i := len(records) - 1; i >= 0; i-- {
		if records[i].RoomId == roomId {
			return records[i].Body
		}
	}
	return ""
}
