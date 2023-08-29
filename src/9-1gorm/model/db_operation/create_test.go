package db_operation

import (
	"testing"
)

func TestCreate(t *testing.T) {
	stu := &Stu{Name: "张三"}
	res := db.Create(stu)
	if res.Error != nil {
		t.Errorf("sth wrong, err: %v", res.Error)
	}
}

func TestCreate2(t *testing.T) {
	stu := &Stu{Name: "克罗地亚", Age: 25}
	res := db.Create(stu)
	if res.Error != nil {
		t.Errorf("sth wrong, err: %v", res.Error)
	}
}
