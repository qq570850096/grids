package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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
func GetAllProcess(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Process))
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

	var l []Process
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
