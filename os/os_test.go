package os

import "testing"

func TestFileExisterDetectsFileExistenceProperly(t *testing.T) {
	fe := FileExister{}

	if exists := fe.Exists("../README.md"); !exists {
		t.Fatal("Error checking that the README.md file exists")
	}

	if exists := fe.Exists("../NOFILE"); exists {
		t.Fatal("Error checking that the README.md file exists")
	}
}
