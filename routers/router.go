package routers

import (
	"github.com/astaxie/beego"
	"grids/controllers"
)

func init() {
    beego.Router("/", &controllers.UserConn{},"get:Index")

    //user
	beego.Router("/user/userList", &controllers.UserConn{},"get:List")
	beego.Router("/user/userListInfo", &controllers.UserConn{},"get:ListInfo")
	beego.Router("/user/Update", &controllers.UserConn{},"*:Update")
	beego.Router("/user/UpInfo", &controllers.UserConn{},"*:Update")
	beego.Router("/user/Add", &controllers.UserConn{},"*:Add")
	beego.Router("/user/Add", &controllers.UserConn{},"*:Add")
	beego.Router("/user/Del", &controllers.UserConn{},"post:Del")
	beego.Router("/user/Search", &controllers.UserConn{},"get:Search")

	beego.Router("/role/roleList", &controllers.RoleCon{},"get:Get")
}
