package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hi" {
		t.Fatalf("expected \"%s\" got \"%s\"", "hi", string(out))
	}
}
