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
	// --- Debug Print: Defining symbol ---
	fmt.Printf("DEBUG (SymbolTable): Defining '%s' as kind %s in scope %p (numDefinitions: %d)\n", name, kind, s, s.numDefinitions)
	// --- End Debug Print ---

	sym := &Symbol{Name: name, Kind: kind, Index: s.numDefinitions}
	s.store[name] = sym
	s.numDefinitions++

	// --- Debug Print: Store after define ---
	// Uncomment the line below if you want to see the store contents after every definition
	// s.DebugStore()
	// --- End Debug Print ---

	return sym
}

// DefineBuiltin registers a builtin symbol in the given index (globals table).
func (s *SymbolTable) DefineBuiltin(index int, name string) *Symbol {
	sym := &Symbol{Name: name, Kind: Builtin, Index: index}
	s.store[name] = sym // Builtins go directly into the store
	// Builtins don't count towards numDefinitions for local/global index allocation

	// --- Debug Print: Defining builtin ---
	fmt.Printf("DEBUG (SymbolTable): Defining builtin '%s' at index %d in scope %p\n", name, index, s)
	// --- End Debug Print ---

	return sym
}

// Resolve looks up a name in current, then outer, returning the symbol.
// It also records free variables when resolving from a non-global outer function scope.
func (s *SymbolTable) Resolve(name string) (*Symbol, bool) {
	// --- Debug Print: Resolving symbol ---
	fmt.Printf("DEBUG (SymbolTable): Resolving '%s' in scope %p (IsFunc: %v)\n", name, s, s.isFunctionScope)
	// Uncomment the line below if you want to see the store contents at the start of every resolve
	// s.DebugStore()
	// --- End Debug Print ---

	// Check current scope first (for locals and parameters defined locally)
	if sym, ok := s.store[name]; ok {
		// --- Debug Print: Found locally ---
		fmt.Printf("DEBUG (SymbolTable): Found '%s' locally in scope %p: kind %s, index %d\n", name, s, sym.Kind, sym.Index)
		// --- End Debug Print ---
		return sym, true // Found locally!
	}

	// If no outer, not found at all
	if s.Outer == nil {
		// --- Debug Print: Not found, no outer scope ---
		fmt.Printf("DEBUG (SymbolTable): '%s' not found, no outer scope from %p\n", name, s)
		// --- End Debug Print ---
		return nil, false
	}

	// Recursively resolve in outer scope
	// --- Debug Print: Resolving in outer scope ---
	fmt.Printf("DEBUG (SymbolTable): '%s' not found in scope %p, resolving in outer scope %p\n", name, s, s.Outer)
	// --- End Debug Print ---
	sym, ok := s.Outer.Resolve(name)
	if !ok {
		// --- Debug Print: Not found in outer scopes ---
		fmt.Printf("DEBUG (SymbolTable): '%s' not found in outer scopes from %p\n", name, s)
		// --- End Debug Print ---
		return nil, false // Not found in outer scopes
	}

	// If found in global or builtin in an outer scope, just return that symbol.
	// These are not captured as free variables.
	if sym.Kind == Global || sym.Kind == Builtin || sym.Kind == Parameter {
		// --- Debug Print: Found Global/Builtin in outer scope ---
		fmt.Printf("DEBUG (SymbolTable): Found '%s' as %s in outer scope. Not capturing.\n", name, sym.Kind)
		// --- End Debug Print ---
		return sym, true
	}

	// Otherwise, the symbol was found in a non-global/non-builtin outer scope.
	// This means it's a free variable for the current scope.
	// We need to capture it.

	// Check if this free variable has already been captured by this scope.
	// This prevents duplicate entries in freeSymbols for the same variable name.
	for _, fs := range s.freeSymbols {
		if fs.Name == name {
			// Already captured, return the existing free symbol
			fmt.Printf("→ Re-resolved captured free var '%s' as slot %d in scope %p\n", fs.Name, fs.Index, s)
			return fs, true
		}
	}

	// Add the symbol to the current scope's freeSymbols list.
	// The index of the free symbol is its position in this list.
	freeIndex := len(s.freeSymbols)
	freeSym := &Symbol{Name: sym.Name, Kind: Free, Index: freeIndex}
	s.freeSymbols = append(s.freeSymbols, freeSym)

	// Debug print from the original code, kept for continuity
	fmt.Printf("→ Captured free var '%s' as slot %d in scope %p\n", freeSym.Name, freeSym.Index, s)

	// Return the newly created Free symbol for the current scope.
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
	fmt.Printf("Scope %p (funcScope=%v) Symbols:\n", s, s.isFunctionScope)
	s.DebugStore() // Use the new method
	if len(s.freeSymbols) > 0 {
		fmt.Println("Free symbols:")
		for i, fs := range s.freeSymbols {
			// Note: The Kind and Index here are for the Free symbol within THIS scope's freeSymbols list.
			// They don't represent the original kind/index in the outer scope where it was defined.
			fmt.Printf("  [%d] %s (Kind: %s, Index: %d)\n", i, fs.Name, fs.Kind, fs.Index)
		}
	}
	if s.Outer != nil {
		fmt.Printf("Outer scope: %p\n", s.Outer)
		// Uncomment the line below for full recursive symbol table debugging
		// s.Outer.Debug()
	}
}

// DebugStore prints the contents of the store map.
func (s *SymbolTable) DebugStore() {
	fmt.Printf("  Store (%d entries):\n", len(s.store))
	if len(s.store) == 0 {
		fmt.Println("    <empty>")
		return
	}
	for name, sym := range s.store {
		fmt.Printf("    '%s' -> kind=%s idx=%d\n", name, sym.Kind, sym.Index)
	}
}

func (s *SymbolTable) DefineLocal(name string) *Symbol {
	return s.Define(name, Local)
}

func (s *SymbolTable) DefineParameter(name string) *Symbol {
	return s.Define(name, Parameter)
}
