package db_operation

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"model/config"
	"testing"
)

type Stu struct {
	Name string `gorm:"column:stu_name"`
	Age  uint8  `gorm:"column:age;default:18"`
}

func (s Stu) TableName() string {
	return "students"
}

func TestCreate(t *testing.T) {
	var addr, user, pwd, dbName string = config.GetMySQLConfig()
	dsn := user + ":" + pwd + "@tcp(" + addr + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&Stu{})

	stu := &Stu{Name: "张三"}
	res := db.Create(stu)
	if res.Error != nil {
		t.Errorf("sth wrong, err: %v", res.Error)
	}
}

func TestCreate2(t *testing.T) {
	var addr, user, pwd, dbName string = config.GetMySQLConfig()
	dsn := user + ":" + pwd + "@tcp(" + addr + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&Stu{})

	stu := &Stu{Name: "李四", Age: 23}
	res := db.Create(stu)
	if res.Error != nil {
		t.Errorf("sth wrong, err: %v", res.Error)
	}
}
