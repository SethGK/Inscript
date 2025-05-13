package types // New package for shared types

import (
	"fmt"
	"strings"
	// Note: This package does NOT import "compiler" or "vm" to break the cycle.
	// It only defines the interfaces and structs.
)

// Type represents the type of a runtime value.
type Type string

const (
	INTEGER_OBJ  Type = "INTEGER"
	FLOAT_OBJ    Type = "FLOAT"
	STRING_OBJ   Type = "STRING"
	BOOLEAN_OBJ  Type = "BOOLEAN"
	NULL_OBJ     Type = "NULL"
	LIST_OBJ     Type = "LIST"
	TABLE_OBJ    Type = "TABLE"
	FUNCTION_OBJ Type = "FUNCTION" // For CompiledFunction
	CLOSURE_OBJ  Type = "CLOSURE"
	ITERATOR_OBJ Type = "ITERATOR" // For iterators
	ERROR_OBJ    Type = "ERROR"    // For runtime errors
)

// Value interface represents any runtime value.
// Defined in the types package.
type Value interface {
	Type() Type
	Inspect() string                       // String representation for printing
	Equals(other Value) bool               // Check for equality
	Compare(other Value) (int, error)      // Compare for ordering (<, >, <=, >=)
	GetIterator() (Iterator, error)        // Method to get an iterator for iterable types
	GetIndex(index Value) (Value, error)   // Method to get an element by index
	SetIndex(index Value, val Value) error // Method to set an element by index
}

// Iterator interface for implementing iteration (for loops).
// It now embeds the Value interface.
type Iterator interface {
	Value                       // An Iterator is also a Value
	Next() (Value, bool, error) // Returns next value, true if successful, error
}

// Integer value
type Integer struct { // Defined in the types package
	Value int64
}

