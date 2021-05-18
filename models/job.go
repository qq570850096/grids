package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Job struct {
	Id             int       `orm:"column(jobId);auto" description:"鑱屽姟Id"`
	JobName        string    `orm:"column(jobName);size(50);null" description:"鑱屽姟鍚嶇О"`
	SuperiorId     int64     `orm:"column(superiorId);null" description:"涓婄骇鑱屽姟Id"`
	CreateTime     time.Time `orm:"column(createTime);type(datetime);null" description:"鍒涘缓鏃堕棿"`
	ModifiedTime   time.Time `orm:"column(modifiedTime);type(datetime);null" description:"淇敼鏃堕棿"`
	CreateUserId   int64     `orm:"column(createUserId);null" description:"鍒涘缓鑰匢d"`
	ModifiedUserId int64     `orm:"column(modifiedUserId);null" description:"淇敼鑰匢d"`
}

func (t *Job) TableName() string {
	return "job"
}

func init() {
	orm.RegisterModel(new(Job))
}

// AddJob insert a new Job into database and returns
// last inserted Id on success.
func AddJob(m *Job) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetJobById retrieves Job by Id. Returns error if
// Id doesn't exist
func GetJobById(id int) (v *Job, err error) {
	o := orm.NewOrm()
	v = &Job{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllJob retrieves all Job matches certain condition. Returns empty list if
// no records exist
func GetAllJob() (jobs []Job) {
	o := orm.NewOrm()
	o.QueryTable("job").All(&jobs)
	return jobs
}

// UpdateJob updates Job by Id and returns error if
// the record to be updated doesn't exist
func UpdateJobById(m *Job) (err error) {
	o := orm.NewOrm()
	v := Job{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteJob deletes Job by Id and returns error if
// the record to be deleted doesn't exist
func DeleteJob(id int) (err error) {
	o := orm.NewOrm()
	v := Job{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Job{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
