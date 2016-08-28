package routers

import (
	"github.com/astaxie/beego"
	"github.com/rrmartins/gobreak/go2break/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// clients
	beego.Router("/clients", &controllers.ClientsController{}, "get:GetAll")
	beego.Router("/clients/create", &controllers.ClientsController{}, "post:Post")
	beego.Router("/clients/:id/show", &controllers.ClientsController{}, "get:GetOne")
	beego.Router("/clients/:id/edit", &controllers.ClientsController{}, "put:Put")
	beego.Router("/clients/:id/delete", &controllers.ClientsController{}, "delete:Delete")
}
