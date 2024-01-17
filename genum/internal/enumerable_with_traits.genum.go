// Code generated by genum DO NOT EDIT.
package internal

import (
	"encoding/json"
	"fmt"
	reflect "reflect"
	"slices"
	"strings"
	stupidTime "time"

	"github.com/drshriveer/gtools/genum"
	"gopkg.in/yaml.v3"
)

var _EnumerableWithTraitsValues = []EnumerableWithTraits{
	E1,
	E2,
	E3,
}

// Timeout returns the enum's associated trait of the same name.
// If no trait exists for the enumeration a default value will be returned.
func (e EnumerableWithTraits) Timeout() stupidTime.Duration {
	switch e {
	case E1:
		return _Timeout
	case E2:
		return 1 * stupidTime.Minute
	case E3:
		return 2 * stupidTime.Minute
	}

	return *new(stupidTime.Duration)
}

// Trait returns the enum's associated trait of the same name.
// If no trait exists for the enumeration a default value will be returned.
func (e EnumerableWithTraits) Trait() string {
	switch e {
	case E1:
		return _Trait
	case E2:
		return "trait 2"
	case E3:
		return "trait 3"
	}

	return *new(string)
}

// TypedStringTrait returns the enum's associated trait of the same name.
// If no trait exists for the enumeration a default value will be returned.
func (e EnumerableWithTraits) TypedStringTrait() OtherType {
	switch e {
	case E1:
		return _TypedStringTrait
	case E2:
		return OtherType("OtherType2")
	case E3:
		return OtherType("OtherType3")
	}

	return *new(OtherType)
}

// IsValid returns true if the enum value is, in fact, valid.
func (e EnumerableWithTraits) IsValid() bool {
	for _, v := range _EnumerableWithTraitsValues {
		if v == e {
			return true
		}
	}
	return false
}

// Values returns a list of all potential values of this enum.
func (EnumerableWithTraits) Values() []EnumerableWithTraits {
	return slices.Clone(_EnumerableWithTraitsValues)
}

// StringValues returns a list of all potential values of this enum as strings.
// Note: This does not return duplicates.
func (EnumerableWithTraits) StringValues() []string {
	return []string{
		"E1",
		"E2",
		"E3",
	}
}

// String returns a string representation of this enum.
// Note: in the case of duplicate values only the first alphabetical definition will be choosen.
func (e EnumerableWithTraits) String() string {
	switch e {
	case E1:
		return "E1"
	case E2:
		return "E2"
	case E3:
		return "E3"
	default:
		return fmt.Sprintf("UndefinedEnumerableWithTraits:%d", e)
	}
}

// ParseString will return a value as defined in string form.
func (e EnumerableWithTraits) ParseString(text string) (EnumerableWithTraits, error) {
	return ParseEnumerableWithTraits(text)
}

// ParseEnumerableWithTraits will attempt to parse the value of a EnumerableWithTraits from either its string form
// or any value of a trait flagged with the --parsableByTrait flag
func ParseEnumerableWithTraits(input any) (EnumerableWithTraits, error) {
	switch input {
	case "E1":
		return E1, nil
	case "E2":
		return E2, nil
	case "E3":
		return E3, nil
	default:
		if text, ok := input.(string); ok {
			switch strings.ToLower(text) {
			case "e1":
				return E1, nil
			case "e2":
				return E2, nil
			case "e3":
				return E3, nil
			}
		}
		return 0, fmt.Errorf("`%+v` could not be parsed to enum of type EnumerableWithTraits", input)
	}
}

// ParseGeneric calls TypedEnum.Parse but returns the result
// in the generic genum.Enum interface. Which is useful when you are only able to work with
// the un-typed interface.
func (e EnumerableWithTraits) ParseGeneric(input any) (genum.Enum, error) {
	return ParseEnumerableWithTraits(input)
}

// MarshalJSON implements the json.Marshaler interface for EnumerableWithTraits.
func (e EnumerableWithTraits) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for EnumerableWithTraits.
func (e *EnumerableWithTraits) UnmarshalJSON(data []byte) error {
	// We always support strings.
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		var err error
		*e, err = ParseEnumerableWithTraits(s)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("unable to unmarshal EnumerableWithTraits from `%v`", data)
}

// MarshalText implements the encoding.TextMarshaler interface for EnumerableWithTraits.
func (e EnumerableWithTraits) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for EnumerableWithTraits.
func (e *EnumerableWithTraits) UnmarshalText(text []byte) error {
	s := string(text)
	var err error
	*e, err = ParseEnumerableWithTraits(s)
	if err == nil {
		return nil
	}

	return fmt.Errorf("unable to unmarshal EnumerableWithTraits from `%s`", s)
}

