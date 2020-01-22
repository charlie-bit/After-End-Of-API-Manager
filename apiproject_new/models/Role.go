package models
/**
	Author:charlie
	Description:用户角色表
	Time:2019-1-17
*/
type Role struct {
	Id 		int 			`gorm:"primary_key;column:role_id";json:"id"`
	Role 	string 			`gorm:"varchar(400);column:use_role";json:"use_role"`
	Leval   int				`gorm:"column:use_leval";json:"use_leval"`
}

/**
	Author:charlie
	Description:用户表
	Time:2019-1-17
*/
func (r *Role)SelectRoleAll() []Role {
	var roles []Role
	GetDB().Find(&roles)
	return roles
}
