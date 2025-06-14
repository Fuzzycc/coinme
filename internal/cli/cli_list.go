package cli

import (
	iu "coinme/internal/utils"
	"io"
)

func PrintAll(r io.Reader) {
	out, err := iu.IORead(r)
	iu.CrashErr(err, "coinme-utils")
	iu.Print(out)
}
