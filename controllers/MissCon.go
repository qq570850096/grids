package controllers

import (
	"grids/models"
	"grids/utils"
	"strconv"
	"strings"
	"time"
)

type MissCon struct {
	MainController
}

type ret struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
func (c *MissCon) Mission(){
	c.TplName = "mission/myMission.html"
}
func (c *MissCon) WorkPlace(){
	map1 := make(map[string][]string)
	m := models.GetAllProcessApply()
	ids := []string{}
	for _,v := range m {
		tmp := utils.CheckString(v.ProcessName) + "申请审批"
		map1[tmp] = append(map1[tmp],v.ProcessName)
		ids = append(ids,strconv.Itoa(v.Id))
	}
	c.Data["MAP"] = map1
	c.Data["IDS"] = ids
	c.TplName = "mission/workPlace.html"
}

func (c *MissCon) ApplyPlace(){
	map1 := make(map[string][]models.Process)
	m := models.GetAllProcessApply()
	for _,v := range m {
		tmp := utils.CheckString(v.ProcessName) + "申请"
		map1[tmp] = append(map1[tmp],v)
	}
	c.Data["MAP"] = map1
	c.TplName = "mission/applyPlace.html"
}

func(c *MissCon) WorkAssigned(){
	c.TplName = "mission/workAssigned.html"
}
func(c *MissCon) DoMission(){
	c.TplName = "mission/doMission.html"
}

func (c *MissCon) Apply() {
	ret := ret{}
	id,_ := strconv.Atoi(c.GetString("userId"))
	prid,_ := strconv.Atoi(c.GetString("prid"))
	user,_ := models.GetUserById(id)
	pins := models.Processinstance{
		Id:           0,
		InstanceName: strings.Trim(c.GetString("proname"),"模板"),
		ProcessId:    prid,
		UserId:       user.Id,
		CreateTime:   time.Now(),
	}
	insid,err := models.AddProcessinstance(&pins)
	if err != nil {
		ret.Code = 1
		ret.Msg = err.Error()
	} else {
		ret.Code = 0
	}
	tins := models.Taskinstance{
		Id:         0,
		TaskName:   strings.Trim(c.GetString("proname"),"模板"),
		JobId:      user.JobId,
		DepartId:   user.DepartId,
		UserId:     user.Id,
		TaskType:   "审批",
		InstanceId: int(insid),
		TaskSort:   1,
		TaskStatus: "待办理",
	}
	_,err = models.AddTaskinstance(&tins)
	if err != nil {
		ret.Code = 1
		ret.Msg = err.Error()
	} else {
		ret.Code = 0
	}
	c.Data["json"] = ret
	c.ServeJSON()
}