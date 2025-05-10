package vm

import (
	"fmt"
)

// Type represents the type of a runtime value.
type Type string

const (
	NULL_OBJ    Type = "NULL"
	BOOLEAN_OBJ Type = "BOOLEAN"
	INTEGER_OBJ Type = "INTEGER"
	FLOAT_OBJ   Type = "FLOAT"
	STRING_OBJ  Type = "STRING"
	// Add other types later: ARRAY_OBJ, HASH_OBJ, FUNCTION_OBJ, ITERATOR_OBJ, etc.
)

// Value is the interface that all runtime values must implement.
type Value interface {
	Type() Type      // Returns the type of the value
	Inspect() string // Returns a string representation of the value for debugging/printing
	// Add methods for operations here later, e.g.,
	// Add(Value) (Value, error)
	// Subtract(Value) (Value, error)
	// ... etc. - Moved to VM helper functions for now for centralized type checking
}

// Null represents the 'nil' value.
type Null struct{}

func (n *Null) Type() Type      { return NULL_OBJ }
func (n *Null) Inspect() string { return "nil" }

// Boolean represents boolean values (true/false).
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type      { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Integer represents integer values.
type Integer struct {
	Value int64
}

func (i *Integer) Type() Type      { return INTEGER_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Float represents floating-point values.
type Float struct {
	Value float64
}

func (f *Float) Type() Type      { return FLOAT_OBJ }
func (f *Float) Inspect() string { return fmt.Sprintf("%f", f.Value) }

// String represents string values.
type String struct {
	Value string
}

func (s *String) Type() Type      { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value }

// Define singleton instances for Null and Booleans for efficiency.
var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

// GoBoolToBoolean converts a Go boolean to a VM Boolean value.
func GoBoolToBoolean(input bool) *Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// IsTruthy checks if a Value is considered "truthy" in the language.
// Typically, nil and false are falsy, everything else is truthy.
func IsTruthy(obj Value) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		// Numbers (non-zero?), strings (non-empty?), arrays (non-empty?), hashes (non-empty?), functions are typically truthy.
		// For now, assume any non-null/non-boolean is truthy. Refine later for collection types.
		return true
	}
}

// Add more value types (Array, Hash, Function, etc.) and their methods here as you implement them.
