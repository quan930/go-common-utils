package ini

import "testing"

func TestIniGetKey(t *testing.T) {
	Init("./test.ini")
	name := IniGetKey("default", "Name")
	if name != "test" {
		t.Errorf("IniGetKey(%v,%v) = %v; expected %v", "default", "Name", name, "test")
	}
}
