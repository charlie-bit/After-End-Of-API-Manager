package models

import (
	"apiproject_new/utils"
	"github.com/dgrijalva/jwt-go"
	"regexp"
)

/**
	Author:charlie
	Description:用户表
	Time:2019-12-31
 */
type User struct {
	Id                 int     		`gorm:"primary_key;column:user_id";json:"id"`								//主键
	Ip                 string		`gorm:"varchar(400);column:user_ip";json:"user_ip"`							//ip地址
	UserName           string  		`gorm:"varchar(20);unique_index;column:user_name";json:"username"`			//用户名
	Password           string  		`gorm:"varchar(20);column:user_password";json:"password"`					//密码
	Token              string  		`gorm:"varchar(400);column:user_token";json:"token"`						//token标识符
	EffectiveTime      string		`gorm:"varchar(400);column:effective_time";json:"effective_time"`			//token有效时间
	ExpirationTime     string		`gorm:"varchar(400);column:expiration_time";json:"expiration_time"`			//token失效时间
	UseState		   int     		`gorm:"column:use_state";json:"use_state"` 									//使用状态
	Role               string       `gorm:"varchar(400);column:use_role";json:"use_role"`						//用户角色
}
/**
	Author:charlie
	Description:用户表
	Time:2019-12-31
 */
type Token struct {
	UserId int
	jwt.StandardClaims
}
/**
	Author:charlie
	Description:验证用户，密码
				用户名：2个英文字母以上
				密码：6位数的密码
	Time:2019-12-31
 */
func (user *User)Validate() (map[string]interface{}, bool) {

	if match,_ := regexp.MatchString("^[a-zA-Z]{2,10}$",user.UserName); !match {
		return utils.Message("用户名格式不正确"),false
	}

	if match, _ := regexp.MatchString("^\\d{6}$",user.Password); !match {
		return utils.Message("密码格式不正确"),false
	}

	 return utils.Message("该用户名，密码均可用"), true
}

/**
	Author:charlie
	Description:新增用户
				用户名：2个英文字母以上
				密码：6位数的密码
	Time:2019-1-10
*/
func (user *User)CreatUser()  {
	GetDB().Create(user)
}
/**
	Author:charlie
	Description:根据用户名查询对象
	Time:2019-1-10
*/
func (user *User)SelectUserByName(name string) User {
	var userTest User
	GetDB().Table("users").Where("user_name = ?",user.UserName).First(&userTest)
	return userTest
}
/**
	Author:charlie
	Description:根据用户名修改用户所有字段
	Time:2019-1-10
*/
func (u *User)UpdateUser(user User)  {
	GetDB().Save(&user)
}
/**
	Author:charlie
	Description:根据用户名和角色查询对象
	Time:2019-1-14
*/
func (u *User)SelectUserByNameAndRole(name string,role string) User {
	var userTest User
	GetDB().Where(&User{UserName:name,Role:role}).First(&userTest)
	return userTest
}
/**
	Author:charlie
	Description:查询user表
	Time:2019-1-15
*/
func (u *User)SelectUserAll(page,size int,table string) interface{} {
	var tables interface{}
	if table == "users" {
		var users []User
		GetDB().Raw("select * from "+table+" limit ?,?",page*size-size,size).Scan(&users)
		tables = users
	}
	if table == "api_records" {
		var records []ApiRecord
		GetDB().Raw("select * from "+table+" limit ?,?",page*size-size,size).Scan(&records)
		tables = records
	}
	if table == "api_url_limits"{
		var limit []ApiUrlLimit
		GetDB().Raw("select * from "+table+" limit ?,?",page*size-size,size).Scan(&limit)
		tables = limit
	}
	return tables
}
/**
	Author:charlie
	Description:查询user总页数
	Time:2019-1-16
*/
func (user *User)SelectUserTotal(size int,table string) int {
	var count,total int
	GetDB().Table(table).Count(&count)
	if count%size == 0 {
		total = count/size
	}else{
		total = count/size + 1
	}
	return total
}
/**
	Author:charlie
	Description:更新多个字段
	Time:2019-1-16
*/
func (user *User)UpdateUserByMap(m map[string]interface{})  {
	GetDB().Model(&user).Updates(m)
}
/**
	Author:charlie
	Description:根据id查询对象
	Time:2019-1-16
*/
func (user *User)SelectUserById(id int) User {
	var userTest User
	GetDB().Table("users").Where("user_id = ?",id).First(&userTest)
	return userTest
}
/**
	Author:charlie
	Description:根据id删除对象
	Time:2019-1-16
*/
func (user *User)DeleteUserById(id int)  {
	user.Id = id
	GetDB().Delete(&user)
}
