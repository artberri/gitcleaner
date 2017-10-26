package columnize

import "github.com/ryanuber/columnize"

// Columnizer create columnes in plain text
type Columnizer struct{}

// Columnize create columnes in plain text
func (c *Columnizer) Columnize(rows []string) string {
	return columnize.SimpleFormat(rows)
}
