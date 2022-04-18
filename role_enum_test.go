package modelrole

import (
	"testing"
)

func TestEnumLevel(t *testing.T) {
	level := RoleAdmin.Level()
	expected := 2

	if level != expected {
		t.Errorf("should have %d, got %d", expected, level)
	}
}

func TestEnumValue(t *testing.T) {
	value, _ := RoleAdmin.Value()
	expected := "admin"

	if value != expected {
		t.Errorf("should have %s, got %s", expected, value)
	}
}