func (i *Integer) Type() Type      { return INTEGER_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Equals(other Value) bool {
	if o, ok := other.(*Integer); ok {
		return i.Value == o.Value
	}
	return false
}
func (i *Integer) Compare(other Value) (int, error) {
	if o, ok := other.(*Integer); ok {
		if i.Value < o.Value {
			return -1, nil
		}
		if i.Value > o.Value {
			return 1, nil
		}
		return 0, nil
	}
	if o, ok := other.(*Float); ok {
		fVal := float64(i.Value)
		if fVal < o.Value {
			return -1, nil
		}
		if fVal > o.Value {
			return 1, nil
		}
		return 0, nil
	}
	return 0, fmt.Errorf("comparison not supported between Integer and %s", other.Type())
}
func (i *Integer) GetIterator() (Iterator, error) { return nil, fmt.Errorf("integer is not iterable") }
func (i *Integer) GetIndex(index Value) (Value, error) {
	return nil, fmt.Errorf("integer is not indexable")
}
func (i *Integer) SetIndex(index Value, val Value) error {
	return fmt.Errorf("integer is not indexable")
}

// NewInteger helper
func NewInteger(i int64) *Integer { return &Integer{Value: i} }

// Float value
type Float struct { // Defined in the types package
	Value float64
}

func (f *Float) Type() Type      { return FLOAT_OBJ }
func (f *Float) Inspect() string { return fmt.Sprintf("%f", f.Value) }
func (f *Float) Equals(other Value) bool {
	if o, ok := other.(*Float); ok {
		return f.Value == o.Value
	}
	return false
}
func (f *Float) Compare(other Value) (int, error) {
	if o, ok := other.(*Float); ok {
		if f.Value < o.Value {
			return -1, nil
		}
		if f.Value > o.Value {
			return 1, nil
		}
		return 0, nil
	}
	if o, ok := other.(*Integer); ok {
		iVal := float64(o.Value)
		if f.Value < iVal {
			return -1, nil
		}
		if f.Value > iVal {
			return 1, nil
		}
		return 0, nil
	}
	return 0, fmt.Errorf("comparison not supported between Float and %s", other.Type())
}
func (f *Float) GetIterator() (Iterator, error) { return nil, fmt.Errorf("float is not iterable") }
func (f *Float) GetIndex(index Value) (Value, error) {
	return nil, fmt.Errorf("float is not indexable")
}
func (f *Float) SetIndex(index Value, val Value) error { return fmt.Errorf("float is not indexable") }

// NewFloat helper
func NewFloat(f float64) *Float { return &Float{Value: f} }

// String value
type String struct { // Defined in the types package
	Value string
}

func (s *String) Type() Type      { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value } // Or add quotes: fmt.Sprintf(`"%s"`, s.Value)
func (s *String) Equals(other Value) bool {
	if o, ok := other.(*String); ok {
		return s.Value == o.Value
	}
	return false
}
func (s *String) Compare(other Value) (int, error) {
	if o, ok := other.(*String); ok {
		return strings.Compare(s.Value, o.Value), nil
	}
	return 0, fmt.Errorf("comparison not supported between String and %s", other.Type())
}
func (s *String) GetIterator() (Iterator, error) { return NewStringIterator(s), nil } // String is iterable
func (s *String) GetIndex(index Value) (Value, error) {
	idxInt, ok := index.(*Integer)
	if !ok {
		return nil, fmt.Errorf("string index must be an integer, got %s", index.Type())
	}
	idx := idxInt.Value
	if idx < 0 || idx >= int64(len(s.Value)) {
		return nil, fmt.Errorf("string index out of bounds: %d", idx)
	}
	return &String{Value: string(s.Value[idx])}, nil // Return a new string for the character
}
func (s *String) SetIndex(index Value, val Value) error {
	return fmt.Errorf("string does not support item assignment")
}

// NewString helper
func NewString(s string) *String { return &String{Value: s} }

// Boolean value
type Boolean struct { // Defined in the types package
	Value bool
}

func (b *Boolean) Type() Type      { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Equals(other Value) bool {
	if o, ok := other.(*Boolean); ok {
		return b.Value == o.Value
	}
	return false
}
func (b *Boolean) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for Boolean")
}
func (b *Boolean) GetIterator() (Iterator, error) { return nil, fmt.Errorf("boolean is not iterable") }
func (b *Boolean) GetIndex(index Value) (Value, error) {
	return nil, fmt.Errorf("boolean is not indexable")
}
func (b *Boolean) SetIndex(index Value, val Value) error {
	return fmt.Errorf("boolean is not indexable")
}

// Nil value
type Nil struct{}              // Defined in the types package
func (n *Nil) Type() Type      { return NULL_OBJ }
func (n *Nil) Inspect() string { return "nil" }
func (n *Nil) Equals(other Value) bool {
	_, ok := other.(*Nil)
	return ok
}
func (n *Nil) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for Nil")
}
func (n *Nil) GetIterator() (Iterator, error)        { return nil, fmt.Errorf("nil is not iterable") }
func (n *Nil) GetIndex(index Value) (Value, error)   { return nil, fmt.Errorf("nil is not indexable") }
func (n *Nil) SetIndex(index Value, val Value) error { return fmt.Errorf("nil is not indexable") }

// CompiledFunction value (represents the compiled code of a function)
// Defined in the types package.
type CompiledFunction struct {
	Instructions  []byte // The bytecode for this function (using byte slice)
	NumLocals     int    // Number of local variables (including parameters)
	NumParameters int    // Number of parameters the function expects
	// Free []Value // Free variables (for closures - TODO later)
}

func (cf *CompiledFunction) Type() Type              { return FUNCTION_OBJ }
func (cf *CompiledFunction) Inspect() string         { return fmt.Sprintf("<function at %p>", cf) }
func (cf *CompiledFunction) Equals(other Value) bool { return cf == other } // Identity equality for functions
func (cf *CompiledFunction) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for Function")
}
func (cf *CompiledFunction) GetIterator() (Iterator, error) {
	return nil, fmt.Errorf("function is not iterable")
}
func (cf *CompiledFunction) GetIndex(index Value) (Value, error) {
	return nil, fmt.Errorf("function is not indexable")
}
func (cf *CompiledFunction) SetIndex(index Value, val Value) error {
	return fmt.Errorf("function is not indexable")
}

