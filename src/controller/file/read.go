package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"server/configs"
	"server/src/service"
	"server/src/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取文件目录
func Catalog(ctx *gin.Context) {
	type Params struct {
		Filename    string `form:"filename" binding:"required"`
		IsRecursion bool   `form:"isRecursion"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	baseUrl := configs.Config.AssetsUrl
	sli := utils.GetCatalog(baseUrl+params.Filename, params.IsRecursion)
	for i, v := range sli {
		sli[i].Path = strings.Replace(v.Path, baseUrl, "", 1)
	}
	service.State.SuccessData(ctx, sli)
}

// 读取文件内容
func ReadFile(ctx *gin.Context) {
	type Params struct {
		Filename string `form:"filename" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	file, err := os.Open(params.Filename)
	if err != nil {
		service.State.ErrorCustom(ctx, err.Error())
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		service.State.ErrorCustom(ctx, err.Error())
		return
	}

	info, _ := file.Stat()
	obj := utils.FileInfo{
		Name:  info.Name(),
		Path:  params.Filename,
		Ext:   filepath.Ext(params.Filename),
		Size:  info.Size(),
		IsDir: false,
		Time:  info.ModTime().Unix(),
		Body:  string(data),
	}

	service.State.SuccessData(ctx, obj)
}
