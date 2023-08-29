package db_operation

import "testing"

func TestMain(m *testing.M) {
	DBinit()
	m.Run()
}
