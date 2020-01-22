package models


/**
	Author:charlie
	Description:车
	Time:2019-1-14
*/
type Car struct {
	Id int
	Name string
}

/**
	Author:charlie
	Description:车--车牌
	Time:2019-1-14
*/
type CarToCard struct {
	Id int
	CarName string
	CarNum string
}

/**
	Author:charlie
	Description:返回的结构体
	Time:2019-1-14
*/
type Data struct {
	CarName string
	CarNum []string
}

/**
	Author:charlie
	Description:查询车辆--车牌并返回json数据
	Time:2019-1-14
*/
func SelectAll() map[string]interface{} {
	var car []Car
	var data []Data
	//首先拿到所有的车名 这是一句话可以查到结果
	//rows,_ := GetDB().Raw("SELECT * from cars LEFT JOIN car_to_cards on cars.`name`=car_to_cards.car_name").Rows()
	//你要保证结果 我分两次查
	//将查询到的车名存储到切片中
	GetDB().Raw("SELECT * from cars").Scan(&car)
	//遍历切片 查询车名对应的车牌号
	for _,val := range car {
		//查询所有的车牌
		var d Data
		GetDB().Raw("SELECT car_to_cards.car_num from cars Right " +
			"JOIN car_to_cards on cars.`name`=car_to_cards.car_name WHERE car_name='?'",val.Name).Scan(&d.CarNum)
		d.CarName = val.Name
		data = append(data, d)
	}
	resp := map[string]interface{}{"cars":data}
	return resp
}
