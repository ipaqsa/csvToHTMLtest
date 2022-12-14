package formatParser

import "golang.org/x/text/encoding"

type Parser struct {
	File            string
	Separator       rune
	SkipFirstLine   bool
	SkipEmptyValues bool
	PRNReader       func(raw string) (line []string, err error)
	Decoder         *encoding.Decoder
}
