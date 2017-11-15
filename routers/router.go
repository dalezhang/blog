package routers

import (
	"github.com/astaxie/beego"
	"github.com/dalezhang/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/404.html", &controllers.BaseController{}, "*:Go404")

	beego.Router("/article", &controllers.ListArticleController{})
	beego.Router("/article/:id", &controllers.ShowArticleController{})

	beego.Router("/login", &controllers.LoginUserController{})
	beego.Router("/logout", &controllers.LogoutUserController{})

	beego.Router("/article/add", &controllers.AddArticleController{})
	beego.Router("/article/edit/:id", &controllers.EditArticleController{})

	beego.Router("/comment/add", &controllers.AddCommentController{})
	beego.Router("/comment/edit/status", &controllers.EditCommentController{})

	//beego.Router("/album", &controllers.ListAlbumController{})
	//beego.Router("/album/upload", &controllers.UpdateAlbumController{})
	//beego.Router("/album/edit", &controllers.EditAlbumController{})

	//beego.Router("/about", &controllers.AboutUserController{})

	beego.Router("/uploadmulti", &controllers.UploadMultiController{})
	beego.Router("/upload", &controllers.UploadController{})

	//beego.Router("/article/ajax/add", &controller.AddArticleController{}, "*.AddPost")
	//beego.Router("/article/add", &controller.AddArticleController{}, "*.Add")
}
