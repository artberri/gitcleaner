package domain

import (
	"fmt"
	"strconv"
)

// ListCommand implements Command and lists heavier file objects in the repository history
type ListCommand struct {
	Converter     Converter
	Columnizer    Columnizer
	ObjectManager ObjectManager
}

// Exec executes the command
func (lc *ListCommand) Exec(path string, max int, humanReadable bool, unique bool) error {
	objects, err := lc.ObjectManager.Get(path)
	if err != nil {
		return err
	}

	output := []string{}
	count := len(objects)

	if unique {
		objects = lc.ObjectManager.GroupObjectsByFile(objects)
	}

	if max == 0 || count < max {
		max = count
	}
	for i := 0; i < max; i++ {
		var size string
		if humanReadable {
			size = lc.Converter.HumanReadable(objects[i].size)
		} else {
			size = strconv.FormatUint(objects[i].size, 10)
		}
		output = append(output, size+" | "+objects[i].path+" | "+objects[i].sha)
	}

	result := lc.Columnizer.Columnize(output)
	fmt.Println(result)

	return nil
}
