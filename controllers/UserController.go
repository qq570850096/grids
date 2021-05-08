package controllers

import (
	"fmt"
	"grids/models"
	"strconv"
	"time"
)

type UserConn struct {
	MainController
}

func (c *UserConn) Index() {
	c.TplName = "index.html"
}

func (c *UserConn) List() {
	c.TplName = "user/userList.html"
}

func (c *UserConn) Update() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "user/userUpdate.html"
	} else {
		rolel := models.GetAllRole()
		id,_ := strconv.Atoi(c.GetString("userId"))
		models.DeleteUserRoleU(id)
		for _,v := range rolel {
			ro := "role" + strconv.Itoa(v.Id)
			if c.GetString(ro) != "" {
				tmp:=models.UserRole{
					UserId: id,
					RoleId: v.Id,
				}
				models.AddUserRole(&tmp)
			}
		}
		user,_ := models.GetUserById(id)
		user.Gender = c.GetString("gender")
		user.UserName = c.GetString("userName")
		user.RealName = c.GetString("realName")
		user.Phone = c.GetString("phone")
		user.ModifiedTime = time.Now()
		user.JobId,_ = strconv.Atoi(c.GetString("jobId"))
		user.DepartId,_ = strconv.Atoi(c.GetString("departId"))
		err := models.UpdateUserById(user)

		if err != nil {
			c.Data["json"] = 0
		} else {
			c.Data["json"] = 1
		}
		c.ServeJSON()
	}
}

func(c *UserConn) Add(){
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "user/userAdd.html"
	} else {
		user := models.User{}
		user.Gender = c.GetString("gender")
		user.Password = c.GetString("password")
		user.UserName = c.GetString("userName")
		user.RealName = c.GetString("realName")
		user.Phone = c.GetString("phone")
		user.CreateTime = time.Now()
		user.JobId,_ = strconv.Atoi(c.GetString("jobId"))
		user.DepartId,_ = strconv.Atoi(c.GetString("departId"))
		user.CreateTime = time.Now()
		ur := models.UserRole{}
		ur.RoleId,_ = strconv.Atoi(c.GetString("roleName"))

		id,_ := models.AddUser(&user)
		ur.UserId = int(id)
		if _,err :=models.AddUserRole(&ur); err != nil {
			c.Data["json"] = 0
		} else {
			c.Data["json"] = 1
		}
		c.ServeJSON()
	}
}

func (c *UserConn) ListInfo() {
	if name := c.GetString("key[title]");name == "" {
		if users,err := models.UserList(); err != nil{
			fmt.Println(err)
		} else {
			for i:=0;i < len(users);i++{
				if i < len(users) -1 {
					if users[i].Id == users[i+1].Id {
						users[i+1].RoleName = users[i].RoleName + "、" + users[i+1].RoleName
						users = append(users[:i],users[i+1:]...)
						i--
					}
				}
			}
			ret := make(map[string]interface{})
			ret["code"] = 0
			ret["count"] = len(users)
			ret["msg"] = ""
			ret["data"] = users
			c.Data["json"] = ret
		}
	} else {
		ret := make(map[string]interface{})
		users,_ := models.SearchUser(name)
		ret["code"] = 0
		ret["count"] = len(users)
		ret["msg"] = ""
		ret["data"] = users
		c.Data["json"] = ret
	}

	c.ServeJSON()
}

func (c *UserConn) Del(){
	id,_ := strconv.Atoi(c.GetString("id"))
	if err := models.DeleteUser(id); err != nil {
		models.DeleteUserRoleU(id)
		fmt.Println(err)
		c.Data["json"] = 0
	} else {
		c.Data["json"] = 1
	}
	c.ServeJSON()
}

func (c *UserConn) Search(){

}