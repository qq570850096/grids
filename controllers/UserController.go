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

func (c *UserConn) Login() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "login.html"
	} else {
		ret := make(map[string]interface{})
		username := c.GetString("username")
		passowrd := c.GetString("password")
		v,_ := models.Login(username,passowrd)
		if v != nil {
			ret["code"] = 0
			ret["msg"] = "success"
			ret["userId"] = v.Id
		} else {
			ret["code"] = 1
			ret["msg"] = "failed"
		}
		c.Data["json"] = ret
		c.ServeJSON()
	}
}

func (c *UserConn) Index() {
	uid := c.Ctx.GetCookie("userId")
	uid_int,_ := strconv.Atoi(uid)
	ops := make(map[string][]models.Menu)
	var childs []models.Menu
	var parentName string
	ur,_ := models.GetUserRoleByUId(uid_int)
	for _,v := range ur {
		menu_roles,_ := models.GetAllMenuByRid(v.RoleId)
		for _,x := range menu_roles {
			menu, _ := models.GetMenuById(x.MenuId)
			if menu.ParentId == 0 {
				childs = models.GetAllMenubyParent(menu.Id)
				parentName = menu.MenuName
			} else {
				for _,y := range childs {
					if menu.Id == y.Id {
						ops[parentName] = append(ops[parentName],y)
					}

				}
			}

		}
	}
	fmt.Println(ops)
	c.Data["Options"] = ops
	c.TplName = "index.html"
}

func (c *UserConn) List() {
	c.TplName = "user/userList.html"
}

func (c *UserConn) Update() {
	if c.Ctx.Request.Method == "GET" {
		departs := models.GetAllDepart()
		jobs := models.GetAllJob()
		c.Data["Depart"] = departs
		c.Data["Jobs"] = jobs
		c.TplName = "user/userUpdate.html"
	} else {
		rolel := models.GetAllRole()
		id, _ := strconv.Atoi(c.GetString("userId"))
		models.DeleteUserRoleU(id)
		for _, v := range rolel {
			ro := "role" + strconv.Itoa(v.Id)
			if c.GetString(ro) != "" {
				tmp := models.UserRole{
					UserId: id,
					RoleId: v.Id,
				}
				models.AddUserRole(&tmp)
			}
		}
		user, _ := models.GetUserById(id)
		user.Gender = c.GetString("gender")
		user.UserName = c.GetString("userName")
		user.RealName = c.GetString("realName")
		user.Phone = c.GetString("phone")
		user.ModifiedTime = time.Now()
		user.JobId, _ = strconv.Atoi(c.GetString("jobId"))
		user.DepartId, _ = strconv.Atoi(c.GetString("departId"))
		err := models.UpdateUserById(user)

		if err != nil {
			c.Data["json"] = 0
		} else {
			c.Data["json"] = 1
		}
		c.ServeJSON()
	}
}

func (c *UserConn) Add() {
	if c.Ctx.Request.Method == "GET" {
		departs := models.GetAllDepart()
		jobs := models.GetAllJob()
		c.Data["Depart"] = departs
		c.Data["Jobs"] = jobs
		c.TplName = "user/userAdd.html"
	} else {
		user := models.User{}
		user.Gender = c.GetString("gender")
		user.Password = c.GetString("password")
		user.UserName = c.GetString("userName")
		user.RealName = c.GetString("realName")
		user.Phone = c.GetString("phone")
		user.CreateTime = time.Now()
		user.JobId, _ = strconv.Atoi(c.GetString("jobId"))
		user.DepartId, _ = strconv.Atoi(c.GetString("departId"))
		user.CreateTime = time.Now()

		id, _ := models.AddUser(&user)
		rolel := models.GetAllRole()
		models.DeleteUserRoleU(int(id))
		for _, v := range rolel {
			ro := "role" + strconv.Itoa(v.Id)
			if c.GetString(ro) != "" {
				tmp := models.UserRole{
					UserId: int(id),
					RoleId: v.Id,
				}
				models.AddUserRole(&tmp)
			}
		}
		c.Data["json"] = 1
		c.ServeJSON()
	}
}

func (c *UserConn) ListInfo() {
	if name := c.GetString("key[title]"); name == "" {
		if users, err := models.UserList(); err != nil {
			fmt.Println(err)
		} else {
			for i := 0; i < len(users); i++ {
				if i < len(users)-1 {
					if users[i].Id == users[i+1].Id {
						users[i+1].RoleName = users[i].RoleName + "???" + users[i+1].RoleName
						users = append(users[:i], users[i+1:]...)
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
		users, _ := models.SearchUser(name)
		ret["code"] = 0
		ret["count"] = len(users)
		ret["msg"] = ""
		ret["data"] = users
		c.Data["json"] = ret
	}

	c.ServeJSON()
}

func (c *UserConn) Del() {
	id, _ := strconv.Atoi(c.GetString("id"))
	if err := models.DeleteUser(id); err != nil {
		models.DeleteUserRoleU(id)
		fmt.Println(err)
		c.Data["json"] = 0
	} else {
		c.Data["json"] = 1
	}
	c.ServeJSON()
}


