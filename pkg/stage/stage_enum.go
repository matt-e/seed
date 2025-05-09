// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package stage

import (
	"fmt"
	"strings"
)

const (
	// Test is a Stage of type Test.
	Test Stage = iota
	// Dev is a Stage of type Dev.
	Dev
	// Staging is a Stage of type Staging.
	Staging
	// Prod is a Stage of type Prod.
	Prod
)

var ErrInvalidStage = fmt.Errorf("not a valid Stage, try [%s]", strings.Join(_StageNames, ", "))

const _StageName = "testdevstagingprod"

var _StageNames = []string{
	_StageName[0:4],
	_StageName[4:7],
	_StageName[7:14],
	_StageName[14:18],
}

// StageNames returns a list of possible string values of Stage.
func StageNames() []string {
	tmp := make([]string, len(_StageNames))
	copy(tmp, _StageNames)
	return tmp
}

var _StageMap = map[Stage]string{
	Test:    _StageName[0:4],
	Dev:     _StageName[4:7],
	Staging: _StageName[7:14],
	Prod:    _StageName[14:18],
}

// String implements the Stringer interface.
func (x Stage) String() string {
	if str, ok := _StageMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Stage(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Stage) IsValid() bool {
	_, ok := _StageMap[x]
	return ok
}

var _StageValue = map[string]Stage{
	_StageName[0:4]:                    Test,
	strings.ToLower(_StageName[0:4]):   Test,
	_StageName[4:7]:                    Dev,
	strings.ToLower(_StageName[4:7]):   Dev,
	_StageName[7:14]:                   Staging,
	strings.ToLower(_StageName[7:14]):  Staging,
	_StageName[14:18]:                  Prod,
	strings.ToLower(_StageName[14:18]): Prod,
}

// ParseStage attempts to convert a string to a Stage.
func ParseStage(name string) (Stage, error) {
	if x, ok := _StageValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _StageValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Stage(0), fmt.Errorf("%s is %w", name, ErrInvalidStage)
}

// MarshalText implements the text marshaller method.
func (x Stage) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Stage) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseStage(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
