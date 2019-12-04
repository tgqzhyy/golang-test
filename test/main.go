package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

type Content struct {
	Name         string `json:"name"`                //文件名
	Size         int64  `json:"size"`                //大小
	Extension    string `json:"extension,omitempty"` //类型/扩展名
	IsDir        bool   `josn:"is_dir"`              //是否是目录
	LastModified string `json:"lastModified"`        //修改日期
	//Md5 string`json:"md5"`
}

// 获取文件夹列表
func getDirContent(p string) ([]Content, error) {
	files, _ := ioutil.ReadDir(p)

	var contents []Content

	for _, f := range files {
		content := Content{
			Name:         f.Name(),
			Size:         f.Size() / 1024,
			IsDir:        f.IsDir(),
			LastModified: f.ModTime().String(),
		}
		if !f.IsDir() {
			content.Extension = path.Ext(p + f.Name())
		}
		contents = append(contents, content)
	}
	return contents, nil
}
func main() {

	b, _ := getDirContent("/home/landv/")
	//aaa := b[1]
	//fmt.Println(aaa)
	for _, v := range b {
		fmt.Println(v.Name)
		fmt.Println(v.Size)
		fmt.Println(v.LastModified)
		fmt.Println()
	}
	//fmt.Println("b:",b)
}