// Closure value (represents a function instance with its environment)
// Defined in the types package.
type Closure struct {
	Fn *CompiledFunction // The compiled function - Referring to CompiledFunction directly
	// Free []Value // Free variables (for closures - TODO later)
}

func (c *Closure) Type() Type              { return CLOSURE_OBJ }
func (c *Closure) Inspect() string         { return fmt.Sprintf("<closure at %p>", c) }
func (c *Closure) Equals(other Value) bool { return c == other } // Identity equality for closures
func (c *Closure) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for Closure")
}
func (c *Closure) GetIterator() (Iterator, error) { return nil, fmt.Errorf("closure is not iterable") }
func (c *Closure) GetIndex(index Value) (Value, error) {
	return nil, fmt.Errorf("closure is not indexable")
}
func (c *Closure) SetIndex(index Value, val Value) error {
	return fmt.Errorf("closure is not indexable")
}

// List value
type List struct { // Defined in the types package
	Elements []Value
}

func (l *List) Type() Type { return LIST_OBJ }
func (l *List) Inspect() string {
	var elements []string
	for _, el := range l.Elements {
		elements = append(elements, el.Inspect())
	}
	return "[" + strings.Join(elements, ", ") + "]"
}
func (l *List) Equals(other Value) bool {
	o, ok := other.(*List)
	if !ok {
		return false
	}
	if len(l.Elements) != len(o.Elements) {
		return false
	}
	for i := range l.Elements {
		if !l.Elements[i].Equals(o.Elements[i]) {
			return false
		}
	}
	return true
}
func (l *List) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for List")
}
func (l *List) GetIterator() (Iterator, error) { return NewListIterator(l), nil } // List is iterable
func (l *List) GetIndex(index Value) (Value, error) {
	idxInt, ok := index.(*Integer)
	if !ok {
		return nil, fmt.Errorf("list index must be an integer, got %s", index.Type())
	}
	idx := idxInt.Value
	if idx < 0 || idx >= int64(len(l.Elements)) {
		return nil, fmt.Errorf("list index out of bounds: %d", idx)
	}
	return l.Elements[idx], nil
}
func (l *List) SetIndex(index Value, val Value) error {
	idxInt, ok := index.(*Integer)
	if !ok {
		return fmt.Errorf("list index must be an integer, got %s", index.Type())
	}
	idx := idxInt.Value
	if idx < 0 || idx >= int64(len(l.Elements)) {
		return fmt.Errorf("list index out of bounds: %d", idx)
	}
	l.Elements[idx] = val
	return nil
}

// NewList helper
func NewList(elements ...Value) *List { return &List{Elements: elements} }

// --- Changes for Ordered Table Start ---

// TablePair represents a single key-value pair within an ordered Table.
type TablePair struct {
	Key   string // Table keys are assumed to be strings
	Value Value
}

// Table value - Now stores pairs in order of insertion.
type Table struct { // Defined in the types package
	// Pairs stores the key-value pairs in the order they were added.
	Pairs []TablePair
	// Lookup provides O(1) average time complexity for accessing values by key.
	// It maps the key string to the index of the pair in the Pairs slice.
	Lookup map[string]int // Changed to uppercase 'L' to be exported
}

func (t *Table) Type() Type { return TABLE_OBJ }

// Inspect iterates over the ordered Pairs slice for consistent output.
func (t *Table) Inspect() string {
	var fields []string
	// Iterate over the ordered slice instead of the unordered map
	for _, pair := range t.Pairs {
		fields = append(fields, fmt.Sprintf("%s: %s", pair.Key, pair.Value.Inspect()))
	}
	return "{" + strings.Join(fields, ", ") + "}"
}

// Equals compares two Tables based on the order of their pairs.
func (t *Table) Equals(other Value) bool {
	o, ok := other.(*Table)
	if !ok {
		return false
	}
	// Compare lengths of the ordered slices
	if len(t.Pairs) != len(o.Pairs) {
		return false
	}
	// Compare each pair in order
	for i := range t.Pairs {
		if t.Pairs[i].Key != o.Pairs[i].Key || !t.Pairs[i].Value.Equals(o.Pairs[i].Value) {
			return false
		}
	}
	return true
}

