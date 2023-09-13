package genum

import (
	"github.com/drshriveer/gtools/math"
)

// EnumLike is a generic type for something that looks like an enum.
type EnumLike math.Integer

// Enum is the base interface all generated enums implement.
type Enum interface {
	// IsValid returns true if the enum is valid.
	IsValid() bool

	// StringValues returns a list of values as strings.
	StringValues() []string

	// returns the string value of an enum.
	String() string

	// IsEnum does nothing but help define the interface.
	IsEnum()
}

// TypedEnum is extended, generic interface that enums extend.
// I'm not actually sure where this will be useful...
type TypedEnum[T EnumLike] interface {
	Enum

	// Values returns all valid values of an enum.
	Values() []T

	// ParseString converts text into a type if valid.
	// returns true if the enum is valid, and false otherwise.
	ParseString(text string) (T, error)
}
