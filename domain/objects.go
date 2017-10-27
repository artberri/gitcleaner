package domain

import (
	"bufio"
	"errors"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// GitObject Git object with properties splitted
type GitObject struct {
	sha            string
	size           uint64
	compressedSize uint64
	path           string
}

// BySizeDesc implements sort.Interface for []GitObject based on size
type BySizeDesc []GitObject

// Len implementation for sorting
func (objects BySizeDesc) Len() int {
	return len(objects)
}

// Swap implementation for sorting
func (objects BySizeDesc) Swap(i, j int) {
	objects[i], objects[j] = objects[j], objects[i]
}

// Less implementation for sorting
func (objects BySizeDesc) Less(i, j int) bool {
	return objects[i].size > objects[j].size
}

// GitObjectManager manages git objects
type GitObjectManager struct {
	Git RepoManager
}

// Get gets all objects including their size and other data, sorted by size
// Based on https://stackoverflow.com/questions/10622179/how-to-find-identify-large-files-commits-in-git-history
func (gom *GitObjectManager) Get(path string) ([]GitObject, error) {
	var errVerifyPath, errRevList, errPath error
	var verifyPackScanner, revListScanner *bufio.Scanner
	var objects map[string]GitObject

	path, errPath = gom.Git.EnsureRepoPath(path)
	if errPath != nil {
		return nil, errPath
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		verifyPackScanner, errVerifyPath = gom.Git.VerifyPack(path)
	}()
	go func() {
		defer wg.Done()

		revListScanner, errRevList = gom.Git.RevList(path)
	}()
	wg.Wait()

	if errVerifyPath != nil {
		return nil, errors.New(errVerifyPath.Error())
	}
	if errRevList != nil {
		return nil, errors.New(errRevList.Error())
	}

	objects, err := gom.parseVerifyPack(verifyPackScanner)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return gom.parseRevList(revListScanner, objects), nil
}

// GroupObjectsByFile groups git objects by file to calculate size
func (gom *GitObjectManager) GroupObjectsByFile(oldObjects []GitObject) []GitObject {
	gitObjects := []GitObject{}
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

	sort.Sort(BySizeDesc(gitObjects))

	return gitObjects
}

// Create a map of objects with filnames
func (gom *GitObjectManager) parseVerifyPack(scanner *bufio.Scanner) (map[string]GitObject, error) {
	objects := map[string]GitObject{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) > 5 {
			size, err := strconv.ParseUint(line[4], 10, 64)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			compressedSize, err := strconv.ParseUint(line[5], 10, 64)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			sha := line[0]
			objects[sha] = GitObject{sha, size, compressedSize, ""}
		}
	}

	return objects, nil
}

// Complete a map of objects with filenames
func (gom *GitObjectManager) parseRevList(scanner *bufio.Scanner, objects map[string]GitObject) []GitObject {
	gitObjects := []GitObject{}

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

	sort.Sort(BySizeDesc(gitObjects))

	return gitObjects
}
