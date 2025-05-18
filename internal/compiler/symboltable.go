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
	Outer           *SymbolTable
	store           map[string]*Symbol
	freeSymbols     []*Symbol
	numDefinitions  int
	isFunctionScope bool
}

// NewSymbolTable creates a new global symbol table.
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store:           make(map[string]*Symbol),
		freeSymbols:     []*Symbol{},
		isFunctionScope: false,
	}
}

// NewEnclosedSymbolTable creates a new nested symbol table; isFunc indicates function scope.
func NewEnclosedSymbolTable(outer *SymbolTable, isFunc bool) *SymbolTable {
	return &SymbolTable{
		Outer:           outer,
		store:           make(map[string]*Symbol),
		freeSymbols:     []*Symbol{},
		isFunctionScope: isFunc,
	}
}

// DefineGlobal adds a new global symbol (module-level).
func (s *SymbolTable) DefineGlobal(name string) *Symbol {
	sym := &Symbol{Name: name, Kind: Global, Index: s.numDefinitions}
	s.store[name] = sym
	s.numDefinitions++
	return sym
}

// DefineLocal adds a new local symbol (block or function scope).
func (s *SymbolTable) DefineLocal(name string) *Symbol {
	sym := &Symbol{Name: name, Kind: Local, Index: s.numDefinitions}
	s.store[name] = sym
	s.numDefinitions++
	return sym
}

// DefineParameter adds a new parameter symbol in a function scope.
func (s *SymbolTable) DefineParameter(name string) *Symbol {
	sym := &Symbol{Name: name, Kind: Parameter, Index: s.numDefinitions}
	s.store[name] = sym
	s.numDefinitions++
	return sym
}

// DefineBuiltin registers a builtin symbol in the given index.
func (s *SymbolTable) DefineBuiltin(index int, name string) *Symbol {
	sym := &Symbol{Name: name, Kind: Builtin, Index: index}
	s.store[name] = sym
	return sym
}

// Resolve looks up a name in current then outer scopes, capturing free variables.
func (s *SymbolTable) Resolve(name string) (*Symbol, bool) {
	if sym, ok := s.store[name]; ok {
		return sym, true
	}

	if s.Outer == nil {
		return nil, false
	}

	sym, ok := s.Outer.Resolve(name)
	if !ok {
		return nil, false
	}

	if sym.Kind == Global || sym.Kind == Builtin || sym.Kind == Parameter {
		return sym, true
	}

	// Capture free symbol
	freeSym := &Symbol{Name: sym.Name, Kind: Free, Index: len(s.freeSymbols)}
	s.freeSymbols = append(s.freeSymbols, freeSym)
	return freeSym, true
}

// FreeSymbols returns symbols captured from outer scopes.
func (s *SymbolTable) FreeSymbols() []*Symbol {
	return s.freeSymbols
}

// NumDefinitions returns the number of symbols defined in this table.
func (s *SymbolTable) NumDefinitions() int {
	return s.numDefinitions
}

// Debug prints the symbol table for development.
func (s *SymbolTable) Debug() {
	fmt.Printf("SymbolTable %p (funcScope=%v):\n", s, s.isFunctionScope)
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
