package test

import (
	"bytes"
	"fmt"
	"io"
	"server/src/service"
	"server/src/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-store")

	sli := []string{"hello", ", ", "world", "!"}
	count := 0
	ctx.Stream(func(w io.Writer) bool {
		time.Sleep(time.Millisecond * 10)
		str := sli[count]

		w.Write([]byte("data:{\"data\":\"" + str + "\"}\n\n"))

		if count == len(sli)-1 {
			return false
		} else {
			count++
			return true
		}
	})
}

func Test2(ctx *gin.Context) {
	url := "http://10.0.7.15:8080/wentuApi/selectExtensionVirtual"
	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJsb2dpblR5cGUiOiJsb2dpbiIsImxvZ2luSWQiOiJzeXNfdXNlcjoxIiwicm5TdHIiOiI2ekFMRlBwYXdTeG5iSUFaeklsZURqUG1IZ0phTXVYeCIsInVzZXJJZCI6MX0.af3nLZxQdZZQoUwEf6kikECUlL5_oOs1CbjZMZv6jHY"
	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
	}
	body := []byte(``)
	res, err := utils.Request("POST", url, header, bytes.NewBuffer(body))
	if err != nil {
		service.State.ErrorCustom(ctx, err.Error())
		return
	}

	buffer := make([]byte, 1024*1024*10)
	for {
		r, err := res.Body.Read(buffer)
		if err != nil {
			break
		}
		if r == '\n' {
			break
		}
		fmt.Println("--------------------------\n", string(buffer))
	}
	fmt.Println("end")

	service.State.SuccessData(ctx, "success")
}

func Test3(ctx *gin.Context) {
	url := "http://localhost:9100/trail/basic/api/test"
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
