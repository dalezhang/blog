package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dalezhang/blog/utils"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//单文件上传
//<input id="imgFile" name="imgFile" type="file" />
type UploadController struct {
	BaseController
}

func (this *UploadController) Post() {
	if !this.isLogin {
		this.Data["json"] = map[string]interface{}{"error": 1, "message": "你没有权限上传"}
		this.ServeJSON()
		return
	}
	//imgFile
	f, h, err := this.GetFile("imgFile")
	defer f.Close()

	//生成上传路径
	now := time.Now()
	dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
	err1 := os.MkdirAll(dir, 0755)
	if err1 != nil {
		this.Data["json"] = map[string]interface{}{"error": 1, "message": "目录权限不够"}
		this.ServeJSON()
		return
	}
	//生成新的文件名
	filename := h.Filename
	ext := utils.SubString(filename, strings.LastIndex(filename, "."), 5)
	filename = utils.GetGuid() + ext

	if err != nil {
		this.Data["json"] = map[string]interface{}{"error": 1, "message": err}
	} else {
		this.SaveToFile("imgFile", dir+"/"+h.Filename)
		this.Data["json"] = map[string]interface{}{"error": 0, "url": strings.Replace(dir, ".", "", 1) + "/" + filename}
	}
	this.ServeJSON()
}

//多文件上传
//<input id="albumUpload" name="uploadFiles" type="file" multiple class="file-loading" data-allowed-file-extensions='["jpg", "jpeg", "png", "gif"]'>
type UploadMultiController struct {
	BaseController
}

func (this *UploadMultiController) Post() {
	if !this.isLogin {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "你没有权限上传"}
		this.ServeJSON()
		return
	}

	files, err := this.GetFiles("uploadFiles")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "目录权限不够"}
		this.ServeJSON()
		return
	}

	//生成上传路径
	now := time.Now()
	dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
	err1 := os.MkdirAll(dir, 0755)
	if err1 != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "目录权限不够"}
		this.ServeJSON()
		return
	}

	resfilestr := ""
	resfilename := ""
	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
			this.ServeJSON()
			return
		}

		//生成新的文件名
		filename := files[i].Filename
		resfilename += utils.GetFileSuffix(filename) + "||"

		ext := utils.SubString(filename, strings.LastIndex(filename, "."), 5)
		filename = utils.GetGuid() + ext
		dst, err := os.Create(dir + "/" + filename)
		defer dst.Close()
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
			this.ServeJSON()
			return
		}
		if _, err := io.Copy(dst, file); err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
			this.ServeJSON()
			return
		}
		resfilestr += strings.Replace(dir, ".", "", 1) + "/" + filename + "||"
	}
	this.SetSession("uploadMultiPic", resfilestr)
	this.SetSession("uploadMultiName", resfilename)

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功", "url": resfilestr}
	this.ServeJSON()
	return
}
