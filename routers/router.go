package routers

import (
	"tutorial2_web_mysql_simple/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/testjson", &controllers.RequestController{})
	beego.Router("/api/measurements", &controllers.MeasurementsController{})
}
