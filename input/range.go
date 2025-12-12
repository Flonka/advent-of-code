package input

import (
	"errors"
	"strconv"
	"strings"
)

// Range is used for handling input integers given in a range separated with a dash
type Range struct {
	Start int
	End   int
}

// NewRangeFromString creates from problem input
// format start-end
func NewRangeFromString(input string) (Range, error) {
	s := strings.Split(input, "-")

	r := Range{}
	i1, err := strconv.Atoi(s[0])
	if err != nil {
		return r, errors.Join(err)
	}
	i2, err := strconv.Atoi(s[1])
	if err != nil {
		return r, errors.Join(err)
	}

	r.Start = i1
	r.End = i2

	return r, nil
}
