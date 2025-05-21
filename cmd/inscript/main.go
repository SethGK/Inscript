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

	// Enable tracing and diagnostics (optional, for clearer output during debugging)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// lexer.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Use the default error strategy
	p.SetErrorHandler(antlr.NewDefaultErrorStrategy())

	// Parse the top-level rule
	parseTree := p.Program()

	// 3. Build AST
	builder := ast.NewASTBuilder()
	astProgram := parseTree.Accept(builder).(*ast.Program)

	// 4. Compile
	comp := compiler.New()
	bytecode, err := comp.Compile(astProgram)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Compilation error: %v\n", err)
		os.Exit(1)
	}

	// --- Bytecode Inspection ---
	fmt.Println("--- Main Program Bytecode ---")
	// Use the String() method on compiler.Instructions
	fmt.Println(bytecode.Instructions.String())
	fmt.Println("---------------------------")

	fmt.Println("--- Constants ---")
	for i, constant := range bytecode.Constants {
		fmt.Printf("%d: %s\n", i, constant.Inspect())

		// If the constant is a CompiledFunction (raw function bytecode)
		if compiledFn, ok := constant.(*types.CompiledFunction); ok {
			fmt.Printf("  --- Bytecode for Compiled Function (Constant %d) ---\n", i)
			// Call String() method on the function's instructions
			fmt.Println(compiler.Instructions(compiledFn.Instructions).String())
			fmt.Println("  ------------------------------------------")
		} else if closure, ok := constant.(*types.Closure); ok {
			// If it's a Closure, inspect its underlying CompiledFunction's bytecode
			fmt.Printf("  --- Bytecode for Closure's Function (Constant %d) ---\n", i)
			fmt.Println(compiler.Instructions(closure.Fn.Instructions).String())
			fmt.Printf("  Free Variables Count: %d\n", len(closure.Free))
			for j, freeVar := range closure.Free {
				fmt.Printf("    Free[%d]: %s\n", j, freeVar.Inspect())
			}
			fmt.Println("  ------------------------------------------")
		}
	}
	fmt.Println("-----------------")
	// --- End of bytecode inspection ---

	// 5. Execute
	vm := vm.New(bytecode)

	err = vm.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Execution finished.")
}
