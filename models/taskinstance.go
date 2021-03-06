package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Taskinstance struct {
	Id         int    `orm:"column(taskId);auto" description:"娴佺▼瀹炰緥鐨勪换鍔d"`
	TaskName   string `orm:"column(taskName);size(255);null" description:"浠诲姟鍚嶇О"`
	JobId      int  `orm:"column(jobId);null" description:"璐熻矗鑱屽姟Id"`
	DepartId   int  `orm:"column(departId);null" description:"璐熻矗閮ㄩ棬Id"`
	UserId     int  `orm:"column(userId);null" description:"璐熻矗浜篒d"`
	TaskType   string `orm:"column(taskType);size(10);null" description:"浠诲姟绫诲瀷锛堟墽琛屻�佸鎵癸級"`
	InstanceId int  `orm:"column(instanceId);null" description:"鎵�灞炴祦绋嬪疄渚婭d"`
	TaskSort   int    `orm:"column(taskSort);null" description:"浠诲姟椤哄簭"`
	TaskStatus string `orm:"column(taskStatus);size(10);null" description:"浠诲姟鐘舵��"`
}

func (t *Taskinstance) TableName() string {
	return "taskinstance"
}

func init() {
	orm.RegisterModel(new(Taskinstance))
}

// AddTaskinstance insert a new Taskinstance into database and returns
// last inserted Id on success.
func AddTaskinstance(m *Taskinstance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTaskinstanceById retrieves Taskinstance by Id. Returns error if
// Id doesn't exist
func GetTaskinstanceById(id int) (v *Taskinstance, err error) {
	o := orm.NewOrm()
	v = &Taskinstance{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetTaskinstanceByUidInsId(uid,insid int,name string) (v []Taskinstance, err error) {
	o := orm.NewOrm()
	if _,err = o.QueryTable("taskinstance").
		Filter("userId",uid).Filter("instanceId",insid).
		Filter("taskName",name).All(&v); err != nil {
			fmt.Println(err)
			return nil,err
	}
	return
}

func GetTaskinstanceBySort(uid,ins,sort int) (v Taskinstance, err error) {
	o := orm.NewOrm()
	if err = o.QueryTable("taskinstance").
		Filter("userId",uid).
		Filter("instanceId",ins).
		Filter("taskSort",sort).One(&v); err != nil {
		fmt.Println(err)
		return
	}
	return
}

// GetAllTaskinstance retrieves all Taskinstance matches certain condition. Returns empty list if
// no records exist
func GetAllTaskinstance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Taskinstance))
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

	var l []Taskinstance
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

// UpdateTaskinstance updates Taskinstance by Id and returns error if
// the record to be updated doesn't exist
func UpdateTaskinstanceById(m *Taskinstance) (err error) {
	o := orm.NewOrm()
	v := Taskinstance{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTaskinstance deletes Taskinstance by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTaskinstance(id int) (err error) {
	o := orm.NewOrm()
	v := Taskinstance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Taskinstance{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
