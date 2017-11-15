package initial

import (
	"github.com/dalezhang/blog/utils"
	"github.com/astaxie/beego"
)

func InitTplFunc() {
	beego.AddFuncMap("date_mh", utils.GetDateMH)
	beego.AddFuncMap("date", utils.GetDate)
	beego.AddFuncMap("avatar", utils.GetGravatar)
}