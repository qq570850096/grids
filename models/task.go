package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Task struct {
	Id        int    `orm:"column(taskId);auto" description:"浠诲姟Id"`
	TaskName  string `orm:"column(taskName);size(255);null" description:"浠诲姟鍚嶇О"`
	JobId     int  `orm:"column(jobId);null" description:"璐熻矗鑱屽姟Id"`
	DepartId  int  `orm:"column(departId);null" description:"璐熻矗閮ㄩ棬Id"`
	UserId    int  `orm:"column(userId);null" description:"璐熻矗浜篒d"`
	TaskType  string `orm:"column(taskType);size(10);null" description:"浠诲姟绫诲瀷锛堟墽琛屻�佸鎵癸級"`
	ProcessId int  `orm:"column(processId);null" description:"鎵�灞炴祦绋婭d"`
	TaskSort  int    `orm:"column(taskSort);null" description:"浠诲姟椤哄簭"`
}

type Task_tmp struct {
	Id        int    `orm:"column(taskId);auto" description:"浠诲姟Id"`
	TaskName  string `orm:"column(taskName);size(255);null" description:"浠诲姟鍚嶇О"`
	JobId     int  `orm:"column(jobId);null" description:"璐熻矗鑱屽姟Id"`
	DepartId  int  `orm:"column(departId);null" description:"璐熻矗閮ㄩ棬Id"`
	UserId    int  `orm:"column(userId);null" description:"璐熻矗浜篒d"`
	TaskType  string `orm:"column(taskType);size(10);null" description:"浠诲姟绫诲瀷锛堟墽琛屻�佸鎵癸級"`
	ProcessId int  `orm:"column(processId);null" description:"鎵�灞炴祦绋婭d"`
	TaskSort  int    `orm:"column(taskSort);null" description:"浠诲姟椤哄簭"`
	JobName string  `orm:"column(jobName);size(50);null" description:"鑱屽姟鍚嶇О"`
}

func (t *Task) TableName() string {
	return "task"
}

func init() {
	orm.RegisterModel(new(Task))
}

// AddTask insert a new Task into database and returns
// last inserted Id on success.
func AddTask(m *Task) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTaskById retrieves Task by Id. Returns error if
// Id doesn't exist
func GetTaskById(id int) (v *Task, err error) {
	o := orm.NewOrm()
	v = &Task{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTask retrieves all Task matches certain condition. Returns empty list if
// no records exist
func GetAllTask(prid int) (tasks []Task) {
	o := orm.NewOrm()

	if _,err:=o.QueryTable("task").Filter("processId",prid).All(&tasks);err != nil {
		fmt.Println(err)
		return
	}
	return
}

// UpdateTask updates Task by Id and returns error if
// the record to be updated doesn't exist
func UpdateTaskById(m *Task) (err error) {
	o := orm.NewOrm()
	v := Task{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTask deletes Task by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTask(id int) (err error) {
	o := orm.NewOrm()
	v := Task{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Task{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetAllTaskByProcessT(pid int) (t []Task) {
	o := orm.NewOrm()
	o.QueryTable("task").Filter("processId",pid).All(&t)
	return t
}
