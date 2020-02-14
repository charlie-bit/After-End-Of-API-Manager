package models

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	"github.com/jinzhu/gorm"
	"net"

	//gorm自己包装的一些驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

/**
	Author:charlie
	Description:数据库连接
	Time:2019-12-31
 */
func Init()  {
	/**
		读取配置文件内容
		用户名 密码 tcp地址 数据库类型 数据库编码
	 */
	conn,err := gorm.Open("mysql",...)/ApiProject?charset=utf8&parseTime=true&loc=Local")//涉及到内部信息 格式可参照gorm官方文档
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	//调试环境--注册表
	db = conn
	db.Debug().AutoMigrate(&User{},&ApiRecord{},&ApiUrlLimit{},&Role{},&Manager{})
}

func GetDB() *gorm.DB {
	return db
}

/**
	Author:charlie
	Description:创造一个缓存空间
	Time:2019-1-2
*/
func GetCache() cache.Cache {
	bm, _ := cache.NewCache("memory", `{"interval":60}`)
	return bm
}

/**
	Author:charlie
	Description:获得我发布到服务器端，客户端访问的本地ip地址
	Time:2019-1-1
*/
func ClientIP() (ip string) {
	//获取所有网卡
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		ip = ""
	}
	for _, value := range addrs{
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil{
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}
