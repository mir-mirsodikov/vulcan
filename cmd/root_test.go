package cmd

import (
	"bytes"
	"io"
	"testing"
)

func Test_RootCommand(t *testing.T) {
	testCmd := NewRootCommand()
	b := bytes.NewBufferString("")

	testCmd.SetOut(b)
	testCmd.Execute()

	got, _ := io.ReadAll(b)
	want := "Hello, World!\n"

	if string(got) != want {
		t.Errorf("got %q want %q", string(got), want)
	}
}