func (t *Table) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for Table")
}

// GetIterator returns a TableIterator that iterates over the ordered keys.
func (t *Table) GetIterator() (Iterator, error) { return NewTableIterator(t), nil } // Table is iterable (iterates over keys in order)

// GetIndex retrieves a value by key using the Lookup map for efficiency.
func (t *Table) GetIndex(index Value) (Value, error) {
	keyStr, ok := index.(*String) // Assuming table keys are strings
	if !ok {
		return nil, fmt.Errorf("table index must be a string, got %s", index.Type())
	}

	// Use the Lookup map for O(1) average time complexity access
	if t.Lookup != nil { // Changed to uppercase 'L'
		if idx, found := t.Lookup[keyStr.Value]; found { // Changed to uppercase 'L'
			return t.Pairs[idx].Value, nil
		}
	} else {
		// Fallback to linear scan if Lookup map is not initialized (less efficient)
		for _, pair := range t.Pairs {
			if pair.Key == keyStr.Value {
				return pair.Value, nil
			}
		}
	}

	// Return nil for non-existent keys
	return &Nil{}, nil
}

// SetIndex sets or adds a value by key, maintaining insertion order.
func (t *Table) SetIndex(index Value, val Value) error {
	keyStr, ok := index.(*String) // Assuming table keys are strings
	if !ok {
		return fmt.Errorf("table index must be a string, got %s", index.Type())
	}

	key := keyStr.Value

	// Check if the key already exists using the Lookup map
	if t.Lookup != nil { // Changed to uppercase 'L'
		if idx, found := t.Lookup[key]; found { // Changed to uppercase 'L'
			t.Pairs[idx].Value = val // Update existing value in the ordered slice
			return nil
		}
	} else {
		// Fallback to linear scan if Lookup map is not initialized (less efficient)
		for i, pair := range t.Pairs {
			if pair.Key == key {
				t.Pairs[i].Value = val // Update existing value
				return nil
			}
		}
	}

	// If the key does not exist, append a new pair to the end (maintains insertion order)
	t.Pairs = append(t.Pairs, TablePair{Key: key, Value: val})
	// If using a Lookup map, add the new key and its index
	if t.Lookup != nil { // Changed to uppercase 'L'
		t.Lookup[key] = len(t.Pairs) - 1 // Changed to uppercase 'L'
	}

	return nil
}

// NewTable helper - Now takes a slice of TablePair to create an ordered table.
// Note: If you still need to convert from a map, you would need a different helper
// that accepts a map and defines an ordering (e.g., by sorting keys).
// This helper is intended for creating tables where the order is known at creation.
func NewTable(pairs []TablePair) *Table {
	// Initialize the Lookup map
	lookup := make(map[string]int, len(pairs))
	for i, pair := range pairs {
		lookup[pair.Key] = i
	}
	return &Table{Pairs: pairs, Lookup: lookup} // Changed to uppercase 'L'
}

// --- Changes for Ordered Table End ---

// StringIterator for iterating over strings
type StringIterator struct {
	str   *String
	index int
}

func NewStringIterator(s *String) *StringIterator  { return &StringIterator{str: s, index: 0} }
func (si *StringIterator) Type() Type              { return ITERATOR_OBJ } // Iterators have a type
func (si *StringIterator) Inspect() string         { return fmt.Sprintf("<string iterator at %p>", si) }
func (si *StringIterator) Equals(other Value) bool { return si == other } // Identity equality
func (si *StringIterator) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for Iterator")
}                                                         // Implement Compare
func (si *StringIterator) GetIterator() (Iterator, error) { return si, nil } // Iterators are their own iterators
func (si *StringIterator) GetIndex(index Value) (Value, error) {
	return nil, fmt.Errorf("iterator is not indexable")
} // Implement GetIndex
func (si *StringIterator) SetIndex(index Value, val Value) error {
	return fmt.Errorf("iterator is not indexable")
} // Implement SetIndex

