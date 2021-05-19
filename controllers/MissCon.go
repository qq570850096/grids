package controllers

import (
	"fmt"
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
		TaskStatus: changestring("审批",1),
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
			if uid != 0 {
				u,_ := models.GetUserById(uid)
				tmp_cpy[i].UserId = u.Id
				tmp_cpy[i].UserName = u.UserName
			}
		}
	}
	c.Data["json"] = 1
	c.ServeJSON()
}

func (c *MissCon) Cancel() {
	//tmp_cpy = tmp_cpy[0:0]
	//tmp = tmp[0:0]
	c.Data["json"] = 1
	c.ServeJSON()
}

func (c *MissCon) Publish() {
	for _,v := range tmp_cpy {
		if v.DepartId == 0 {
			goto RETURN
		}
		x,_ := models.GetProcessinstanceByProId(v.ProcessId)
		if v.UserId == 0 && v.DepartId != 0 && v.JobId != 0 {
			users,_ := models.GetUserByLevel(v.DepartId,v.JobId)
			for _,t := range users {
				tIns := models.Taskinstance{
					Id:         0,
					TaskName:   v.TaskName,
					JobId:      v.JobId,
					DepartId:   v.DepartId,
					UserId:     t.Id,
					TaskType:   v.TaskType,
					InstanceId: x.Id,
					TaskSort:   v.TaskSort,
					TaskStatus: changestring(v.TaskType,v.TaskSort),
				}
				if _,err := models.AddTaskinstance(&tIns);err != nil {
					c.Data["json"] = 1
					fmt.Println(err)
				} else {
					c.Data["json"] = 0
				}
			}
		} else {
			tIns := models.Taskinstance{
				Id:         0,
				TaskName:   v.TaskName,
				JobId:      v.JobId,
				DepartId:   v.DepartId,
				UserId:     v.UserId,
				TaskType:   v.TaskType,
				InstanceId: x.Id,
				TaskSort:   v.TaskSort,
				TaskStatus: changestring(v.TaskType,v.TaskSort),
			}
			if _,err := models.AddTaskinstance(&tIns);err != nil {
				c.Data["json"] = 1
				fmt.Println(err)
			} else {
				c.Data["json"] = 0
			}
		}

	}
	RETURN:
	tmp = tmp_cpy
	c.ServeJSON()
}

func (c *MissCon)GetDoMiss ()  {
	uid,_ := strconv.Atoi(c.GetString("userId"))
	prid,_ := strconv.Atoi(c.GetString("prId"))
	name := c.GetString("name")
	pi,_ := models.GetProcessinstanceByProId(prid)
	tasks,_ := models.GetTaskinstanceByUidInsId(uid,pi.Id,name)
	type ret struct {
		Id int
		TaskName string
		DepartName string
		TaskType string
		InsId int
	}
	var retList []ret
	for _,v := range tasks{
		if( v.TaskStatus == "已办理" || v.TaskStatus == "已通过" ||
			v.TaskStatus == "已驳回" || v.TaskStatus == "已取消" || v.TaskStatus == "未开启") {
			continue
		}
		d,_ := models.GetDepartById(v.DepartId)
		tmp := ret{
			Id: v.Id,
			TaskName:   v.TaskName,
			DepartName: d.DepartName,
			TaskType: v.TaskType,
			InsId:pi.Id,
		}
		retList = append(retList, tmp)
	}
	retmap := make(map[string]interface{})
	retmap["code"] = 0
	retmap["data"] = retList
	retmap["count"] = len(retList)
	retmap["msg"] = "success"
	c.Data["json"] = retmap
	c.ServeJSON()
}

func(c *MissCon)Deal(){
	id,_ := strconv.Atoi(c.GetString("id"))
	uid,_ := strconv.Atoi(c.GetString("uid"))
	tins,_ := models.GetTaskinstanceById(id)
	if tins.TaskType == "执行" {
		tins.TaskStatus = "已办理"
	} else {
		tins.TaskStatus = "已通过"
	}
	if err := models.UpdateTaskinstanceById(tins);err != nil {
		fmt.Println(err)
		c.Data["json"] = 0
	}
	tins.TaskSort++
	v,_ := models.GetTaskinstanceBySort(uid,tins.InstanceId,tins.TaskSort)
	if v.TaskType == "执行" {
		v.TaskStatus = "待办理"
	} else {
		v.TaskStatus = "待审批"
	}
	if err := models.UpdateTaskinstanceById(&v);err != nil {
		fmt.Println(err)
		c.Data["json"] = 0
	}
	c.Data["json"] = 1
	c.ServeJSON()
}
func(c *MissCon)Undeal(){
	id,_ := strconv.Atoi(c.GetString("id"))
	uid,_ := strconv.Atoi(c.GetString("uid"))
	tins,_ := models.GetTaskinstanceById(id)
	if tins.TaskType == "执行" {
		tins.TaskStatus = "已取消"
	} else {
		tins.TaskStatus = "已驳回"
	}
	models.UpdateTaskinstanceById(tins)
	tins.TaskSort--
	v,_ := models.GetTaskinstanceBySort(uid,tins.InstanceId,tins.TaskSort)
	if v.TaskType == "执行" {
		v.TaskStatus = "待办理"
	} else {
		v.TaskStatus = "待审批"
	}
	if err := models.UpdateTaskinstanceById(&v);err != nil {
		fmt.Println(err)
		c.Data["json"] = 0
	}
	c.Data["json"] = 1
	c.ServeJSON()
}


func (c *MissCon) Check() {
	id,_ := strconv.Atoi(c.GetString("id"))
	uid,_ := strconv.Atoi(c.GetString("uid"))
	c.Data["json"] = check(id,uid)
	c.ServeJSON()
}

func changestring(str string,sort int) string{
	if sort != 1 {
		return "未开启"
	}
	if str == "执行" {
		return "待办理"
	} else if str == "审批" {
		return "待审批"
	}
	return ""
}

func check(id,uid int) string {
	var msg string
	t,_ := models.GetTaskinstanceById(id)
	if t.TaskSort > 1 {
		t.TaskSort -= 1
	} else {
		return ""
	}
	v,_ := models.GetTaskinstanceBySort(uid,t.InstanceId,t.TaskSort)
	if v.TaskType == "执行" && v.TaskSort > 1 && (v.TaskStatus == "未开启" || v.TaskStatus == "已取消") {
		msg = v.TaskName+"还未办理，请先办理"+ v.TaskName
	} else if v.TaskType == "审批" && v.TaskStatus == "待审批" {
		msg = v.TaskName+"还未审批，请先审批"+ v.TaskName
	} else if v.TaskType == "审批" && v.TaskSort > 1 && (v.TaskStatus == "未开启" || v.TaskStatus == "已驳回") {
		msg = v.TaskName+"还未审批，请先审批"+ v.TaskName
	} else if v.TaskType == "执行" && v.TaskStatus == "待办理" {
		msg = v.TaskName+"还未办理，请先办理"+ v.TaskName
	}
	return msg
}
