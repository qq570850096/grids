package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Process struct {
	Id             int       `orm:"column(processId);auto" description:"娴佺▼Id"`
	ProcessName    string    `orm:"column(processName);size(255);null" description:"娴佺▼鍚嶇О"`
	ProcessType    string    `orm:"column(processType);size(10);null" description:"娴佺▼绫诲瀷锛堟寚娲俱�佺敵璇凤級"`
	CreateTime     time.Time `orm:"column(createTime);type(datetime);null" description:"鍒涘缓鏃堕棿"`
	ModifiedTime   time.Time `orm:"column(modifiedTime);type(datetime);null" description:"淇敼鏃堕棿"`
	CreateUserId   int64     `orm:"column(createUserId);null" description:"鍒涘缓鑰匢d"`
	ModifiedUserId int64     `orm:"column(modifiedUserId);null" description:"淇敼鑰匢d"`
}

type ProcessTask struct {
	P Process
	Tasks			[]Task
}

type ProcessJ struct {
	Id             int       `orm:"column(processId);auto" description:"娴佺▼Id"`
	ProcessName    string    `orm:"column(processName);size(255);null" description:"娴佺▼鍚嶇О"`
	ProcessType    string    `orm:"column(processType);size(10);null" description:"娴佺▼绫诲瀷锛堟寚娲俱�佺敵璇凤級"`
	JobName        string    `orm:"column(jobName);size(50);null" description:"鑱屽姟鍚嶇О"`
	CreateTime     time.Time `orm:"column(createTime);type(datetime);null" description:"鍒涘缓鏃堕棿"`
	ModifiedTime   time.Time `orm:"column(modifiedTime);type(datetime);null" description:"淇敼鏃堕棿"`
}


func (t *Process) TableName() string {
	return "process"
}

func init() {
	orm.RegisterModel(new(Process))
}

// AddProcess insert a new Process into database and returns
// last inserted Id on success.
func AddProcess(m *Process) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProcessById retrieves Process by Id. Returns error if
// Id doesn't exist
func GetProcessById(id int) (v *Process, err error) {
	o := orm.NewOrm()
	v = &Process{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProcess retrieves all Process matches certain condition. Returns empty list if
// no records exist
func GetAllProcessApply() (m []Process){
	o := orm.NewOrm()
	o.QueryTable("process").Filter("processType","申请").All(&m)
	return m
}

func GetAllProcess() (m []Process){
	o := orm.NewOrm()
	o.QueryTable("process").All(&m)
	return m
}

func GetAllProcessCheck() (m []ProcessTask){
	plist := make([]Process,0)
	o := orm.NewOrm()
	o.QueryTable("process").Filter("processType","指派").All(&plist)

	for _,v := range plist{
		tmp := ProcessTask{
			P:     v,
			Tasks: []Task{},
		}
		tmp.Tasks = GetAllTaskByProcessT(v.Id)
		m = append(m,tmp)
	}

	return m
}
// UpdateProcess updates Process by Id and returns error if
// the record to be updated doesn't exist
func UpdateProcessById(m *Process) (err error) {
	o := orm.NewOrm()
	v := Process{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProcess deletes Process by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProcess(id int) (err error) {
	o := orm.NewOrm()
	v := Process{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Process{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func ProcessList() (processlist []ProcessJ, err error) {
	o := orm.NewOrm()
	if _,err = o.Raw("SELECT p.*,j.jobName " +
		"FROM process p, process_job pj, job j " +
		"WHERE p.processId = pj.processId AND j.jobId = pj.jobId").QueryRows(&processlist); err != nil {
		return nil ,err
	}
	return processlist,err
}

func SearchProcess(name string) (processlist []ProcessJ,err error) {
	o := orm.NewOrm()
	name1 := "%" + name + "%"
	name2 := "%" + name + "%"
	if _,err = o.Raw("SELECT p.*,j.jobName " +
		"FROM process p, process_job pj, job j " +
		"WHERE (p.processId = pj.processId AND j.jobId = pj.jobId AND j.jobName LIKE ?) OR "+
		"(p.processId = pj.processId AND j.jobId = pj.jobId AND p.processType LIKE ?)", name1, name2).QueryRows(&processlist); err != nil {
		return nil, err
		fmt.Println(err)
	}
	return processlist,nil
}

