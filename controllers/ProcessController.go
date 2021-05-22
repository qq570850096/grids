package controllers

import (
	"fmt"
	"grids/models"
	"strconv"
	"time"
)

type ProcessConn struct {
	MainController
}

func (c *ProcessConn) List() {
	c.TplName = "process/processList.html"
}

func (c *ProcessConn) ListInfo() {
	if name := c.GetString("key[title]");name == "" {
		if processes,err := models.ProcessList(); err != nil{
			fmt.Println(err)
		} else {
			for i:=0;i < len(processes);i++{
				if i < len(processes) -1 {
					if processes[i].Id == processes[i+1].Id {
						processes[i+1].JobName = processes[i].JobName + "、" + processes[i+1].JobName
						processes = append(processes[:i],processes[i+1:]...)
						i--
					}
				}
			}
			ret := make(map[string]interface{})
			ret["code"] = 0
			ret["count"] = len(processes)
			ret["msg"] = ""
			ret["data"] = processes
			c.Data["json"] = ret
		}
	} else {
		ret := make(map[string]interface{})
		processes,_ := models.SearchProcess(name)
		for i:=0;i < len(processes);i++{
			if i < len(processes) -1 {
				if processes[i].Id == processes[i+1].Id {
					processes[i+1].JobName = processes[i].JobName + "、" + processes[i+1].JobName
					processes = append(processes[:i],processes[i+1:]...)
					i--
				}
			}
		}
		ret["code"] = 0
		ret["count"] = len(processes)
		ret["msg"] = ""
		ret["data"] = processes
		c.Data["json"] = ret
	}

	c.ServeJSON()
}

func(c *ProcessConn) Add(){
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "process/processAdd.html"
	} else {
		process:= models.Process{}
		process.ProcessName = c.GetString("processName")
		process.ProcessType = c.GetString("processType")
		process.CreateTime = time.Now()
		process.ModifiedTime = time.Now()

		id,_ := models.AddProcess(&process)
		jobs := models.GetAllJob()
		models.DeleteProcessJobP(int(id))
		for _,v := range jobs {
			job := "job" + strconv.Itoa(v.Id)
			if c.GetString(job) != "" {
				tmp:=models.ProcessJob{
					ProcessId: int(id),
					JobId: v.Id,
				}
				models.AddProcessJob(&tmp)
			}
		}
		c.Data["json"] = 1
		c.ServeJSON()
	}
}

//task相关还没删除
func (c *ProcessConn) Del(){
	id,_ := strconv.Atoi(c.GetString("id"))
	if err := models.DeleteProcess(id); err != nil {
		models.DeleteProcessJobP(id)
		fmt.Println(err)
		c.Data["json"] = 0
	} else {
		c.Data["json"] = 1
	}
	c.ServeJSON()
}

func (c *ProcessConn) Update() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "process/processUpdate.html"
	} else {
		jobs := models.GetAllJob()
		id,_ := strconv.Atoi(c.GetString("processId"))
		models.DeleteProcessJobP(id)
		for _,v := range jobs {
			ro := "job" + strconv.Itoa(v.Id)
			if c.GetString(ro) != "" {
				tmp:=models.ProcessJob{
					ProcessId: id,
					JobId: v.Id,
				}
				models.AddProcessJob(&tmp)
			}
		}
		process,_ := models.GetProcessById(id)
		process.ProcessName = c.GetString("processName")
		process.ProcessType = c.GetString("processType")
		process.ModifiedTime = time.Now()
		err := models.UpdateProcessById(process)

		if err != nil {
			c.Data["json"] = 0
		} else {
			c.Data["json"] = 1
		}
		c.ServeJSON()
	}
}

func (c *ProcessConn) Task() {
	departs := models.GetAllDepart()
	jobs := models.GetAllJob()
	users,_ := models.GetAllUser()
	prid,_ := strconv.Atoi(c.GetString("prid"))
	task := models.GetAllTask(prid)
	fmt.Println(task)
	if len(task) == 0 {
		c.Data["IsEmpty"] = true
	} else {
		c.Data["IsEmpty"] = false
	}
	c.Data["Tasks"] = task
	c.Data["Departs"] = departs
	c.Data["Jobs"] = jobs
	c.Data["Users"] = users
	str:=strconv.Itoa(len(task))
	c.Ctx.SetCookie("count",str,120)
	fmt.Println(c.Ctx.GetCookie("count"))
	c.TplName = "process/editTask.html"
}

func(c *ProcessConn) TaskInfo(){
	
	str1 := "taskId"
	str2 := "jobId"
	str3 := "departId"
	str4 := "userId"
	str5 := "taskName"
	str6 := "taskType"
	index := 1
	procid,_ := strconv.Atoi(c.GetString("procId"))
	for {

		taskId,_ := strconv.Atoi(c.GetString(fmt.Sprintf("%s%d",str1,index)))
		userid,_ := strconv.Atoi(c.GetString(fmt.Sprintf("%s%d",str4,index)))
		jobid,_ := strconv.Atoi(c.GetString(fmt.Sprintf("%s%d",str2,index)))
		depid,_ := strconv.Atoi(c.GetString(fmt.Sprintf("%s%d",str3,index)))
		taskname := fmt.Sprintf("%s%d",str5,index)
		fmt.Println(c.GetString(taskname))
		// 如果 taskId 还不为空则说明该数据是从后端返回的，是需要更新的数据
		if c.GetString(fmt.Sprintf("%s%d",str1,index))!="" {

			task := models.Task{
				Id:        taskId,
				TaskName:  c.GetString(taskname),
				JobId:     jobid,
				DepartId:  depid,
				UserId:    userid,
				TaskType:  c.GetString(fmt.Sprintf("%s%d",str6,index)),
				ProcessId: procid,
				TaskSort:  index,
			}
			if err := models.UpdateTaskById(&task);err != nil {
				c.Data["json"] = 0
				break
			}
			index++
			// 如果 taskId 还为空则说明该数据是前端编辑的，有可能还有新增数据

		} else if c.GetString(taskname)!= "" {
			task := models.Task{
				Id:        taskId,
				TaskName:  c.GetString(taskname),
				JobId:     jobid,
				DepartId:  depid,
				UserId:    userid,
				TaskType:  c.GetString(fmt.Sprintf("%s%d",str6,index)),
				ProcessId: procid,
				TaskSort:  index,
			}
			if _,err := models.AddTask(&task);err != nil {
				fmt.Println(err)
				c.Data["json"] = 0
				break
			}
			index++
		} else {
			break
		}
	}
	c.Data["json"] = 1
	c.ServeJSON()
}