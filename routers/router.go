package routers

import (
	"beego_web_mysql/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/testjson", &controllers.RequestController{})
	beego.Router("/api/measurements", &controllers.MeasurementsController{})
}
