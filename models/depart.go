package models

import (
	"fmt"
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
func GetAllDepart() (ml []Depart) {
	o := orm.NewOrm()
	o.QueryTable("depart").All(&ml)
	return ml
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
