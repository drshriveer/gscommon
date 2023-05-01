package internal

import (
	"encoding/json"
	"fmt"
)

// IsValid has a terrible implementation, but returns true if the value is, well, valid.
func (e MyEnum) IsValid() bool {
	_, err := e.ParseString(e.String())
	return err == nil
}

// Values returns a list of all potential values of this enum.
func (e MyEnum) Values() []MyEnum {
	return []MyEnum{
		IntentionallyNegative,
		EnumOneComplicationZero,
		EnumThreeComplicationOne,
		EnumTwoComplicationOne,
		EnumTwoComplicationZero,
		UNSET,
		EnumOneComplicationOne,
		EnumThreeComplicationThree,
		EnumTwoComplicationThree,
		EnumTwoComplicationTwo,
		ValueOne,
		EnumOneComplicationTwo,
		ValueTwo,
		EnumThreeComplicationTwo,
		EnumThreeComplicationZero,
		ValueSeven,
	}
}

// StringValues returns a list of all potential values of this enum as strings.
func (e MyEnum) StringValues() []string {
	return []string{
		"IntentionallyNegative",
		"EnumOneComplicationZero",
		"EnumThreeComplicationOne",
		"EnumTwoComplicationOne",
		"EnumTwoComplicationZero",
		"UNSET",
		"EnumOneComplicationOne",
		"EnumThreeComplicationThree",
		"EnumTwoComplicationThree",
		"EnumTwoComplicationTwo",
		"ValueOne",
		"EnumOneComplicationTwo",
		"ValueTwo",
		"EnumThreeComplicationTwo",
		"EnumThreeComplicationZero",
		"ValueSeven",
	}
}

// String returns a string representation of this enum.
// Note: in the case of duplicate values only the first alphabetical definition will be choosen.
func (e MyEnum) String() string {
	switch e {
	case IntentionallyNegative:
		return "IntentionallyNegative"
	case UNSET:
		return "UNSET"
	case EnumOneComplicationOne:
		return "EnumOneComplicationOne"
	case EnumOneComplicationTwo:
		return "EnumOneComplicationTwo"
	case EnumThreeComplicationTwo:
		return "EnumThreeComplicationTwo"
	default:
		return fmt.Sprintf("UndefinedMyEnum:%d", e)
	}
}

// ParseString will return a value as defined in string form.
func (e MyEnum) ParseString(text string) (MyEnum, error) {
	switch text {
	case "IntentionallyNegative":
		return IntentionallyNegative, nil
	case "EnumOneComplicationZero":
		return EnumOneComplicationZero, nil
	case "EnumThreeComplicationOne":
		return EnumThreeComplicationOne, nil
	case "EnumTwoComplicationOne":
		return EnumTwoComplicationOne, nil
	case "EnumTwoComplicationZero":
		return EnumTwoComplicationZero, nil
	case "UNSET":
		return UNSET, nil
	case "EnumOneComplicationOne":
		return EnumOneComplicationOne, nil
	case "EnumThreeComplicationThree":
		return EnumThreeComplicationThree, nil
	case "EnumTwoComplicationThree":
		return EnumTwoComplicationThree, nil
	case "EnumTwoComplicationTwo":
		return EnumTwoComplicationTwo, nil
	case "ValueOne":
		return ValueOne, nil
	case "EnumOneComplicationTwo":
		return EnumOneComplicationTwo, nil
	case "ValueTwo":
		return ValueTwo, nil
	case "EnumThreeComplicationTwo":
		return EnumThreeComplicationTwo, nil
	case "EnumThreeComplicationZero":
		return EnumThreeComplicationZero, nil
	case "ValueSeven":
		return ValueSeven, nil
	default:
		return 0, fmt.Errorf("`%s` is not a valid enum of type MyEnum", text)
	}
}

// MarshalJSON implements the json.Marshaler interface for MyEnum.
func (e MyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for MyEnum.
func (e *MyEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		var err error
		*e, err = MyEnum(0).ParseString(s)
		return err
	}
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*e = MyEnum(i)
		if e.IsValid() {
			return nil
		}
	}

	return fmt.Errorf("unable to unmarshal MyEnum from `%v`", data)
}

// MarshalText implements the encoding.TextMarshaler interface for MyEnum.
func (e MyEnum) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for MyEnum.
func (e *MyEnum) UnmarshalText(text []byte) error {
	var err error
	*e, err = MyEnum(0).ParseString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for MyEnum.
func (e MyEnum) MarshalYAML() (any, error) {
	return e.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for MyEnum.
func (e *MyEnum) UnmarshalYAML(unmarshal func(any) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		var err error
		*e, err = MyEnum(0).ParseString(s)
		return err
	}
	var i int
	if err := unmarshal(&i); err == nil {
		*e = MyEnum(i)
		if e.IsValid() {
			return nil
		}
	}

	return fmt.Errorf("unable to unmarshal MyEnum from yaml")
}
