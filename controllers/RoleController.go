package controllers

import (
	"fmt"
	"grids/models"
	"grids/utils"
	"strconv"
	"time"
)

type RoleCon struct {
	MainController
}

func (c *RoleCon) Get() {
	rolelist := models.GetAllRole()
	id,_ := strconv.Atoi(c.GetString("id"))
	rl := models.GetUr(id)
	ret := make(map[string]interface{})
	ret["list"] = rolelist
	ret["role"] = rl
	c.Data["json"] = ret
	c.ServeJSON()
}

func (c *RoleCon) UpInfo() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "role/roleUpdate.html"
	}else {
		roleId,_ := c.GetInt("roleId")
		tmps := c.GetStrings("resourceIds[]")
		rolename := c.GetString("roleName")
		ids := utils.StringsToInts(tmps)
		role := models.Role{
			Id:             roleId,
			RoleName:       rolename,
			ModifiedTime:   time.Now(),
			CreateUserId:   0,
			ModifiedUserId: 0,
		}
		models.UpdateRoleById(&role)
		models.DeleteMenuRoleByRid(roleId)
		for _, v := range ids {
			mr := models.MenuRole{
				Id:     0,
				MenuId: v,
				RoleId: int(roleId),
			}
			models.AddMenuRole(&mr)
		}
		fmt.Println(ids)
		ret := make(map[string]interface{})
		ret["code"] = 0
		ret["msg"] = "success"
		c.Data["json"] = ret
		c.ServeJSON()
		c.Data["json"] = 0
		c.ServeJSON()
	}
}

func (c *RoleCon) DelRole() {
	ret := make(map[string]interface{})
	roleId,_ := c.GetInt("roleId")
	err := models.DeleteRole(roleId)
	if err == nil {
		ret["code"] = 1
		ret["msg"] = "failed"
	}

	err = models.DeleteMenuRoleByRid(roleId)
	if err == nil {
		ret["code"] = 1
		ret["msg"] = "failed"
	}
	err = models.DeleteUserRoleR(roleId)
	if err == nil {
		ret["code"] = 1
		ret["msg"] = "failed"
	}

	ret["code"] = 0
	ret["msg"] = "success"
	c.Data["json"] = ret
	c.ServeJSON()
	c.Data["json"] = 0
	c.ServeJSON()
}

func (c *RoleCon) ListInfo() {
	if name := c.GetString("key[title]");name == "" {
		roleList := models.GetAllRole()

		ret := make(map[string]interface{})
		ret["code"] = 0
		ret["count"] = len(roleList)
		ret["msg"] = ""
		ret["data"] = roleList
		c.Data["json"] = ret

	} else {
		ret := make(map[string]interface{})
		roleList,_ := models.SearchRole(name)
		ret["code"] = 0
		ret["count"] = len(roleList)
		ret["msg"] = ""
		ret["data"] = roleList
		c.Data["json"] = ret
	}

	c.ServeJSON()
}

func (c *RoleCon) GetAdd() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "role/roleAdd.html"
	}else {
		tmps := c.GetStrings("resourceIds[]")
		rolename := c.GetString("roleName")
		ids := utils.StringsToInts(tmps)
		role := models.Role{
			Id:             0,
			RoleName:       rolename,
			CreateTime:     time.Now(),
			ModifiedTime:   time.Now(),
			CreateUserId:   0,
			ModifiedUserId: 0,
		}
		roleid , _ := models.AddRole(&role)
		for _, v := range ids {
			mr := models.MenuRole{
				Id:     0,
				MenuId: v,
				RoleId: int(roleid),
			}
			models.AddMenuRole(&mr)
		}
		fmt.Println(ids)
		ret := make(map[string]interface{})
		ret["code"] = 0
		ret["msg"] = "success"
		c.Data["json"] = ret
		c.ServeJSON()
	}


}

func (c *RoleCon) GetRoleList() {
	c.TplName = "role/roleList.html"
}

func (c *RoleCon) GetMenu() {
	retdata := []models.TreeOV{}
	if roleId,_ := c.GetInt("roleId"); roleId != 0 {
		menus_role,_ := models.GetAllMenuByRid(roleId)
		menus := models.GetAllMenu()
		for _,v := range menus{
			if v.ParentId == 0 {
				tree := models.TreeOV{
					MenuName: v.MenuName,
					Id:       v.Id,
					Checked:  false,
					Children: getallchildrenCheck(v.Id,menus_role),
				}
				retdata = append(retdata,tree)
			}
		}
	} else {
		menus := models.GetAllMenu()

		for _,v := range menus {
			if v.ParentId == 0 {
				tree := models.TreeOV{
					MenuName: v.MenuName,
					Id:       v.Id,
					Checked:  false,
					Children: getallchildren(v.Id),
				}
				retdata = append(retdata,tree)
			}
		}
	}

	ret := make(map[string]interface{})
	ret["code"] = 0
	ret["data"] = retdata
	ret["msg"] = "success"
	c.Data["json"] = ret
	c.ServeJSON()
}

func getallchildrenCheck(pid int,id []models.MenuRole)( m []models.TreeOV) {
	child := models.GetAllMenubyParent(pid)
	for _,v := range child {
		tree := models.TreeOV{
			MenuName: v.MenuName,
			Id:       v.Id,
			Checked:  contains(v.Id,id),
			Children: nil,
		}
		m = append(m, tree)
	}
	return m
}

func getallchildren(pid int)( m []models.TreeOV) {
	child := models.GetAllMenubyParent(pid)
	for _,v := range child {
		tree := models.TreeOV{
			MenuName: v.MenuName,
			Id:       v.Id,
			Checked:  false,
			Children: nil,
		}
		m = append(m, tree)
	}
	return m
}

func contains (id int,ids []models.MenuRole) bool {
	for _,v := range ids {
		if id == v.MenuId {
			return true
		}
	}
	return false
}