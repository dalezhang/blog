package main

import (
	"github.com/astaxie/beego"
	_ "github.com/dalezhang/blog/initial"
	_ "github.com/dalezhang/blog/routers"
	//"fmt"
	"html/template"
	"net/http"
)

func main() {
	beego.ErrorHandler("404", page_not_found)
	beego.Run()
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")

	data := make(map[string]interface{})
	data["content"] = "page not fount"
	t.Execute(rw, data)
}
