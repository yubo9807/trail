package stream

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func EventSource(ctx *gin.Context) {
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
