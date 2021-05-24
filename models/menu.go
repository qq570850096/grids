package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	Id       int    `orm:"column(menuId);pk" description:"鑿滃崟Id"`
	ParentId int64  `orm:"column(parentId);null" description:"鐖惰彍鍗旾d"`
	MenuName string `orm:"column(menuName);size(50);null" description:"鑿滃崟鍚嶅瓧"`
	MenuType int    `orm:"column(menuType);null" description:"鑿滃崟绫诲瀷锛�0锛氱洰褰曘��1锛氳彍鍗曪級"`
	Url      string `orm:"column(url);size(255);null" description:"璇锋眰鍦板潃"`
	Sort     int    `orm:"column(sort);null" description:"鎺掑簭"`
}

type TreeOV struct {
	MenuName string `json:"title"`
	Id       int    `json:"id"`
	Checked  bool 	`json:"checked"`
	Children []TreeOV `json:"children"`
	ResourceIds []int `json:"resourceIds"`
	CreatUserId int `json:"creatUserId"`
	ModifiedUserId int `json:"modifiedUserId"`
}

func (t *Menu) TableName() string {
	return "menu"
}

func init() {
	orm.RegisterModel(new(Menu))
}

// AddMenu insert a new Menu into database and returns
// last inserted Id on success.
func AddMenu(m *Menu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMenuById retrieves Menu by Id. Returns error if
// Id doesn't exist
func GetMenuById(id int) (v *Menu, err error) {
	o := orm.NewOrm()
	v = &Menu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMenu retrieves all Menu matches certain condition. Returns empty list if
// no records exist
func GetAllMenu() (menus []Menu) {
	o:= orm.NewOrm()
	if _,err := o.QueryTable("menu").All(&menus);err != nil{
		fmt.Println(err)
		return nil
	}
	return
}

func GetAllMenubyParent(pid int) (menus []Menu) {
	o:= orm.NewOrm()
	if _,err := o.QueryTable("menu").Filter("parentId",pid).
		All(&menus);err != nil{
		fmt.Println(err)
		return nil
	}
	return
}
// UpdateMenu updates Menu by Id and returns error if
// the record to be updated doesn't exist
func UpdateMenuById(m *Menu) (err error) {
	o := orm.NewOrm()
	v := Menu{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMenu deletes Menu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMenu(id int) (err error) {
	o := orm.NewOrm()
	v := Menu{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Menu{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
