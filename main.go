package main

import (
	_ "api_beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "api_beego/routers"
	"api_beego/models"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
)

func init(){ // init instead of int
     beego.Debug("Filters init...")

	// CORS for https://foo.* origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://127.0.0.1:8112"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
    orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", 
        "postgres",
        "user=postgres password=12345 host=127.0.0.1 port=5432 dbname=auction sslmode=disable");
	orm.RegisterModel(new(models.Collateral))	
    orm.RunSyncdb("default", false, true)
	orm.RunCommand()
}
 func main() {
	    _, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache", "EmbedExpiry": "120"}`)
	
		if err != nil {
			logs.Error(err)
		}
		beego.Run()
 }