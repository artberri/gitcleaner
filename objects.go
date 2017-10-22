package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

// List all objects including their size, sort by size
// Based on http://stubbisms.wordpress.com/2009/07/10/git-script-to-show-largest-pack-objects-and-trim-your-waist-line/
func getObjects(path string) (GitObjects, *cli.ExitError) {
	objects := map[string]GitObject{}
	scanner, err := verifyPack(path)
	if err != nil {
		return nil, cli.NewExitError(err.Error, 1)
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) > 5 {
			size, err := strconv.ParseUint(line[4], 10, 64)
			if err != nil {
				return nil, cli.NewExitError(err.Error, 1)
			}
			compressedSize, err := strconv.ParseUint(line[5], 10, 64)
			if err != nil {
				return nil, cli.NewExitError(err.Error, 1)
			}
			sha := line[0]
			objects[sha] = GitObject{sha, size, compressedSize, ""}
		}
	}

	return completeObjectsWithNames(path, objects)
}

// Complete a map of objects with filnames
func completeObjectsWithNames(path string, objects map[string]GitObject) (GitObjects, *cli.ExitError) {
	gitObjects := GitObjects{}
	scanner, err := revList(path)
	if err != nil {
		return nil, cli.NewExitError(err.Error, 1)
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) > 1 {
			sha, name := line[0], line[1]
			_, exists := objects[sha]
			if exists && name != "" {
				gitObject := objects[sha]
				gitObject.path = name
				gitObjects = append(gitObjects, gitObject)
			}
		}
	}

	sort.Sort(gitObjects)

	return gitObjects, nil
}
