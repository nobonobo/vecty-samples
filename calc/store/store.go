package store

import (
	"strconv"
)

// State ...
var State = &StateDef{
	NowInput: "0",
}

// StateDef ...
type StateDef struct {
	NowInput  string
	LastInput string
	Operator  rune
	Modified  bool
}

// Now ...
func (s *StateDef) Now() (float64, error) {
	return strconv.ParseFloat(s.NowInput, 64)
}

// Last ...
func (s *StateDef) Last() (float64, error) {
	return strconv.ParseFloat(s.LastInput, 64)
}

// IsZero ...
func (s *StateDef) IsZero() bool {
	return s.NowInput == "0" || s.NowInput == "-0"
}