func (si *StringIterator) Next() (Value, bool, error) {
	if si.index >= len(si.str.Value) {
		return &Nil{}, false, nil // Iteration is done
	}
	char := string(si.str.Value[si.index])
	si.index++
	return &String{Value: char}, true, nil // Return the character as a string value
}

// ListIterator for iterating over lists
type ListIterator struct {
	list  *List
	index int
}

func NewListIterator(l *List) *ListIterator      { return &ListIterator{list: l, index: 0} }
func (li *ListIterator) Type() Type              { return ITERATOR_OBJ } // Iterators have a type
func (li *ListIterator) Inspect() string         { return fmt.Sprintf("<list iterator at %p>", li) }
func (li *ListIterator) Equals(other Value) bool { return li == other } // Identity equality
func (li *ListIterator) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for Iterator")
}                                                       // Implement Compare
func (li *ListIterator) GetIterator() (Iterator, error) { return li, nil } // Iterators are their own iterators
func (li *ListIterator) GetIndex(index Value) (Value, error) {
	return nil, fmt.Errorf("iterator is not indexable")
} // Implement GetIndex
func (li *ListIterator) SetIndex(index Value, val Value) error {
	return fmt.Errorf("iterator is not indexable")
} // Implement SetIndex

func (li *ListIterator) Next() (Value, bool, error) {
	if li.index >= len(li.list.Elements) {
		return &Nil{}, false, nil // Iteration is done
	}
	value := li.list.Elements[li.index]
	li.index++
	return value, true, nil // Return the list element
}

// TableIterator for iterating over tables (iterates over keys in order)
type TableIterator struct {
	table *Table
	index int // Current index in the table's Pairs slice
}

// NewTableIterator creates a new iterator iterating over the ordered pairs.
func NewTableIterator(t *Table) *TableIterator {
	// The iterator now directly uses the ordered Pairs slice.
	return &TableIterator{table: t, index: 0}
}

func (ti *TableIterator) Type() Type              { return ITERATOR_OBJ } // Iterators have a type
func (ti *TableIterator) Inspect() string         { return fmt.Sprintf("<table iterator at %p>", ti) }
func (ti *TableIterator) Equals(other Value) bool { return ti == other } // Identity equality
func (ti *TableIterator) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for Iterator")
}                                                        // Implement Compare
func (ti *TableIterator) GetIterator() (Iterator, error) { return ti, nil } // Iterators are their own iterators
func (ti *TableIterator) GetIndex(index Value) (Value, error) {
	return nil, fmt.Errorf("iterator is not indexable")
} // Implement GetIndex
func (ti *TableIterator) SetIndex(index Value, val Value) error {
	return fmt.Errorf("iterator is not indexable")
} // Implement SetIndex

func (ti *TableIterator) Next() (Value, bool, error) {
	// Iterate through the ordered Pairs slice
	if ti.index >= len(ti.table.Pairs) {
		return &Nil{}, false, nil // Iteration is done
	}
	// Return the key of the current pair in order
	key := ti.table.Pairs[ti.index].Key
	ti.index++
	return &String{Value: key}, true, nil // Return the key as a string value
	// If you wanted to iterate over values or pairs, you would return those here instead.
}

// Error value for runtime errors
type Error struct {
	Message string
}

func (e *Error) Type() Type      { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
func (e *Error) Equals(other Value) bool {
	o, ok := other.(*Error)
	if !ok {
		return false
	}
	return e.Message == o.Message
}
func (e *Error) Compare(other Value) (int, error) {
	return 0, fmt.Errorf("comparison not supported for Error")
}
func (e *Error) GetIterator() (Iterator, error) { return nil, fmt.Errorf("error is not iterable") }
func (e *Error) GetIndex(index Value) (Value, error) {
	return nil, fmt.Errorf("error is not indexable")
}
func (e *Error) SetIndex(index Value, val Value) error { return fmt.Errorf("error is not indexable") }

// NewError helper
func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

// TODO: Add other value types as needed (e.g., Builtin)
