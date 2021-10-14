package model

import "testing"

func TestQuery(t *testing.T) {
	employee := Employee{Id: "23010119751201312X"}
	err := employee.Info()
	if err != nil {
		t.Error(err)
	}
	t.Log(employee)
}
