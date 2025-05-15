package main

import (
	"fmt"
	"os"

	"github.com/SethGK/Inscript/internal/ast"          // Import AST package
	"github.com/SethGK/Inscript/internal/compiler"     // Import Compiler package
	"github.com/SethGK/Inscript/internal/types"        // Import types package
	"github.com/SethGK/Inscript/internal/vm"           // Import VM package
	parser "github.com/SethGK/Inscript/parser/grammar" // Import the ANTLR generated parser package
	"github.com/antlr4-go/antlr/v4"                    // Import ANTLR runtime
)

func main() {
	// 1. Read Source Code (Example: from a file specified as a command-line argument)
	if len(os.Args) < 2 {
		// Updated usage message with .ins extension
		fmt.Println("Usage: inscript <source_file.ins>")
		os.Exit(1)
	}
	filePath := os.Args[1]
	input, err := antlr.NewFileStream(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filePath, err)
		os.Exit(1)
	}

	// 2. Lex and Parse
	lexer := parser.NewInscriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewInscriptParser(stream)

	// Add an error listener to capture syntax errors
	// The DiagnosticErrorListener will print errors to stderr automatically.
	// You might want to implement a custom error listener for more control.
	errorListener := antlr.NewDiagnosticErrorListener(true)
	p.AddErrorListener(errorListener)
	lexer.AddErrorListener(errorListener) // Also add to the lexer

	p.SetErrorHandler(antlr.NewDefaultErrorStrategy()) // Use default error recovery

	// Parse the program rule
	// The error listener will report errors during this call.
	parseTree := p.Program()

	// We don't have a direct way to get the *count* of errors from the default listener.
	// If the parsing process itself returned a non-nil result (parseTree) and didn't panic
	// due to unrecoverable errors, we can proceed. The DiagnosticErrorListener will
	// have already printed messages for any syntax errors found.
	// A more robust approach would involve a custom error listener that tracks errors.

	// For now, we'll proceed assuming the DiagnosticErrorListener has informed the user
	// if there were errors. If parsing failed critically, it might have panicked or
	// returned a nil parseTree (though the latter is less common with default error strategy).
	// A simple check for nil parseTree is not reliable for all error cases.
	// Relying on the listener printing errors and potentially exiting manually
	// or using a custom listener is the standard approach.

	// 3. Build AST
	builder := ast.NewASTBuilder() // Use your AST builder
	// Ensure the parseTree is not nil before attempting to build the AST,
	// although with default error strategy, it might not be nil even on errors.
	// The AST builder should ideally handle potential errors/incomplete nodes from the parse tree.
	astProgram := parseTree.Accept(builder).(*ast.Program)

	// Optional: Print a representation of the AST for debugging
	// fmt.Printf("AST: %+v\n", astProgram)

	// 4. Compile
	comp := compiler.New() // Create a new compiler instance
	bytecode, err := comp.Compile(astProgram)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Compilation error: %v\n", err)
		os.Exit(1)
	}

	// --- Add this section to inspect function bytecode ---
	fmt.Println("--- Main Program Bytecode ---")
	// Use the FormatInstructions method from the compiler package (defined in code.go)
	// Note: FormatInstructions is a method on the Bytecode struct itself now.
	fmt.Println(bytecode.FormatInstructions())
	fmt.Println("---------------------------")

	fmt.Println("--- Constants ---")
	// Print the main program's constants
	for i, constant := range bytecode.Constants {
		fmt.Printf("%d: %s\n", i, constant.Inspect())

		// Check if the constant is a Closure (a compiled function)
		if closure, ok := constant.(*types.Closure); ok {
			// If it's a closure, format and print its internal bytecode
			fmt.Printf("  --- Bytecode for Function (Constant %d) ---\n", i)
			// The instructions are in the CompiledFunction inside the Closure
			// We need to create a temporary Bytecode struct to call FormatInstructions.
			// Only the Instructions field is needed for formatting.
			funcBytecode := &compiler.Bytecode{
				Instructions: closure.Fn.Instructions,
				// Constants:     closure.Fn.Constants, // Removed: CompiledFunction does not have a Constants field
				NumLocals:     closure.Fn.NumLocals,
				NumParameters: closure.Fn.NumParameters,
				NumGlobals:    0, // Function bytecode has no globals
			}
			fmt.Println(funcBytecode.FormatInstructions())
			fmt.Println("  ------------------------------------------")
		}
	}
	fmt.Println("-----------------")
	// --- End of inspection section ---

	// 5. Execute
	// Create a new VM instance, passing only the bytecode.
	// The VM will get the number of globals from bytecode.NumGlobals.
	vm := vm.New(bytecode)

	// Run the bytecode. VM.Run now returns only an error.
	err = vm.Run() // Corrected assignment
	if err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		os.Exit(1)
	}

	// The VM.Run method in the latest vm_go_updated immersive returns an error, not a value.
	// If you want to capture the final result, you might need to modify VM.Run
	// to return a value.Value as well, or inspect the stack after Run finishes.
	// For now, we'll remove the result printing as VM.Run returns only error.
	// fmt.Printf("Program finished. Result: %+v\n", result)

	fmt.Println("Execution finished.")
}
