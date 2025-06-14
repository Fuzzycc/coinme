package utils

import (
	"fmt"
	"log"
	"os"
)

// ------------
// Metadata
// ------------
// VERSION
// 0.1.0
// ------------

// ------------
// ErroF section
// ------------


// Wraps fmt.Errof:
//
//- > Errorf formats according to a format specifier and returns the string as a value that satisfies error.
//
//- > If the format specifier includes a %w verb with an error operand, the returned error will implement an Unwrap method returning the operand. If there is more than one %w verb, the returned error will implement an Unwrap method returning a []error containing all the %w operands in the order they appear in the arguments. It is invalid to supply the %w verb with an operand that does not implement the error interface. The %w verb is otherwise a synonym for %v.
func ErroF(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

// ------------
// Dash section
// ------------


// Dash is equivalent to [DashOut(err, "")]
//
// Useful if you want to ignore certain errors, but want to log them to standard out with no added prefix.
// Use [DashBool] instead if you want no error message logging
func Dash(err error) (erred bool) {
	return DashOut(err, "")
}

// Dash ignores an error, but returns whether it was non-nil.
//
// Useful if you want to ignore certain errors, but want to log if any were encountered.
func DashBool(err error) (erred bool) {
	return err != nil
}

// DashBlind ignores any error with no return whatsoever.
//
// Run, run and never look back! Aaaaah!
func DashNull(err error) {
	// What did you expect?
}

// DashStr ignores an error, but returns its string format appended to prefix. Empty if err is nil.
//
// Useful if you want to ignore certain errors, but want to log them.
func DashStr(err error, prefix string) (errString string) {
	if prefix != "" {
		prefix = prefix + ": "
	}
	if err != nil {
		return fmt.Sprintf("%s%s", prefix, err)
	}
	return ""
}

// DashOut ignores an error, but returns whether it was non-nil and prints it to standard out appended to prefix.
// Only when erred is true does it print.
//
// Useful if you want to ignore certain errors, but want to log them to Stdout.
func DashOut(err error, prefix string) (erred bool) {
	if prefix != "" {
		prefix = prefix + ": "
	}
	if err != nil {
		fmt.Printf("%s%s\n", prefix, err)
		return true
	}
	return false
}

// DashStderr ignores an error, but returns whether it was non-nil and prints it to standard error appended to prefix.
// Only when erred is true does it print.
//
// Useful if you want to ignore certain errors, but want to log them to Stderr.
func DashErr(err error, prefix string) (erred bool) {
	if prefix != "" {
		prefix = prefix + ": "
	}
	if err != nil {
		// _, errWrite :=
		fmt.Fprintf(os.Stderr, "%s%s\n", prefix, err)

		// tolerance := 100 // Doubt FprintF will fail to write -tolerance- times, but...
		// for errWrite != nil {
		// 	if tolerance <= 0 {
		// 		break
		// 	}
		// 	_, errWrite = fmt.Fprintf(os.Stderr, "Dashorcrash.DashStderr(): %s\n", err)
		// 	tolerance--
		// } // ... is this necessary?
		// NUKED

		return true
	}
	return false
}

// ------------
// Crash Section
// ------------

// Crash is equivalent to [CrashOut(err, "")]
//
// Useful if you want to terminate on certain errors while logging them to standard out with no added prefixes.
func Crash(err error) {
	CrashOut(err, "")
}

// CrashOut prints err to standard output if non-nil followed by os.Exit(1), else nothing happens.
//
// Useful if you want to terminate on certain errors while logging them to standard out.
func CrashOut(err error, prefix string) {
	if prefix != "" {
		prefix = prefix + ": "
	}
	if err != nil {
		log.Fatalf("%s%s\n", prefix, err)
	}
}

// CrashErr prints err to standard error if non-nil followed os.Exit(1), else nothing happens.
//
// Useful if you want to terminate on certain errors while logging them to standard error.
func CrashErr(err error, prefix string) {
	if prefix != "" {
		prefix = prefix + ": "
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s%s\n", prefix, err)
		os.Exit(1)
	}
}
