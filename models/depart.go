package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Depart struct {
	Id             int       `orm:"column(departId);auto" description:"閮ㄩ棬Id"`
	DepartName     string    `orm:"column(departName);size(100);null" description:"閮ㄩ棬鍚嶇О"`
	SuperiorId     int64     `orm:"column(superiorId);null" description:"涓婄骇閮ㄩ棬Id"`
	DepartTypeId   int64     `orm:"column(departTypeId);null" description:"閮ㄩ棬绫诲瀷Id"`
	CreateTime     time.Time `orm:"column(createTime);type(datetime);null" description:"鍒涘缓鏃堕棿"`
	ModifiedTime   time.Time `orm:"column(modifiedTime);type(datetime);null" description:"淇敼鏃堕棿"`
	CreateUserId   int64     `orm:"column(createUserId);null" description:"鍒涘缓鑰匢d"`
	ModifiedUserId int64     `orm:"column(modifiedUserId);null" description:"淇敼鑰匢d"`
}

func (t *Depart) TableName() string {
	return "depart"
}

func init() {
	orm.RegisterModel(new(Depart))
}

// AddDepart insert a new Depart into database and returns
// last inserted Id on success.
func AddDepart(m *Depart) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDepartById retrieves Depart by Id. Returns error if
// Id doesn't exist
func GetDepartById(id int) (v *Depart, err error) {
	o := orm.NewOrm()
	v = &Depart{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDepart retrieves all Depart matches certain condition. Returns empty list if
// no records exist
func GetAllDepart(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Depart))
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

	var l []Depart
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

// UpdateDepart updates Depart by Id and returns error if
// the record to be updated doesn't exist
func UpdateDepartById(m *Depart) (err error) {
	o := orm.NewOrm()
	v := Depart{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDepart deletes Depart by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDepart(id int) (err error) {
	o := orm.NewOrm()
	v := Depart{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Depart{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
