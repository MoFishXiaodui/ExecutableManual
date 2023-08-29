package db_operation

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"model/config"
	"testing"
)

func TestFirst(t *testing.T) {
	var addr, user, pwd, dbName string = config.GetMySQLConfig()
	dsn := user + ":" + pwd + "@tcp(" + addr + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&Stu{})

	stu := &Stu{}
	res := db.First(stu, "stu_name = ?", "李四") // "李四"会自动替换前面的问号，作为First查询的条件
	if res.Error != nil {
		t.Errorf("sth wrong, err: %v", res.Error)
	}
	if stu.Age != 23 {
		t.Errorf("李四应该是23岁，而拿到的数据库数据却是：%v", stu.Age)
	}
}
