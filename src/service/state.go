package service

import (
	"bytes"
	"errors"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

type storeType struct {
	Code      int
	Data      any
	Message   string
	startTime time.Time
	RunTime   string
	Body      string
}

const key = "_state_store_"

func getStore(ctx *gin.Context) *storeType {
	data, ok := ctx.Get(key)
	if !ok {
		panic(errors.New("state init error"))
	}
	return data.(*storeType)
}
func setStore(ctx *gin.Context, store *storeType) {
	ctx.Set(key, store)
}

type stateType struct{}

var State stateType

// 初始化 store
func (s *stateType) InitState(ctx *gin.Context) {
	data, _ := ctx.GetRawData()                                // body 数据只能被读一次，读完即删
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 回写

	store := &storeType{
		Code:      400,
		Message:   "unknown error",
		startTime: time.Now(),
		Body:      string(data),
	}
	setStore(ctx, store)
}

// 获取状态数据
func (s *stateType) GetStateStore(ctx *gin.Context) *storeType {
	return getStore(ctx)
}

func (s *stateType) Clean(ctx *gin.Context) {
	store := getStore(ctx)
	store.Data = nil
}

// 返回统一格式
func (s *stateType) Result(ctx *gin.Context) {
	store := getStore(ctx)
	store.RunTime = time.Since(store.startTime).String()
	ctx.JSON(200, gin.H{
		"code":    store.Code,
		"data":    store.Data,
		"message": store.Message,
		"runTime": store.RunTime,
	})
	ctx.Abort()
}

// 请求成功，并返回数据
func (s *stateType) SuccessData(ctx *gin.Context, data interface{}) {
	store := getStore(ctx)
	store.Code = 200
	store.Message = "success"
	store.Data = data
}

// 请求成功
func (s *stateType) Success(ctx *gin.Context) {
	s.SuccessData(ctx, "success")
}

func errorParams(ctx *gin.Context, code int, msg string) {
	store := getStore(ctx)
	store.Code = code
	store.Message = msg
	ctx.Abort()
}

// 未授权
func (s *stateType) ErrorUnauthorized(ctx *gin.Context, msg string) {
	errorParams(ctx, 401, msg)
}

// token 失效，重新刷新 token
func (s *stateType) ErrorTokenFailure(ctx *gin.Context) {
	errorParams(ctx, 405, "token is expired")
}

// 参数错误
func (s *stateType) ErrorParams(ctx *gin.Context) {
	errorParams(ctx, 406, "params error")
}

// 请求超时
func (s *stateType) ErrorConnectTimeout(ctx *gin.Context) {
	errorParams(ctx, 504, "connect timeout")
}

// 自定义错误消息
func (s *stateType) ErrorCustom(ctx *gin.Context, msg string) {
	errorParams(ctx, 500, msg)
}
