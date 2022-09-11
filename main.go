package main

import (
	"api_beego/models"
	_ "api_beego/routers"
	"log"
	"os"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/lib/pq"
)

func init() { // init instead of int
	beego.Debug("Filters init...")

	// CORS for https://foo.* origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default",
		"postgres",
		"user=sepzsrojgsurnm password=8471a6d259354847f27e3bc04f9b674b19038278e55e8ee9e94275e2a48310a3 host=ec2-52-200-5-135.compute-1.amazonaws.com port=5432 dbname=d79rc60ef2fovr sslmode=require")
	orm.RegisterModel(new(models.Posts))
	orm.RunSyncdb("default", false, true)
	orm.RunCommand()
}
func main() {
	_, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache", "EmbedExpiry": "120"}`)

	if err != nil {
		logs.Error(err)
	}
	log.Println("Env $PORT :", os.Getenv("PORT"))
	if os.Getenv("PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
			log.Fatal("$PORT must be set")
		}
		log.Println("port : ", port)
		beego.BConfig.Listen.HTTPPort = port
		beego.BConfig.Listen.HTTPSPort = port
	}

	beego.Run()
}
