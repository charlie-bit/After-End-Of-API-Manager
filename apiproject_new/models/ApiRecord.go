package models

/**
	Author:charlie
	Description:用户表
	Time:2019-1-2
*/
type ApiRecord struct {
	Id 				int 		`gorm:"primary_key;column:record_id";json:"record_id"`
	Ip 				string 		`gorm:"varchar(400);column:record_ip";json:"record_ip"`
	UserName 		string 		`gorm:"varchar(400);column:user_name";json:"user_name"`
	ApiEndPoint		string 		`gorm:"varchar(400);column:api_end_point";json:"api_end_point"`
	State      	 	int  		`gorm:"column:record_state";json:"record_state"`
	Message     	string 		`gorm:"varchar(400);column:record_message";json:"record_message"`
	StartTime		string 		`gorm:"varchar(400);column:start_time";json:"start_time"`
	EndTime			string 		`gorm:"varchar(400);column:end_time";json:"end_time"`
	ResponseTime  	string 		`gorm:"varchar(400);column:response_time";json:"response_time"`
}

/**
	Author:charlie
	Description:新增api使用记录
	Time:2019-1-10
*/
func (record *ApiRecord)CreateNewApiRecord()  {
	GetDB().Create(record)
}