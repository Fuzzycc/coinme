package utils

import (
	"fmt"
	"os"
)

// --- --- --- ---
// Write
// --- --- --- ---

// Prints operands to Standard Out, ending with a new line. Operands are space-separated.
func Print(a ...any) {
	PrintOut(a...)
}

func PrintOut(a ...any) {
	fmt.Println(a...)
}

func LogErr(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}
