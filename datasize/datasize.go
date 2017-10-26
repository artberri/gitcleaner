package datasize

import "github.com/c2h5oh/datasize"

// Converter manages CLI
type Converter struct{}

// HumanReadable converts datasize into Human readable format
func (c *Converter) HumanReadable(size uint64) string {
	return datasize.ByteSize(size).HumanReadable()
}
