package actions

// InsertChar ...
type Insert struct {
	Char rune
}

// Invert ...
type Invert struct{}

// Clear ...
type Clear struct {
	All bool
}

// Percent ...
type Percent struct{}

// Operator ...
type Operator struct {
	Char rune
}

// Equal ...
type Equal struct{}
