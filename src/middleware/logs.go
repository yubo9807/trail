package middleware

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"server/configs"
	"server/src/service"
	"server/src/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 重写 Write([]byte) (int, error)
func (w responseWriter) Write(body []byte) (int, error) {
	w.body.Write(body)                  // 向一个 bytes.buffer 中再写一份数据
	return w.ResponseWriter.Write(body) // 完成 gin.Context.Writer.Write() 原有功能
}

func init() {
	err := os.Mkdir("logs", 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Logs(ctx *gin.Context) {

	writer := responseWriter{
		ctx.Writer,
		bytes.NewBuffer([]byte{}),
	}
	ctx.Writer = writer

	ctx.Next()

	response := ""
	if service.State.GetStateStore(ctx).Code != 200 {
		response = "\nresponse: " + writer.body.String()
	}
	LogsWrite(ctx, response)
}

// 写入日志
func LogsWrite(ctx *gin.Context, append string) {
	filename := utils.DateFormater(time.Now(), "YYYY-MM-DD")

	logSrc := LogsGetSrc("logs/" + filename + ".log")
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	log.SetOutput(logSrc)
	log.SetPrefix("\n")

	state := service.State.GetStateStore(ctx)

	reg := regexp.MustCompile(" |\n")
	body := reg.ReplaceAllString(state.Body, "")

	log.Println(
		state.RunTime,
		ctx.ClientIP(),
		ctx.Request.Method,
		ctx.Request.RequestURI,
		utils.If(body == "", "", "\nbody:"+body),
		append,
	)

}

// 获取日志文件路径
func LogsGetSrc(filename string) *os.File {
	src, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	logSrc := src
	if err != nil {
		src, _ := os.Create(filename)
		logSrc = src
		go clearLogs()
	}
	return logSrc
}

// 清理日志
func clearLogs() {
	flag := time.Now().AddDate(0, 0, -configs.Config.LogReserveTime).Unix()
	filepath.Walk("logs", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".log") {
			return nil
		}

		name := strings.Split(info.Name(), ".")[0]
		t, _ := time.Parse("2006-01-02", name)
		if t.Unix() < flag {
			os.Remove(path)
		}
		return nil
	})
}
