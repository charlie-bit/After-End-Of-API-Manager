package models
/**
	Author:charlie
	Description:api权限设置
	Time:2019-12-31
*/
type ApiUrlLimit struct {
	Id 				int			`gorm:"primary_key;column:api_id";json:"api_id"`
	ApiUrl  		string		`gorm:"varchar(400);column:api_url";json:"api_url"`
	ApiDescription  string   	`gorm:"varchar(400);column:api_description";json:"api_description"`
	Role 	 		string		`gorm:"varchar(400);column:use_role";json:"use_role"`
}

/**
	Author:charlie
	Description:根据id删除对象
	Time:2019-1-16
*/
func (url *ApiUrlLimit)DeleteApiUrlLimitsById(id int)  {
	url.Id = id
	GetDB().Delete(&url)
}

/**
	Author:charlie
	Description:根据id删除对象
	Time:2019-1-19
*/
func (aul *ApiUrlLimit)AddApiLimit()  {
	GetDB().Create(aul)
}
/**
	Author:charlie
	Description:根据api路径查是否有重复的
	Time:2019-1-19
*/
func (aul *ApiUrlLimit)SelectApiLimitByApi(api string) ApiUrlLimit {
	var aulTest ApiUrlLimit
	GetDB().Table("api_url_limits").Where("api_url = ?",api).First(&aulTest)
	return aulTest
}
/**
	Author:charlie
	Description:更改用户角色权限
	Time:2019-1-19
*/
func (aul *ApiUrlLimit)UpdateRoleLimit(role string,des string)  {
	GetDB().Model(&aul).Updates(map[string]string{"use_role":role,"api_description":des})
}
/**
	Author:charlie
	Description:根据id查找api
	Time:2019-1-19
*/
func (aul *ApiUrlLimit)SelectApiLimit(id int)  {
	GetDB().First(&aul,id)
}
