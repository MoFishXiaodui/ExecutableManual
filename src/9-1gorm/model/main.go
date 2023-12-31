package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"model/config"
)

type Stu struct {
	Name string `gorm:"column:stu_name"`
	Age  uint8  `gorm:"column:age;default:18"`
}

func (s Stu) TableName() string {
	return "students"
}
func main() {
	var addr, user, pwd, dbName string = config.GetMySQLConfig()
	dsn := user + ":" + pwd + "@tcp(" + addr + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&Stu{})
}
