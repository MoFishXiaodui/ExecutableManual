package db_operation

import "testing"

// 无效删除
func TestDelete(t *testing.T) {
	stu := &Stu{Name: "赤兔"}
	db.Delete(stu)
}

// 无效删除
func TestDelete2(t *testing.T) {
	stu := &Stu{Name: "赤兔"}
	db.Unscoped().Delete(stu)
}

// 有效删除
func TestDelete3(t *testing.T) {
	stu := &Stu{}
	db.Where("stu_name = ?", "赤兔").Delete(stu)
}

// 有效删除一条数据
func TestDelete4(t *testing.T) {
	stu := &Stu{}
	db.First(stu, "stu_name = ?", "赤兔").Delete(stu)
	// First会找到第一条符合条件的数据，并在stu结构体中指定主键，所以在执行Delete操作时，可以确认删除的具体行
}
