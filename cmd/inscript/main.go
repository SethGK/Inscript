package main

import (
	"fmt"
	"os"

	"github.com/SethGK/Inscript.git/internal/vm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: inscript <script.ins>")
		return
	}

	scriptPath := os.Args[1]
	source, err := os.ReadFile(scriptPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// TODO: Tokenize, parse, compile, and run source code.
	fmt.Println("Executing:", scriptPath)
	_ = vm.Execute(string(source))
}
