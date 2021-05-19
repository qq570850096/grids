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
type task_tmp struct {
	models.Task
	JobName string
	DepartName string
	UserName string
}
var tmp []task_tmp
var tmp_cpy []task_tmp
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
	ch := models.GetAllProcessCheck()
	for _,v := range m {
		tmp := utils.CheckString(v.ProcessName) + "申请审批"
		map1[tmp] = append(map1[tmp],v.ProcessName)
	}
	c.Data["MAP"] = map1
	c.Data["Miss"] = ch
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
		TaskStatus: changestring("审批"),
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

func(c *MissCon) GetWorkAssigned() {

	pid,_ := strconv.Atoi(c.GetString("pid"))


	tasks := models.GetAllTaskByProcessT(pid)
	ret := make(map[string]interface{})

	if len(tmp_cpy) > 0{
		goto JSON
	}
	for _,v := range tasks {
		job,_ := models.GetJobById(v.JobId)
		tmpl := task_tmp{
			v,
			job.JobName,
			"",
			"",
		}
		if v.DepartId != 0 {
			d,_ := models.GetDepartById(v.DepartId)
			tmpl.DepartName = d.DepartName
		}
		if v.UserId != 0 {
			d,_ := models.GetUserById(v.UserId)
			tmpl.UserName = d.UserName
		}
		tmpl.JobName = job.JobName
		tmp = append(tmp,tmpl)
	}
	tmp_cpy = tmp
	JSON:
	ret["code"] = 0
	ret["msg"] = "success"
	ret["len"] = len(tasks)
	ret["data"] = tmp_cpy
	c.Data["json"] = ret
	c.ServeJSON()
}

func (c *MissCon)Incharge()  {
	if c.Ctx.Request.Method == "GET" {
		departs := models.GetAllDepart()
		jobs := models.GetAllJob()
		c.Data["Depart"] = departs
		c.Data["Jobs"] = jobs
		c.Data["IsEmpty"] = true
		c.TplName = "mission/inCharge.html"
	}
}

func (c *MissCon)GetIncharge()  {
	dep,_ := strconv.Atoi(c.GetString("departName"))
	job,_ := strconv.Atoi(c.GetString("jobName"))
	if dep != 0 && job != 0 {
		users,_ := models.GetUserByLevel(dep,job)
		if len(users) > 0 {
			c.Data["json"] = users
		}
	} else {
		c.Data["json"] = nil
	}
	c.ServeJSON()
}

func (c *MissCon) AddUser() {
	task_id,_ := strconv.Atoi(c.GetString("taskId"))
	dep,_ := strconv.Atoi(c.GetString("departName"))
	job,_ := strconv.Atoi(c.GetString("jobName"))
	uid,_ := strconv.Atoi(c.GetString("user"))
	for i,v := range tmp_cpy {
		if v.Id == task_id {
			j,_ := models.GetJobById(job)
			tmp_cpy[i].JobId = j.Id
			tmp_cpy[i].JobName = j.JobName
			d,_ := models.GetDepartById(dep)
			tmp_cpy[i].DepartId = d.Id
			tmp_cpy[i].DepartName = d.DepartName
			u,_ := models.GetUserById(uid)
			tmp_cpy[i].UserId = u.Id
			tmp_cpy[i].UserName = u.UserName
		}
	}
	c.Data["json"] = 1
	c.ServeJSON()
}

func (c *MissCon) Cancel() {
	tmp_cpy = tmp_cpy[0:0]
	tmp = tmp[0:0]
	c.Data["json"] = 1
	c.ServeJSON()
}

func (c *MissCon) Publish() {
	for _,v := range tmp_cpy {
		m := v.Task
		models.UpdateTaskById(&m)
	}
	tmp = tmp_cpy
	c.Data["json"] = 0
	c.ServeJSON()
}

func changestring(str string) string{
	if str == "执行" {
		return "待办理"
	} else if str == "审批" {
		return "待审批"
	}
	return ""
}
