package compiler

import "fmt"

// SymbolKind represents the kind of a symbol (global, local, parameter, builtin, or free).
type SymbolKind string

const (
	Global    SymbolKind = "GLOBAL"
	Local     SymbolKind = "LOCAL"
	Parameter SymbolKind = "PARAMETER"
	Builtin   SymbolKind = "BUILTIN"
	Free      SymbolKind = "FREE"
)

// Symbol represents a compiled symbol (variable, function name, etc.).
type Symbol struct {
	Name  string
	Kind  SymbolKind
	Index int // Index in the global, local, parameter, or free arrays
}

// SymbolTable manages symbols in a scope, including free symbols for closures.
type SymbolTable struct {
	Outer                       *SymbolTable
	store                       map[string]*Symbol
	freeSymbols                 []*Symbol
	numLocalAndParamDefinitions int  // Counts locals and parameters only for THIS table
	nextGlobalIndex             int  // Only used by the top-level global symbol table
	isFunctionScope             bool // True if this symbol table represents a function's scope
}

// NewSymbolTable creates a new global symbol table.
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store:                       make(map[string]*Symbol),
		freeSymbols:                 []*Symbol{},
		numLocalAndParamDefinitions: 0,
		nextGlobalIndex:             0, // Initialize global index counter for the global scope
		isFunctionScope:             false,
	}
}

// NewEnclosedSymbolTable creates a new nested symbol table; isFunc indicates function scope.
func NewEnclosedSymbolTable(outer *SymbolTable, isFunc bool) *SymbolTable {
	return &SymbolTable{
		Outer:                       outer,
		store:                       make(map[string]*Symbol),
		freeSymbols:                 []*Symbol{},
		numLocalAndParamDefinitions: 0, // Reset for new local/function scope
		isFunctionScope:             isFunc,
	}
}

// DefineGlobal adds a new global symbol (module-level).
func (s *SymbolTable) DefineGlobal(name string) *Symbol {
	// Only the top-level global symbol table should define globals.
	if s.Outer != nil {
		panic("DefineGlobal called on a non-global symbol table")
	}
	sym := &Symbol{Name: name, Kind: Global, Index: s.nextGlobalIndex}
	s.store[name] = sym
	s.nextGlobalIndex++ // Increment global index
	return sym
}

// DefineLocal adds a new local symbol (block or function scope).
func (s *SymbolTable) DefineLocal(name string) *Symbol {
	sym := &Symbol{Name: name, Kind: Local, Index: s.numLocalAndParamDefinitions}
	s.store[name] = sym
	s.numLocalAndParamDefinitions++ // Increment local/param definition count
	return sym
}

// DefineParameter adds a new parameter symbol in a function scope.
func (s *SymbolTable) DefineParameter(name string) *Symbol {
	sym := &Symbol{Name: name, Kind: Parameter, Index: s.numLocalAndParamDefinitions}
	s.store[name] = sym
	s.numLocalAndParamDefinitions++ // Increment local/param definition count
	return sym
}

// DefineBuiltin registers a builtin symbol in the given index.
func (s *SymbolTable) DefineBuiltin(index int, name string) *Symbol {
	sym := &Symbol{Name: name, Kind: Builtin, Index: index}
	s.store[name] = sym
	return sym
}

// Resolve looks up a name in current then outer scopes.
// It also handles capturing free variables if the current scope is a function scope
// and the symbol is a local/parameter from an enclosing function scope.
func (s *SymbolTable) Resolve(name string) (*Symbol, bool) {
	// 1. Check current scope
	if sym, ok := s.store[name]; ok {
		return sym, true
	}

	// 2. If not found, check outer scope
	if s.Outer == nil {
		return nil, false // Not found in any scope
	}

	// Recursively resolve in the outer scope
	outerSym, ok := s.Outer.Resolve(name)
	if !ok {
		return nil, false // Not found in outer scopes either
	}

	// 3. If found in an outer scope, determine its kind relative to the current scope.
	// Global and Builtin symbols are always accessed directly, never become free.
	if outerSym.Kind == Global || outerSym.Kind == Builtin {
		return outerSym, true
	}

	// If the current scope is a function scope, and the resolved symbol
	// is a Local, Parameter, or Free from an outer scope, it needs to be captured
	// as a Free variable for the *current* function's closure.
	// This ensures variables from enclosing *function* scopes are correctly handled.
	if s.isFunctionScope && (outerSym.Kind == Local || outerSym.Kind == Parameter || outerSym.Kind == Free) {
		return s.DefineFree(outerSym), true
	}

	// If the current scope is NOT a function scope (i.e., it's a block scope)
	// OR if the symbol is not Local/Parameter/Free (which is handled above by Global/Builtin),
	// then it's still treated as its original kind relative to the current function's frame.
	// This covers cases like `count` in the `for` loop within `length` function.
	return outerSym, true
}

// DefineFree adds a symbol from an outer scope as a free variable in the current scope.
// It ensures that a symbol is only added once to the freeSymbols list.
func (s *SymbolTable) DefineFree(original *Symbol) *Symbol {
	// Check if it's already a free symbol in this table to avoid duplicates.
	for _, fs := range s.freeSymbols {
		if fs.Name == original.Name {
			return fs // Already captured, return the existing free symbol
		}
	}
	// If not already captured, define it as a new free symbol in this table.
	sym := &Symbol{Name: original.Name, Kind: Free, Index: len(s.freeSymbols)}
	s.freeSymbols = append(s.freeSymbols, sym)
	return sym
}

// FreeSymbols returns symbols captured from outer scopes.
func (s *SymbolTable) FreeSymbols() []*Symbol {
	return s.freeSymbols
}

// NumDefinitions returns the number of locals and parameters defined in this table.
func (s *SymbolTable) NumDefinitions() int {
	return s.numLocalAndParamDefinitions
}

// NumGlobalsInTable returns the number of global variables defined in this table (only for global table).
func (s *SymbolTable) NumGlobalsInTable() int {
	if s.Outer != nil {
		panic("NumGlobalsInTable called on a non-global symbol table")
	}
	return s.nextGlobalIndex
}

// Debug prints the symbol table for development.
func (s *SymbolTable) Debug() {
	fmt.Printf("SymbolTable %p (funcScope=%v, numLocalAndParamDefs=%d, nextGlobalIdx=%d):\n", s, s.isFunctionScope, s.numLocalAndParamDefinitions, s.nextGlobalIndex)
	for name, sym := range s.store {
		fmt.Printf("  %s -> Kind: %s, Index: %d\n", name, sym.Kind, sym.Index)
	}
	if len(s.freeSymbols) > 0 {
		fmt.Println("Free symbols:")
		for _, fs := range s.freeSymbols {
			fmt.Printf("  %s (Index: %d)\n", fs.Name, fs.Index)
		}
	}
	if s.Outer != nil {
		fmt.Println("Outer scope:")
		s.Outer.Debug()
	}
}
