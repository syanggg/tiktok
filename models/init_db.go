package models

import(
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/util"
)

var DB *gorm.DB

func InitDB(){
	var err error
	DB ,err = gorm.Open(mysql.Open(DBConnectionString()),&gorm.Config{
		SkipDefaultTransaction: true,   //禁用事务操作。不禁用，写入操作会默认事务操作，处理不好可能会出现写入失效的情况
	})
	if err != nil{
		panic("failed to connect database")
	}

	//迁移
	err = DB.AutoMigrate(&UserInfo{},&UserLogin{},&Video{},&Comment{})

	if err != nil{
		panic(err)
	}

}

// DBConnectionString 获取MySQL数据库链接字符串
func DBConnectionString() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",util.Info.DB.Username,
		util.Info.DB.Password,util.Info.DB.Host,util.Info.DB.Port,util.Info.DB.Database,util.Info.DB.Charset,
		util.Info.DB.ParseTime,util.Info.DB.Loc)
}