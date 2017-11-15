package models

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Article struct {
	Id       int
	Title    string
	Uri      string
	Keywords string
	Summary  string
	Content  string
	Author   string
	Created  int64
	Viewnum  int
	Status   int
}

func (this *Article) TableName() string {
	return "article"
}

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:@/blog?charset=utf8", 30)
	orm.RegisterModel(new(Article))
	//orm.RunSyncdb("default", false, true)
}


func GetArticle(id int) (Article, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Id: id}
	err := o.Read(&art)
	return art, err
}

func AddArticle(updArt Article) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := new(Article)

	art.Title = updArt.Title
	art.Uri = updArt.Uri
	art.Keywords = updArt.Keywords
	art.Summary = updArt.Summary
	art.Content = updArt.Content
	art.Author = updArt.Author
	art.Created = time.Now().Unix()
	art.Viewnum = 1
	art.Status = updArt.Status

	id, err := o.Insert(art)
	return id, err
}

func UpdateArticle(id int, updArt Article) error {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Id: id}
	art.Title = updArt.Title
	art.Uri = updArt.Uri
	art.Keywords = updArt.Keywords
	art.Summary = updArt.Summary
	art.Content = updArt.Content
	art.Author = updArt.Author
	art.Created = time.Now().Unix()
	art.Status = updArt.Status
	_, err := o.Update(&art)
	return err
}

func ListArticle(condArr map[string]string, page int, offset int) (num int64, err error, art []Article) {
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	cond := orm.NewCondition()
	if condArr["title"] != "" {
		cond = cond.And("title__icontains", condArr["title"])
	}
	if condArr["keywords"] != "" {
		cond = cond.Or("keywords__incontains", condArr["keywords"])
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset = 10
	}
	start := (page - 1) * offset
	var articles []Article
	num, err1 := qs.Limit(offset, start).All(&articles)
	return num, err1, articles
}

func CountArticle(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	cond := orm.NewCondition()
	if condArr["title"] != "" {
		cond = cond.And("title__icontains", condArr["title"])
	}
	if condArr["keywords"] != "" {
		cond = cond.Or("keywords__icontains", condArr["keywords"])
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}
