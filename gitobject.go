package main

// GitObject Git object with properties splitted
type GitObject struct {
	sha            string
	size           uint64
	compressedSize uint64
	path           string
}

// GitObjects Slice of GitObject
type GitObjects []GitObject

// Len implementation for sorting
func (objects GitObjects) Len() int {
	return len(objects)
}

// Swap implementation for sorting
func (objects GitObjects) Swap(i, j int) {
	objects[i], objects[j] = objects[j], objects[i]
}

// Less implementation for sorting
func (objects GitObjects) Less(i, j int) bool {
	return objects[i].size > objects[j].size
}
