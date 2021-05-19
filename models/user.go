package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id             int       `orm:"column(userId);auto" description:"鐢ㄦ埛Id"`
	UserName       string    `orm:"column(userName);size(50);null" description:"璐﹀彿"`
	RealName       string    `orm:"column(realName);size(50);null" description:"濮撳悕"`
	Gender         string    `orm:"column(gender);size(10);null" description:"鎬у埆"`
	Password       string    `orm:"column(password);size(20);null" description:"瀵嗙爜"`
	Phone          string    `orm:"column(phone);size(20);null" description:"鎵嬫満鍙风爜"`
	DepartId       int     `orm:"column(departId);null" description:"鎵�灞為儴闂↖d"`
	JobId          int     `orm:"column(jobId);null" description:"鎵�灞炶亴鍔d"`
	CreateTime     time.Time `orm:"column(createTime);type(datetime);null" description:"鍒涘缓鏃堕棿"`
	ModifiedTime   time.Time `orm:"column(modifiedTime);type(datetime);null" description:"淇敼鏃堕棿"`
	CreateUserId   int     `orm:"column(createUserId);null" description:"鍒涘缓鑰匢d"`
	ModifiedUserId int     `orm:"column(modifiedUserId);null" description:"淇敼鑰匢d"`
}

type UserRJD struct {
	Id             int       `orm:"column(userId);auto" description:"鐢ㄦ埛Id"`
	UserName       string    `orm:"column(userName);size(50);null" description:"璐﹀彿"`
	RealName       string    `orm:"column(realName);size(50);null" description:"濮撳悕"`
	Gender         string    `orm:"column(gender);size(10);null" description:"鎬у埆"`
	Phone          string    `orm:"column(phone);size(20);null" description:"鎵嬫満鍙风爜"`
	CreateTime     time.Time `orm:"column(createTime);type(datetime);null" description:"鍒涘缓鏃堕棿"`
	ModifiedTime   time.Time `orm:"column(modifiedTime);type(datetime);null" description:"淇敼鏃堕棿"`
	RoleName       string    `orm:"column(roleName);size(50);null" description:"瑙掕壊鍚嶇О"`
	JobName        string    `orm:"column(jobName);size(50);null" description:"鑱屽姟鍚嶇О"`
	Jobid             int       `orm:"column(departId);auto" description:"閮ㄩ棬Id"`
	Departid             int       `orm:"column(jobId);auto" description:"鑱屽姟Id"`
	DepartName     string    `orm:"column(departName);size(100);null" description:"閮ㄩ棬鍚嶇О"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUserByLevel(dep,job int) (v []User, err error) {
	o := orm.NewOrm()
	_,err = o.QueryTable("user").Filter("departId",dep).Filter("jobId",job).All(&v)
	if err != nil {
		return nil, err
	}
	return
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser() (user []User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	_, err = qs.All(&user)
	if err != nil {
		return nil, err
	}
	return user, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func UserList() (userlist []UserRJD, err error) {
	o := orm.NewOrm()
	if _,err = o.Raw("select r.*,u.* ,d.`departName`,j.`jobName` " +
				"from `role` r,user_role ur,`user` u,depart d,job j " +
				"where  u.`userId` = ur.`userId` and r.`roleId` = ur.`roleId` " +
				"and u.`jobId` = j.`jobId` " +
				"and u.`departId` = d.`departId`").QueryRows(&userlist); err != nil {
		return nil ,err
	}
	return userlist,err
}

func SearchUser(name string) (userlist []UserRJD,err error) {
	o := orm.NewOrm()
	name = "%" + name + "%"
	if _,err = o.Raw("select r.*,u.* ,d.`departName`,j.`jobName` " +
		"from `role` r,user_role ur,`user` u,depart d,job j " +
		"where  u.`userId` = ur.`userId` and r.`roleId` = ur.`roleId` " +
		"and u.`jobId` = j.`jobId` and u.`userName` like ? " +
		"and u.`departId` = d.`departId`", name).QueryRows(&userlist); err != nil {
		return nil, err
		fmt.Println(err)
	}
	return userlist,nil
}

