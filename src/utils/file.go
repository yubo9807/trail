package utils

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	Ext      string     `json:"ext"`
	Time     int64      `json:"time"`
	Size     int64      `json:"size"`
	IsDir    bool       `json:"isDir"`
	Body     string     `json:"body"`
	Children []FileInfo `json:"children"`
}

// 递归获取文件目录
func GetCatalog(filename string, isRecursion bool) []FileInfo {
	sli := []FileInfo{}

	filepath.Walk(filename, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		newPath := strings.Replace(path, filename, "", -1)
		if len(strings.Split(newPath, string(filepath.Separator))) == 2 {
			children := []FileInfo{}
			if isRecursion && info.IsDir() {
				children = append(children, GetCatalog(path, isRecursion)...)
			}
			sli = append(sli, FileInfo{
				Name:     info.Name(),
				Path:     path,
				Ext:      filepath.Ext(path),
				Size:     info.Size(),
				IsDir:    info.IsDir(),
				Time:     info.ModTime().Unix(),
				Children: children,
			})
		}

		return nil
	})

	return sli
}
