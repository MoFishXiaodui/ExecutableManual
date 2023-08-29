package db_operation

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	stu := &Stu{}
	db.First(stu, "stu_name = ?", "克罗地亚")
	ageBefore := stu.Age
	stu.Age++
	db.Save(stu)

	newStu := &Stu{}
	db.First(newStu, "stu_name = ?", "克罗地亚")
	if ageBefore+1 != newStu.Age {
		t.Errorf("克罗地亚的年龄应该是: %v， 而实测结果是: %v\n", ageBefore+1, newStu.Age)
	}
}