// MarshalYAML implements a YAML Marshaler for EnumerableWithTraits.
func (e EnumerableWithTraits) MarshalYAML() (any, error) {
	return e.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for EnumerableWithTraits.
func (e *EnumerableWithTraits) UnmarshalYAML(value *yaml.Node) error {
	var err error

	// first try and parse as a string
	*e, err = ParseEnumerableWithTraits(value.Value)
	if err == nil {
		return nil
	}

	// then try and parse for any string-like traits

	return fmt.Errorf("unable to unmarshal EnumerableWithTraits from yaml `%s`", value.Value)
}

// IsEnum implements an empty function required to implement Enum.
func (EnumerableWithTraits) IsEnum() {}

var _CreaturesValues = []Creatures{
	NotCreature,
	Cat,
	Dog,
	Ant,
	Spider,
	Human,
	SeaAnemone,
}

// IsCreatureMammal returns the enum's associated trait of the same name.
// If no trait exists for the enumeration a default value will be returned.
func (e Creatures) IsCreatureMammal() bool {
	switch e {
	case NotCreature:
		return _IsCreatureMammal
	case Cat:
		return true
	case Dog:
		return true
	case Ant:
		return false
	case Spider:
		return false
	case Human:
		return true
	}

	return *new(bool)
}

// NumCreatureLegs returns the enum's associated trait of the same name.
// If no trait exists for the enumeration a default value will be returned.
func (e Creatures) NumCreatureLegs() int {
	switch e {
	case NotCreature:
		return _NumCreatureLegs
	case Cat:
		return CatLegs
	case Dog:
		return DogLegs
	case Ant:
		return AntLegs
	case Spider:
		return SpiderLegs
	case Human:
		return HumanLegs
	}

	return *new(int)
}

// IsValid returns true if the enum value is, in fact, valid.
func (e Creatures) IsValid() bool {
	for _, v := range _CreaturesValues {
		if v == e {
			return true
		}
	}
	return false
}

// Values returns a list of all potential values of this enum.
func (Creatures) Values() []Creatures {
	return slices.Clone(_CreaturesValues)
}

// StringValues returns a list of all potential values of this enum as strings.
// Note: This does not return duplicates.
func (Creatures) StringValues() []string {
	return []string{
		"NotCreature",
		"Cat",
		"Dog",
		"Ant",
		"Spider",
		"Human",
		"SeaAnemone",
	}
}

// String returns a string representation of this enum.
// Note: in the case of duplicate values only the first alphabetical definition will be choosen.
func (e Creatures) String() string {
	switch e {
	case NotCreature:
		return "NotCreature"
	case Cat:
		return "Cat"
	case Dog:
		return "Dog"
	case Ant:
		return "Ant"
	case Spider:
		return "Spider"
	case Human:
		return "Human"
	case SeaAnemone:
		return "SeaAnemone"
	default:
		return fmt.Sprintf("UndefinedCreatures:%d", e)
	}
}

// ParseString will return a value as defined in string form.
func (e Creatures) ParseString(text string) (Creatures, error) {
	return ParseCreatures(text)
}

// ParseCreatures will attempt to parse the value of a Creatures from either its string form
// or any value of a trait flagged with the --parsableByTrait flag
func ParseCreatures(input any) (Creatures, error) {
	switch input {
	case "NotCreature":
		return NotCreature, nil
	case "Cat":
		return Cat, nil
	case "Feline":
		return Feline, nil
	case "Feline2":
		return Feline2, nil
	case "Dog":
		return Dog, nil
	case "Ant":
		return Ant, nil
	case "Spider":
		return Spider, nil
	case "Human":
		return Human, nil
	case "SeaAnemone":
		return SeaAnemone, nil
	default:
		if text, ok := input.(string); ok {
			switch strings.ToLower(text) {
			case "notcreature":
				return NotCreature, nil
			case "cat":
				return Cat, nil
			case "feline":
				return Feline, nil
			case "feline2":
				return Feline2, nil
			case "dog":
				return Dog, nil
			case "ant":
				return Ant, nil
			case "spider":
				return Spider, nil
			case "human":
				return Human, nil
			case "seaanemone":
				return SeaAnemone, nil
			}
		}
		return 0, fmt.Errorf("`%+v` could not be parsed to enum of type Creatures", input)
	}
}

// ParseGeneric calls TypedEnum.Parse but returns the result
// in the generic genum.Enum interface. Which is useful when you are only able to work with
// the un-typed interface.
func (e Creatures) ParseGeneric(input any) (genum.Enum, error) {
	return ParseCreatures(input)
}

// MarshalJSON implements the json.Marshaler interface for Creatures.
func (e Creatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Creatures.
func (e *Creatures) UnmarshalJSON(data []byte) error {
	// We always support strings.
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		var err error
		*e, err = ParseCreatures(s)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("unable to unmarshal Creatures from `%v`", data)
}

// MarshalText implements the encoding.TextMarshaler interface for Creatures.
func (e Creatures) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Creatures.
func (e *Creatures) UnmarshalText(text []byte) error {
	s := string(text)
	var err error
	*e, err = ParseCreatures(s)
	if err == nil {
		return nil
	}

	return fmt.Errorf("unable to unmarshal Creatures from `%s`", s)
}

// MarshalYAML implements a YAML Marshaler for Creatures.
func (e Creatures) MarshalYAML() (any, error) {
	return e.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Creatures.
func (e *Creatures) UnmarshalYAML(value *yaml.Node) error {
	var err error

	// first try and parse as a string
	*e, err = ParseCreatures(value.Value)
	if err == nil {
		return nil
	}

	// then try and parse for any string-like traits

	return fmt.Errorf("unable to unmarshal Creatures from yaml `%s`", value.Value)
}

// IsEnum implements an empty function required to implement Enum.
func (Creatures) IsEnum() {}

var _EnumWithPackageImportsValues = []EnumWithPackageImports{
	EnumWithPackageImports0,
	EnumWithPackageImports1,
	EnumWithPackageImports2,
}

// Kind returns the enum's associated trait of the same name.
// If no trait exists for the enumeration a default value will be returned.
func (e EnumWithPackageImports) Kind() reflect.Kind {
	switch e {
	case EnumWithPackageImports0:
		return _Kind
	case EnumWithPackageImports1:
		return reflect.Uint64
	case EnumWithPackageImports2:
		return reflect.Bool
	}

	return *new(reflect.Kind)
}

// IsValid returns true if the enum value is, in fact, valid.
func (e EnumWithPackageImports) IsValid() bool {
	for _, v := range _EnumWithPackageImportsValues {
		if v == e {
			return true
		}
	}
	return false
}

// Values returns a list of all potential values of this enum.
func (EnumWithPackageImports) Values() []EnumWithPackageImports {
	return slices.Clone(_EnumWithPackageImportsValues)
}

// StringValues returns a list of all potential values of this enum as strings.
// Note: This does not return duplicates.
func (EnumWithPackageImports) StringValues() []string {
	return []string{
		"EnumWithPackageImports0",
		"EnumWithPackageImports1",
		"EnumWithPackageImports2",
	}
}

// String returns a string representation of this enum.
// Note: in the case of duplicate values only the first alphabetical definition will be choosen.
func (e EnumWithPackageImports) String() string {
	switch e {
	case EnumWithPackageImports0:
		return "EnumWithPackageImports0"
	case EnumWithPackageImports1:
		return "EnumWithPackageImports1"
	case EnumWithPackageImports2:
		return "EnumWithPackageImports2"
	default:
		return fmt.Sprintf("UndefinedEnumWithPackageImports:%d", e)
	}
}

// ParseString will return a value as defined in string form.
func (e EnumWithPackageImports) ParseString(text string) (EnumWithPackageImports, error) {
	return ParseEnumWithPackageImports(text)
}

// ParseEnumWithPackageImports will attempt to parse the value of a EnumWithPackageImports from either its string form
// or any value of a trait flagged with the --parsableByTrait flag
func ParseEnumWithPackageImports(input any) (EnumWithPackageImports, error) {
	switch input {
	case "EnumWithPackageImports0":
		return EnumWithPackageImports0, nil
	case "EnumWithPackageImports1":
		return EnumWithPackageImports1, nil
	case "EnumWithPackageImports2":
		return EnumWithPackageImports2, nil
	default:
		if text, ok := input.(string); ok {
			switch strings.ToLower(text) {
			case "enumwithpackageimports0":
				return EnumWithPackageImports0, nil
			case "enumwithpackageimports1":
				return EnumWithPackageImports1, nil
			case "enumwithpackageimports2":
				return EnumWithPackageImports2, nil
			}
		}
		return 0, fmt.Errorf("`%+v` could not be parsed to enum of type EnumWithPackageImports", input)
	}
}

// ParseGeneric calls TypedEnum.Parse but returns the result
// in the generic genum.Enum interface. Which is useful when you are only able to work with
// the un-typed interface.
func (e EnumWithPackageImports) ParseGeneric(input any) (genum.Enum, error) {
	return ParseEnumWithPackageImports(input)
}

// MarshalJSON implements the json.Marshaler interface for EnumWithPackageImports.
func (e EnumWithPackageImports) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for EnumWithPackageImports.
func (e *EnumWithPackageImports) UnmarshalJSON(data []byte) error {
	// We always support strings.
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		var err error
		*e, err = ParseEnumWithPackageImports(s)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("unable to unmarshal EnumWithPackageImports from `%v`", data)
}

// MarshalText implements the encoding.TextMarshaler interface for EnumWithPackageImports.
func (e EnumWithPackageImports) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for EnumWithPackageImports.
func (e *EnumWithPackageImports) UnmarshalText(text []byte) error {
	s := string(text)
	var err error
	*e, err = ParseEnumWithPackageImports(s)
	if err == nil {
		return nil
	}

	return fmt.Errorf("unable to unmarshal EnumWithPackageImports from `%s`", s)
}

// MarshalYAML implements a YAML Marshaler for EnumWithPackageImports.
func (e EnumWithPackageImports) MarshalYAML() (any, error) {
	return e.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for EnumWithPackageImports.
func (e *EnumWithPackageImports) UnmarshalYAML(value *yaml.Node) error {
	var err error

	// first try and parse as a string
	*e, err = ParseEnumWithPackageImports(value.Value)
	if err == nil {
		return nil
	}

	// then try and parse for any string-like traits

	return fmt.Errorf("unable to unmarshal EnumWithPackageImports from yaml `%s`", value.Value)
}

// IsEnum implements an empty function required to implement Enum.
func (EnumWithPackageImports) IsEnum() {}
