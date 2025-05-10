package compiler

// SymbolKind represents whether a symbol is global, local, parameter, or builtin.
type SymbolKind int

const (
	Global SymbolKind = iota
	Local
	Parameter
	Builtin
)

// Symbol represents information about a variable, function, etc.
type Symbol struct {
	Name  string
	Kind  SymbolKind
	Index int // Index in globals array, locals array, etc.
}

// SymbolTable maps symbol names to Symbols.
// It supports nesting via the Outer field.
type SymbolTable struct {
	store          map[string]*Symbol
	numDefinitions int          // Number of symbols defined directly in this table (for locals/params)
	Outer          *SymbolTable // Link to the parent scope's symbol table
}

// NewSymbolTable creates a new symbol table.
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{store: make(map[string]*Symbol)}
}

// NewEnclosedSymbolTable creates a new symbol table nested within an outer one.
func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

// Define adds a new symbol to the table.
// It returns the created Symbol.
// This method is primarily used for defining Local and Parameter symbols.
// Global symbols are typically defined directly in the top-level compiler's global table.
func (s *SymbolTable) Define(name string, kind SymbolKind) *Symbol {
	if kind == Global {
		// Global symbols should be defined in the top-level compiler's global table.
		// This prevents accidental global definition in nested scopes.
		panic("Define(Global) called on non-global symbol table")
	}

	symbol := &Symbol{Name: name, Kind: kind, Index: s.numDefinitions}
	s.store[name] = symbol
	s.numDefinitions++ // Increment definition count for this scope (locals/params)
	return symbol
}

// Resolve finds a symbol in the table or its outer tables.
// It searches from the current scope outwards to the global scope.
func (s *SymbolTable) Resolve(name string) (*Symbol, bool) {
	// Check current scope
	obj, ok := s.store[name]
	if ok {
		return obj, true
	}

	// If not found in current scope, try outer scopes recursively
	if s.Outer != nil {
		return s.Outer.Resolve(name)
	}

	// Not found in any scope
	return nil, false
}

// NumDefinitions returns the number of symbols defined directly in this table.
// This is useful for determining the number of local variables/parameters in a function scope.
func (s *SymbolTable) NumDefinitions() int {
	return s.numDefinitions
}

// GlobalTable is a helper to get the top-most (global) symbol table.
func (s *SymbolTable) GlobalTable() *SymbolTable {
	current := s
	for current.Outer != nil {
		current = current.Outer
	}
	return current
}
