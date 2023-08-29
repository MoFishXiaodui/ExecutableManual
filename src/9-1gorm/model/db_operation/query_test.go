package db_operation

import (
	"testing"
)

func TestFirst(t *testing.T) {
	stu := &Stu{}
	res := db.First(stu, "stu_name = ?", "李四") // "李四"会自动替换前面的问号，作为First查询的条件
	if res.Error != nil {
		t.Errorf("sth wrong, err: %v", res.Error)
	}
	if stu.Age != 24 {
		t.Errorf("李四应该是23岁，而拿到的数据库数据却是：%v", stu.Age)
	}
}
