package compiler // Package compiler, as it's in the internal/compiler directory

// No imports needed for vm or types packages based on current structure

// SymbolKind represents the kind of a symbol (global, local, parameter, builtin).
type SymbolKind string

const (
	Global    SymbolKind = "GLOBAL"
	Local     SymbolKind = "LOCAL"
	Parameter SymbolKind = "PARAMETER"
	Builtin   SymbolKind = "BUILTIN"
)

// Symbol represents a compiled symbol (variable, function name, etc.).
type Symbol struct {
	Name  string
	Kind  SymbolKind
	Index int // Index in the globals array or locals array
}

// SymbolTable manages symbols in a scope.
type SymbolTable struct {
	Outer *SymbolTable // Outer scope for variable resolution

	store map[string]*Symbol // Maps symbol names to Symbols

	// For local scopes (functions/blocks), tracks the number of defined variables.
	// This determines the size of the locals array needed for a frame.
	numDefinitions int

	// Indicates if this symbol table represents a function scope.
	// This is used by the compiler to determine whether to emit OpGetLocal/OpSetLocal
	// or OpGetGlobal/OpSetGlobal for variable access.
	isFunctionScope bool
}

// NewSymbolTable creates a new symbol table (for the global scope).
// The global scope is not a function scope.
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store:           make(map[string]*Symbol),
		isFunctionScope: false, // Global scope is not a function scope
	}
}

// NewEnclosedSymbolTable creates a new symbol table with an outer scope.
// isFunc should be true if this new table represents a function scope, false otherwise (block scope).
func NewEnclosedSymbolTable(outer *SymbolTable, isFunc bool) *SymbolTable {
	s := NewSymbolTable() // Creates a base table (with isFunctionScope=false initially)
	s.Outer = outer
	s.isFunctionScope = isFunc // Set the correct function scope status
	return s
}

// Define defines a new symbol in the current scope.
// Returns the created Symbol.
func (s *SymbolTable) Define(name string, kind SymbolKind) *Symbol {
	symbol := &Symbol{Name: name, Kind: kind, Index: s.numDefinitions}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

// Resolve looks up a symbol in the current and outer scopes.
// Returns the Symbol and true if found, nil and false otherwise.
func (s *SymbolTable) Resolve(name string) (*Symbol, bool) {
	symbol, ok := s.store[name]
	if !ok && s.Outer != nil {
		// If not found in current scope, try outer scope.
		return s.Outer.Resolve(name)
	}
	return symbol, ok
}
