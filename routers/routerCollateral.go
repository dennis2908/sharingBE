package routers

import (
	"github.com/astaxie/beego"
	"api_beego/controllers"
)

func init() {
	beego.Router("/Collateral/GetAllCollateral", &controllers.CollateralController{}, "get:GetAllCollateral")
	beego.Router("/Collateral/GetCollateralByID", &controllers.CollateralController{}, "get:GetCollateralByID")
	beego.Router("/Collateral/CreateCollateral", &controllers.CollateralController{}, "post:CreateCollateral")
	beego.Router("/Collateral/DeleteCollateral", &controllers.CollateralController{}, "delete:DeleteCollateral")
	beego.Router("/Collateral/EditCollateral", &controllers.CollateralController{}, "put:EditCollateral")
}