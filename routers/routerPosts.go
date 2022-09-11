package routers

import (
	"api_beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/article/:limit/:offset", &controllers.PostsController{}, "get:GetAllPosts")
	beego.Router("/article/:limit/:offset/:status", &controllers.PostsController{}, "get:GetAllPostsStatus")
	beego.Router("/CountArticle/", &controllers.PostsController{}, "get:GetRowCount")
	beego.Router("/article/:id", &controllers.PostsController{}, "get:GetPostsByID")
	beego.Router("/article/", &controllers.PostsController{}, "post:CreatePosts")
	beego.Router("/article/:id", &controllers.PostsController{}, "delete:DeletePosts")
	beego.Router("/article/:id", &controllers.PostsController{}, "put,post,patch:EditPosts")
	beego.Router("/articleUpdate/:id", &controllers.PostsController{}, "put,post,patch:UpdatePosts")
}
