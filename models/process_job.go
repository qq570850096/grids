package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ProcessJob struct {
	Id        int   `orm:"column(id);auto"`
	ProcessId int `orm:"column(processId);null" description:"娴佺▼Id"`
	JobId     int `orm:"column(jobId);null" description:"鍚敤鑱屽姟Id"`
}

func (t *ProcessJob) TableName() string {
	return "process_job"
}

func init() {
	orm.RegisterModel(new(ProcessJob))
}

// AddProcessJob insert a new ProcessJob into database and returns
// last inserted Id on success.
func AddProcessJob(m *ProcessJob) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProcessJobById retrieves ProcessJob by Id. Returns error if
// Id doesn't exist
func GetProcessJobById(id int) (v *ProcessJob, err error) {
	o := orm.NewOrm()
	v = &ProcessJob{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProcessJob retrieves all ProcessJob matches certain condition. Returns empty list if
// no records exist
func GetAllProcessJob(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProcessJob))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ProcessJob
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateProcessJob updates ProcessJob by Id and returns error if
// the record to be updated doesn't exist
func UpdateProcessJobById(m *ProcessJob) (err error) {
	o := orm.NewOrm()
	v := ProcessJob{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProcessJob deletes ProcessJob by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProcessJob(id int) (err error) {
	o := orm.NewOrm()
	v := ProcessJob{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProcessJob{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func CheckPidUid(i,j int) bool {
	o := orm.NewOrm()
	return o.QueryTable("process_job").
		Filter("processId",i).Filter("jobId",j).Exist()
}

func GetPj(processid int)(joblist []Job){
	o := orm.NewOrm()
	process := []ProcessJob{}
	o.QueryTable("process_job").
		Filter("processId",processid).All(&process)

	// reverse query
	for _,v := range process{
		tmp := Job{}
		o.QueryTable("job").
			Filter("jobId",v.JobId).One(&tmp)
		joblist = append(joblist,tmp)
	}
	return joblist
}

func DeleteProcessJobP(id int) (err error) {
	o := orm.NewOrm()
	v := ProcessJob{ProcessId: id}
	// ascertain id exists in the database
	_,err = o.Delete(&v,"processId")
	return
}