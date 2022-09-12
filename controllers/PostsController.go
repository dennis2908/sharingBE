package controllers

import (
	models "api_beego/models"
	"context"
	"encoding/json"
	"fmt"
	_ "fmt"
	"time"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/leekchan/accounting"
	_ "github.com/shopspring/decimal"
)

var bm, vvv = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache", "EmbedExpiry": "1200000000"}`)

type PostsController struct {
	beego.Controller
}

type ambilPosts struct {
	Id       int
	Title    string
	Content  string
	Category string
	Status   string
}

type ambilPostsX struct {
	Id       int
	Title    string
	Category string
	Status   string
}

type countPosts struct {
	status string
	value  int
}

type cekPosts struct {
	Title    string
	Content  string
	Category string
	Status   string
}

type cekPostsMin struct {
	Title    string
	Category string
	Status   string
}

func (api *PostsController) GetAllPosts() {
	o := orm.NewOrm()
	o.Using("default")
	sql := "select * from posts limit " + api.Ctx.Input.Param(":limit") + " offset " + api.Ctx.Input.Param(":offset")
	var Posts []models.Posts
	_, err := o.Raw(sql).QueryRows(&Posts)

	if err == nil {
		// ... handle error
		api.Data["json"] = Posts
	}

	api.ServeJSON()
}

func (api *PostsController) GetRowCount() {
	o := orm.NewOrm()
	o.Using("default")
	// res := new(countPosts)
	var list orm.ParamsList
	sql := "select concat(status,':',count(*)) as countData from posts group by status"
	// var countposts []countPosts
	_, err := o.Raw(sql).ValuesFlat(&list)

	if err == nil {
		fmt.Println(list)
		api.Data["json"] = list
	}
	fmt.Println(list)
	api.Data["json"] = list

	api.ServeJSON()
}

func (api *PostsController) GetAllPostsStatus() {
	o := orm.NewOrm()
	o.Using("default")
	sql := "select * from posts where status = '" + api.Ctx.Input.Param(":status") + "' limit " + api.Ctx.Input.Param(":limit") + " offset " + api.Ctx.Input.Param(":offset")
	var Posts []models.Posts
	_, err := o.Raw(sql).QueryRows(&Posts)

	if err == nil {
		// ... handle error
		api.Data["json"] = Posts
	}

	api.ServeJSON()
}

func GetAllCollateralCheck() {
	_, err := bm.Get(context.Background(), "data")
	if err != nil {

		o := orm.NewOrm()
		o.Using("default")
		var sql string
		var Posts []ambilPosts
		sql = "select * from collateral"
		o.Raw(sql).QueryRows(&Posts)

		// put
		bm.Put(context.Background(), "data", Posts, time.Second*10000000)

	}
}

func (api *PostsController) GetPostsByID() {
	// var tt = CollateralIdCheck(api)

	// if tt == 1 {
	// 	api.Ctx.WriteString("Data is incomplete")
	// 	return

	// }

	// idInt, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	o.Using("default")
	// var sql string
	var Posts []models.Posts
	idInt, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	// sql = "select id,title,content,category,status from posts where id = '" + api.Ctx.Input.Param(":id") + "'"

	// var posts []*Post
	qs := o.QueryTable("posts")
	num, _ := qs.Filter("id", idInt).All(&Posts)
	// num, _ := o.Raw("SELECT id,title,content,category,status FROM user WHERE id = ?", idInt).Values(&maps)
	if num > 0 {
		api.Data["json"] = Posts[0]
	}
	api.ServeJSON()
}

func (api *PostsController) DeletePosts() {
	o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Posts []ambilPosts
	id := api.Ctx.Input.Param(":id")
	sql = "delete from posts where id = '" + id + "'"
	o.Raw(sql).QueryRows(&Posts)

	api.Data["json"] = "success delete with id = " + id

	api.ServeJSON()
}

func (api *PostsController) ToTrash() {
	o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Posts []ambilPosts
	id := api.Ctx.Input.Param(":id")
	sql = "update posts set status='thrash' where id = '" + id + "'"
	o.Raw(sql).QueryRows(&Posts)

	api.Data["json"] = "succeed move to thrash with id = " + id

	api.ServeJSON()
}

