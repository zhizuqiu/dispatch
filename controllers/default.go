package controllers

import (
	"dispatch/models"
	"dispatch/models/util"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type MainController struct {
	beego.Controller
}

/**
处理错误
*/
func (this *MainController) handleJson(jsonObj interface{}) {
	this.Data["json"] = jsonObj
	this.ServeJSON()
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) Upload() {
	c.TplName = "upload.html"
}

func (c *MainController) Upload2() {
	c.TplName = "upload2.html"
}

func (this *MainController) UploadFile() {
	path := this.GetString("Path")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	name := "upload"
	f, h, _ := this.GetFile(name)
	localPath := "data" + path + h.Filename
	defer f.Close()
	err := this.SaveToFile(name, localPath)
	if err != nil {
		this.handleJson(models.InternalServerError(err.Error()))
		return
	}
	this.handleJson(models.Success("success"))
}

type PathSlice []*Path

// 按名字排序
func (s PathSlice) Len() int      { return len(s) }
func (s PathSlice) Swap(i, j int) { *s[i], *s[j] = *s[j], *s[i] }
func (s PathSlice) Less(i, j int) bool {
	if (*s[i]).IsDir && !(*s[j]).IsDir {
		return true
	}
	if !(*s[i]).IsDir && (*s[j]).IsDir {
		return false
	}
	return (*s[i]).Name < (*s[j]).Name
}

type Path struct {
	Id           int
	IsDir        bool
	LastModified string
	Length       int64
	Name         string
	Path         string
}

func (this *MainController) List() {
	path := this.GetString("Path")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	fmt.Println("Path=" + path)

	localPath := "data" + path

	lists := make(PathSlice, 0)
	rd, err := ioutil.ReadDir(localPath)
	if err != nil {
		fmt.Println(err)
		this.handleJson(lists)
		return
	}

	index := 0
	for _, f := range rd {
		if !f.IsDir() {
			// 文件
			lists = append(lists, &Path{Id: index, IsDir: false, LastModified: f.ModTime().Format("2006-01-02 15:04:05"), Length: f.Size(), Name: f.Name(), Path: path + f.Name()})
			index++
		} else {
			// 目录
			lists = append(lists, &Path{Id: index, IsDir: true, LastModified: f.ModTime().Format("2006-01-02 15:04:05"), Length: f.Size(), Name: f.Name() + "/", Path: path + f.Name() + "/"})
			index++
		}
	}
	sort.Sort(lists)
	this.handleJson(lists)
}

func (this *MainController) Delete() {
	path := this.GetString("Path")
	isDir, err := this.GetBool("IsDir", false)
	if err != nil {
		fmt.Println(err)
		this.handleJson(models.InternalServerError(err.Error()))
		return
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !isDir {
		err = os.Remove("data" + path)
		if err != nil {
			fmt.Println(err)
			this.handleJson(models.InternalServerError(err.Error()))
			return
		}
	} else {
		fmt.Println(path)
		err = os.RemoveAll("data" + path)
		if err != nil {
			fmt.Println(err)
			this.handleJson(models.InternalServerError(err.Error()))
			return
		}
	}

	this.handleJson(models.Success(nil))
}
func (this *MainController) NewDir() {
	path := this.GetString("Path")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	fmt.Println(path)

	localPath := "data" + path

	oldMask, _ := util.Umask(0)
	err := os.MkdirAll(localPath, 0755)
	util.Umask(oldMask)
	if err != nil {
		fmt.Println(err)
		this.handleJson(models.InternalServerError(err.Error()))
		return
	}

	this.handleJson(models.Success(nil))
}
