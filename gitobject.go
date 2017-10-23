package main

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
