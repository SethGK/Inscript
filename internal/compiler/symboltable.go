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
	Index int // Index in the globals array, locals array, or free array
}

// SymbolTable manages symbols in a scope, including free symbols for closures.
type SymbolTable struct {
	Outer           *SymbolTable       // Outer scope for resolution
	store           map[string]*Symbol // Maps names to symbols in this scope
	freeSymbols     []*Symbol          // Symbols captured from outer (non-global) scopes
	numDefinitions  int                // Number of locally defined symbols
	isFunctionScope bool               // True if this is a function scope
}

// NewSymbolTable creates a new global symbol table.
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store:           make(map[string]*Symbol),
		freeSymbols:     []*Symbol{},
		isFunctionScope: false,
	}
}

// NewEnclosedSymbolTable creates a new nested symbol table; isFunc indicates a function scope.
func NewEnclosedSymbolTable(outer *SymbolTable, isFunc bool) *SymbolTable {
	return &SymbolTable{
		Outer:           outer,
		store:           make(map[string]*Symbol),
		freeSymbols:     []*Symbol{},
		isFunctionScope: isFunc,
	}
}

// Define adds a new symbol in the current scope.
func (s *SymbolTable) Define(name string, kind SymbolKind) *Symbol {
	sym := &Symbol{Name: name, Kind: kind, Index: s.numDefinitions}
	s.store[name] = sym
	s.numDefinitions++
	return sym
}

// DefineBuiltin registers a builtin symbol in the given index (globals table).
func (s *SymbolTable) DefineBuiltin(index int, name string) *Symbol {
	sym := &Symbol{Name: name, Kind: Builtin, Index: index}
	s.store[name] = sym
	return sym
}

// Resolve looks up a name in current, then outer, returning the symbol.
// It also records free variables when resolving from a non-global outer function scope.
func (s *SymbolTable) Resolve(name string) (*Symbol, bool) {
	// Check current scope
	if sym, ok := s.store[name]; ok {
		return sym, true
	}
	// If no outer, not found
	if s.Outer == nil {
		return nil, false
	}
	// Recursively resolve in outer
	sym, ok := s.Outer.Resolve(name)
	if !ok {
		return nil, false
	}
	// If found in global or builtin, just return
	if sym.Kind == Global || sym.Kind == Builtin {
		return sym, true
	}
	// Otherwise, it's a free variable for this scope
	freeSym := &Symbol{Name: sym.Name, Kind: Free, Index: len(s.freeSymbols)}
	s.freeSymbols = append(s.freeSymbols, sym)
	// Register in current store so subsequent resolves use local free slot
	s.store[name] = freeSym
	return freeSym, true
}

// FreeSymbols returns the list of captured free symbols.
func (s *SymbolTable) FreeSymbols() []*Symbol {
	return s.freeSymbols
}

// NumDefinitions returns the number of symbols defined in this table.
func (s *SymbolTable) NumDefinitions() int {
	return s.numDefinitions
}

// Debug prints the contents of the symbol table (for development).
func (s *SymbolTable) Debug() {
	fmt.Printf("Symbols (funcScope=%v):\n", s.isFunctionScope)
	for name, sym := range s.store {
		fmt.Printf("  %s -> kind=%s idx=%d\n", name, sym.Kind, sym.Index)
	}
	if len(s.freeSymbols) > 0 {
		fmt.Println("Free symbols:")
		for i, fs := range s.freeSymbols {
			fmt.Printf("  [%d] %s\n", i, fs.Name)
		}
	}
}
