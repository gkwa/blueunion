package blueunion_test

import (
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
	"github.com/taylormonacelli/blueunion"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"blueunion": blueunion.Execute,
	}))
}

func TestHello(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
