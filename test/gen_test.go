package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "golang.org/x/tools/go/packages"
	"os"
	"os/exec"
	"testing"
)

func TestGen(t *testing.T) {
	cmd := exec.Command("go", "run", "../generate", "./ent/schema")
	dir, _ := os.Getwd()
	cmd.Dir = dir
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	require.NoError(t, cmd.Run(), stderr.String(), stdout.String())
}

func TestMapArray(t *testing.T) {
	x := make(map[string][]string, 0)
	assert.Len(t, x["?"], 0)
}
