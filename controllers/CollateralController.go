package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/shopspring/decimal"
    _ "github.com/leekchan/accounting"
	_ "github.com/astaxie/beego/validation"
	models "api_beego/models"
	"context"
	"time"
	
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/core/logs"
	_ "strconv"
)

var bm,vvv = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache", "EmbedExpiry": "1200000000"}`)

type CollateralController struct {
	beego.Controller
}

type ambilCollateral struct {
    Coll_id int
	Acc_id  int
	Type_id int	
	Owner_name string
	Coll_location string
	Initial_col_price int 
	Final_col_price int
	Ljk_id int
	Document_path string
}

func (api *CollateralController) GetAllCollateral() {
    GetAllCollateralCheck() 
	Result, _ := bm.Get(context.Background(), "data")
	api.Data["json"] = Result
	api.ServeJSON()
}

func GetAllCollateralCheck() {
	_, err := bm.Get(context.Background(), "data")
	if err != nil {
	
    o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Collateral [] ambilCollateral
	sql = "select * from collateral"
	o.Raw(sql).QueryRows(&Collateral)
	
	// put
	bm.Put(context.Background(), "data", Collateral, time.Second*10000000)
	
	}
}

func GetAllCollateralRenew() {
    o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Collateral [] ambilCollateral
	sql = "select * from collateral"
	o.Raw(sql).QueryRows(&Collateral)
	
	// put
	bm.Put(context.Background(), "data", Collateral, time.Second*10000000)
}

func (api *CollateralController) GetCollateralByID() {
    var tt = CollateralIdCheck(api)

	if tt == 1 {
		api.Ctx.WriteString("Data is incomplete")
		return

	}
    o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Collateral [] ambilCollateral
	sql = "select * from collateral where coll_id = '"+api.GetString("coll_id")+"'"
	num, err := o.Raw(sql).QueryRows(&Collateral)
	if err != orm.ErrNoRows && num > 0 {
		api.Data["json"] = Collateral[0]
	}
	api.ServeJSON()
}

func (api *CollateralController) DeleteCollateral() {
    var tt = CollateralIdCheck(api)

	if tt == 1 {
		api.Ctx.WriteString("Data is incomplete")
		return

	}
    o := orm.NewOrm()
	o.Using("default")
	var sql string
	var Collateral [] ambilCollateral
	coll_id := api.GetString("coll_id")
	sql = "delete from collateral where coll_id = '"+coll_id+"'"
	o.Raw(sql).QueryRows(&Collateral)
	
	GetAllCollateralRenew()
	
	api.Data["json"] = "success delete user account with id = "+coll_id
	
	api.ServeJSON()
}

func CollateralIdCheck(api  *CollateralController) int{
    coll_id := api.GetString("coll_id")
    if coll_id == "" {
        return 1
	}
	return 0
	
	
}	

func AllCollateralCheck(api  *CollateralController) int{
    acc_id := api.GetString("acc_id")
    if acc_id == "" {
        return 1
        
    } 
	type_id := api.GetString("type_id")
    if type_id == "" {
        return 1
        
    } 
	owner_name := api.GetString("owner_name")
    if owner_name == "" {
        return 1
        
    } 
	coll_location := api.GetString("coll_location")
    if coll_location == "" {
        return 1
        
    } 
	initial_col_price := api.GetString("initial_col_price")
    if initial_col_price == "" {
        return 1
        
    } 
	final_col_price := api.GetString("final_col_price")
    if final_col_price == "" {
        return 1
        
    } 
	ljk_id := api.GetString("ljk_id")
    if ljk_id == "" {
		return 1
        
    }
	document_path := api.GetString("document_path")
    if document_path == "" {
        return 1
        
	}
	
	return 0
}

func (api *CollateralController) CreateCollateral() {

	var tt = AllCollateralCheck(api)

	if tt == 1 {
		api.Ctx.WriteString("Data is incomplete")
		return

	}
    o := orm.NewOrm()
	o.Using("default")
	var Collateral []*models.Collateral
	var sql string
	sql = "INSERT INTO collateral (acc_id, type_id, owner_name, coll_location, initial_col_price, final_col_price, ljk_id, document_path) VALUES ('"+api.GetStrings("acc_id")[0]+"'"
	sql += ",'"+api.GetStrings("type_id")[0]+"','"+api.GetStrings("owner_name")[0]+"','"+api.GetStrings("coll_location")[0]+"','"+api.GetStrings("initial_col_price")[0]+"'"
	sql += ",'"+api.GetStrings("final_col_price")[0]+"','"+api.GetStrings("ljk_id")[0]+"','"+api.GetStrings("document_path")[0]+"')"
	o.Raw(sql).QueryRows(&Collateral)
	GetAllCollateralRenew()
	api.Data["json"] = "Successfully add new data"
	api.ServeJSON()
}


func (api *CollateralController) EditCollateral() {
    var tt = AllCollateralCheck(api)

	if tt == 1 {
		api.Ctx.WriteString("Data is incomplete")
		return
	}
    o := orm.NewOrm()
	o.Using("default")
	var Collateral []*models.Collateral
	var sql string
	sql = "select coll_id from collateral where coll_id = '"+api.GetString("coll_id")+"'";
	num, err := o.Raw(sql).QueryRows(&Collateral)
	if err != orm.ErrNoRows && num == 0 {
		api.Data["json"] = "91"
	    api.ServeJSON()   
	}
	sql = "UPDATE collateral SET acc_id = '"+api.GetStrings("acc_id")[0]+"', type_id = '"+api.GetStrings("type_id")[0]+"', owner_name = '"+api.GetStrings("owner_name")[0]+"'"
	sql += ", coll_location = '"+api.GetStrings("coll_location")[0]+"', initial_col_price = '"+api.GetStrings("initial_col_price")[0]+"', final_col_price = '"+api.GetStrings("final_col_price")[0]+"'"
	sql += ", ljk_id = '"+api.GetStrings("ljk_id")[0]+"', document_path = '"+api.GetStrings("document_path")[0]+"' where coll_id = '"+api.GetStrings("coll_id")[0]+"'" 
	o.Raw(sql).QueryRows(&Collateral)
	GetAllCollateralRenew()
	api.Data["json"] = "successfully edit collateral with coll_id "+api.GetStrings("coll_id")[0]
	api.ServeJSON()
}