func AllPostsCheck(api *PostsController) string {
	valid := validation.Validation{}

	// var Posts []*models.Posts

	frm := api.Ctx.Input.RequestBody
	ul := &ambilPosts{}
	json.Unmarshal(frm, ul)
	// fmt.Println(u.Category)

	Title := ul.Title
	Category := ul.Category
	Status := ul.Status
	Content := ul.Content

	u := cekPosts{Title, Content, Category, Status}
	valid.Required(u.Title, "Title")
	valid.Required(u.Content, "Content")
	valid.Required(u.Category, "Category")
	valid.Required(u.Status, "Status")
	valid.MinSize(u.Title, 20, "Title")
	valid.MinSize(u.Content, 200, "Content")
	valid.MinSize(u.Category, 3, "Category")

	if valid.HasErrors() {
		// If there are error messages it means the validation didn't pass
		// Print error message
		for _, err := range valid.Errors {
			return err.Key + err.Message
		}
	}

	if IsValidCategory(Status) == false {
		return "Status is not valid. Choose publish, draft, or thrash"
	}

	return ""
}

func MinPostsCheck(api *PostsController) string {
	valid := validation.Validation{}

	frm := api.Ctx.Input.RequestBody
	u := &ambilPostsX{}
	json.Unmarshal(frm, u)

	Title := u.Title
	Category := u.Category
	Status := u.Status

	um := cekPostsMin{Title, Category, Status}
	valid.Required(um.Title, "Title")
	valid.Required(um.Category, "Category")
	valid.Required(um.Status, "Status")
	valid.MinSize(um.Title, 20, "Title")
	valid.MinSize(um.Category, 3, "Category")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return err.Key + err.Message
		}
	}

	if IsValidCategory(Status) == false {
		return "Status is not valid. Choose publish, draft, or thrash"
	}

	return ""
}

func IsValidCategory(category string) bool {
	switch category {
	case
		"publish",
		"draft",
		"thrash":
		return true
	}
	return false
}

func (api *PostsController) CreatePosts() {

	frm := api.Ctx.Input.RequestBody
	if AllPostsCheck(api) != "" {
		api.Data["json"] = AllPostsCheck(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	u := &ambilPosts{}
	json.Unmarshal(frm, u)
	idInt, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	Title := u.Title
	Content := u.Content
	Category := u.Category
	Status := u.Status
	PostsQry := models.Posts{Id: idInt, Title: Title, Content: Content, Category: Category, Status: Status}

	// insert
	_, err := o.Insert(&PostsQry)
	// sql = "INSERT INTO posts (Title, Content, Category, status,created_date,updated_date) VALUES ('" + Title + "'"
	// sql += ",'" + Content + "','" + Category + "','" + status + "')"
	// _, err := o.Raw(sql).QueryRows(&Posts)

	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully edit data " + api.Ctx.Input.Param(":id")
	api.ServeJSON()
}

func (api *PostsController) EditPosts() {
	frm := api.Ctx.Input.RequestBody
	if AllPostsCheck(api) != "" {
		api.Data["json"] = AllPostsCheck(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	u := &ambilPosts{}
	json.Unmarshal(frm, u)
	idInt, _ := strconv.Atoi(api.Ctx.Input.Param(":id"))
	Title := u.Title
	Content := u.Content
	Category := u.Category
	Status := u.Status
	PostsQry := models.Posts{Id: idInt, Title: Title, Content: Content, Category: Category, Status: Status}

	// insert
	_, err := o.Update(&PostsQry)
	// sql = "INSERT INTO posts (Title, Content, Category, status,created_date,updated_date) VALUES ('" + Title + "'"
	// sql += ",'" + Content + "','" + Category + "','" + status + "')"
	// _, err := o.Raw(sql).QueryRows(&Posts)

	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully edit data " + api.Ctx.Input.Param(":id")
	api.ServeJSON()
}

func (api *PostsController) UpdatePosts() {

	if MinPostsCheck(api) != "" {
		api.Data["json"] = MinPostsCheck(api)
		api.ServeJSON()
		return
	}

	o := orm.NewOrm()
	o.Using("default")

	frm := api.Ctx.Input.RequestBody
	u := &ambilPostsX{}
	json.Unmarshal(frm, u)

	Title := u.Title
	Category := u.Category
	Status := u.Status

	var Posts []ambilPosts

	sql := "update posts set status='" + Status + "',title='" + Title + "', category='" + Category + "' where id = '" + api.Ctx.Input.Param(":id") + "'"
	o.Raw(sql).QueryRows(&Posts)
	api.Data["json"] = "Successfully edit data " + api.Ctx.Input.Param(":id")
	api.ServeJSON()
}
