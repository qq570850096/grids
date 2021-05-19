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

	beego.Router("/mission", &controllers.MissCon{},"get:Mission")
	beego.Router("/workPlace", &controllers.MissCon{},"get:WorkPlace")
	beego.Router("/applyPlace", &controllers.MissCon{},"get:ApplyPlace")
	beego.Router("/miss/workAssigned", &controllers.MissCon{},"get:WorkAssigned")
	beego.Router("/miss/getWorkAssigned", &controllers.MissCon{},"get:GetWorkAssigned")
	beego.Router("/miss/doMission", &controllers.MissCon{},"get:DoMission")
	beego.Router("/miss/apply", &controllers.MissCon{},"post:Apply")
	beego.Router("/miss/inCharge", &controllers.MissCon{},"get:Incharge")
	beego.Router("/miss/getInCharge", &controllers.MissCon{},"get:GetIncharge")
	beego.Router("/miss/addUser", &controllers.MissCon{},"post:AddUser")
	beego.Router("/miss/cancel", &controllers.MissCon{},"get:Cancel")
	beego.Router("/miss/publish", &controllers.MissCon{},"get:Publish")
}
