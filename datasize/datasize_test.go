package datasize

import "fmt"

func ExampleConverter() {
	var output string
	cnv := &Converter{}

	output = cnv.HumanReadable(1)
	fmt.Println(output)
	output = cnv.HumanReadable(1024)
	fmt.Println(output)
	output = cnv.HumanReadable(1025)
	fmt.Println(output)
	output = cnv.HumanReadable(4096)
	fmt.Println(output)
	output = cnv.HumanReadable(2345246256)
	fmt.Println(output)
	output = cnv.HumanReadable(404235234234596)
	fmt.Println(output)
	// Output:
	// 1 B
	// 1024 B
	// 1.0 KB
	// 4.0 KB
	// 2.2 GB
	// 367.6 TB
}
