package controllers

import (
	"grids/models"
	"strconv"
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
