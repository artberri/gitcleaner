package main

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/urfave/cli"
)

// List all objects including their size, sort by size
// Based on https://stackoverflow.com/questions/10622179/how-to-find-identify-large-files-commits-in-git-history
func getObjects(path string) (GitObjects, *cli.ExitError) {
	var verifyPackErr, revListErr *cli.ExitError
	var verifyPackScanner, revListScanner *bufio.Scanner

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		verifyPackScanner, verifyPackErr = verifyPack(path)
	}()
	go func() {
		defer wg.Done()

		revListScanner, revListErr = revList(path)
	}()
	wg.Wait()

	if verifyPackErr != nil {
		return nil, cli.NewExitError(verifyPackErr.Error(), 1)
	}
	if revListErr != nil {
		return nil, cli.NewExitError(revListErr.Error(), 1)
	}

	objects, err := parseVerifyPack(verifyPackScanner)
	if err != nil {
		return nil, cli.NewExitError(err.Error(), 1)
	}

	return parseRevList(revListScanner, objects)
}

// Create a map of objects with filnames
func parseVerifyPack(scanner *bufio.Scanner) (map[string]GitObject, *cli.ExitError) {
	objects := map[string]GitObject{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) > 5 {
			size, err := strconv.ParseUint(line[4], 10, 64)
			if err != nil {
				return nil, cli.NewExitError(err.Error(), 1)
			}
			compressedSize, err := strconv.ParseUint(line[5], 10, 64)
			if err != nil {
				return nil, cli.NewExitError(err.Error(), 1)
			}
			sha := line[0]
			objects[sha] = GitObject{sha, size, compressedSize, ""}
		}
	}

	return objects, nil
}

// Complete a map of objects with filenames
func parseRevList(scanner *bufio.Scanner, objects map[string]GitObject) (GitObjects, *cli.ExitError) {
	gitObjects := GitObjects{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) > 1 {
			sha, name := line[0], line[1]
			if gitObject, exists := objects[sha]; exists && name != "" {
				gitObject.path = name
				gitObjects = append(gitObjects, gitObject)
			}
		}
	}

	sort.Sort(gitObjects)

	return gitObjects, nil
}

func groupObjectsByFile(oldObjects GitObjects) GitObjects {
	gitObjects := GitObjects{}
	objectMap := map[string]GitObject{}

	for _, oldObject := range oldObjects {
		if object, exists := objectMap[oldObject.path]; exists {
			object.size += oldObject.size
			objectMap[oldObject.path] = object
		} else {
			oldObject.sha = ""
			objectMap[oldObject.path] = oldObject
		}
	}

	for _, object := range objectMap {
		gitObjects = append(gitObjects, object)
	}

	sort.Sort(gitObjects)

	return gitObjects
}
