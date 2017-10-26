package exec

import (
	"fmt"
)

func ExampleBashRunner() {
	br := BashRunner{}

	scanner, _ := br.Run("ls ./")

	for scanner.Scan() {
		fmt.Print(scanner.Text() + "\n")
	}
	// Output:
	// exec.go
	// exec_test.go
}
