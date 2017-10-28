package domain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestExecThrowsErrorIfGetObjectsThrowsError(t *testing.T) {
	mockObjectManager := &MockObjectManager{}
	mockObjectManager.GetFn = func(path string) ([]GitObject, error) {
		return nil, errors.New("error get")
	}

	listCommand := ListCommand{}
	listCommand.ObjectManager = mockObjectManager
	err := listCommand.Exec("/path/repo", 10, false, false)

	if err == nil || err.Error() != "error get" {
		t.Fatal("Not error throw")
	}
}

func ExampleListCommand() {
	// Example1
	mockObjectManager := &MockObjectManager{}
	mockObjectManager.GetFn = func(path string) ([]GitObject, error) {
		return []GitObject{{
			sha:            "aa11",
			size:           1234,
			compressedSize: 700,
			path:           "./file1.txt",
		}, {
			sha:            "aa12",
			size:           2234,
			compressedSize: 720,
			path:           "./file2.txt",
		}, {
			sha:            "aa13",
			size:           3234,
			compressedSize: 730,
			path:           "./file3.txt",
		}, {
			sha:            "aa14",
			size:           4234,
			compressedSize: 740,
			path:           "./file4.txt",
		}, {
			sha:            "aa15",
			size:           51234,
			compressedSize: 7550,
			path:           "./file5.txt",
		}}, nil
	}
	mockObjectManager.GroupObjectsByFileFn = func(oldObjects []GitObject) []GitObject {
		return []GitObject{{
			size:           1234,
			compressedSize: 700,
			path:           "./file21.txt",
		}, {
			size:           2234,
			compressedSize: 720,
			path:           "./file22.txt",
		}, {
			size:           3234,
			compressedSize: 730,
			path:           "./file23.txt",
		}}
	}

	mockConverter := &MockConverter{}
	mockConverter.HumanReadableFn = func(size uint64) string {
		return strconv.FormatUint(size, 10) + "test"
	}

	var toColumnizeRows []string
	mockColumnizer := &MockColumnizer{}
	mockColumnizer.ColumnizeFn = func(rows []string) string {
		toColumnizeRows = rows
		return "exit"
	}

	listCommand := ListCommand{}
	listCommand.ObjectManager = mockObjectManager
	listCommand.Columnizer = mockColumnizer
	listCommand.Converter = mockConverter
	listCommand.Exec("/path/repo", 10, false, false)
	fmt.Println(mockObjectManager.GroupObjectsByFileInvoked)
	fmt.Println(mockConverter.HumanReadableInvoked)
	for _, row := range toColumnizeRows {
		fmt.Println(row)
	}

	listCommand = ListCommand{}
	listCommand.ObjectManager = mockObjectManager
	listCommand.Columnizer = mockColumnizer
	listCommand.Converter = mockConverter
	listCommand.Exec("/path/repo", 2, true, true)
	fmt.Println(mockObjectManager.GroupObjectsByFileInvoked)
	fmt.Println(mockConverter.HumanReadableInvoked)
	for _, row := range toColumnizeRows {
		fmt.Println(strings.TrimSpace(row))
	}

	// Output:
	// exit
	// false
	// false
	// 1234 | ./file1.txt | aa11
	// 2234 | ./file2.txt | aa12
	// 3234 | ./file3.txt | aa13
	// 4234 | ./file4.txt | aa14
	// 51234 | ./file5.txt | aa15
	// exit
	// true
	// true
	// 1234test | ./file21.txt |
	// 2234test | ./file22.txt |
}
