package stream

import (
	"fmt"
	"server/src/service"
	"server/src/utils"

	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	url := "http://localhost:9100/trail/basic/api/stream/sse"
	res, err := utils.Request("GET", url, map[string]string{
		"Content-Type": "application/json",
	}, nil)
	if err != nil {
		service.State.ErrorCustom(ctx, err.Error())
		return
	}

	count := 0
	buffer := make([]byte, 1024)
	for {
		r, err := res.Body.Read(buffer)
		if err != nil {
			break
		}
		count++
		fmt.Println("-----------------------------\n", string(buffer[0:r]))
	}
	fmt.Println("end: ", count)

	service.State.SuccessData(ctx, "success")
}
