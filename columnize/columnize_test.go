package columnize

import (
	"fmt"
)

func ExampleColumnizer() {
	clm := &Columnizer{}

	output := clm.Columnize([]string{
		"column1 | column2 | column 3",
		"short | very long very long very long very long very long very long | whatever",
		"hey | ho | let's go",
	})

	fmt.Println(output)
	// Output:
	// column1  column2                                                      column 3
	// short    very long very long very long very long very long very long  whatever
	// hey      ho                                                           let's go
}
