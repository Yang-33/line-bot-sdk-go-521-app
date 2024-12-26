package main

import (
	"os/exec"
	"testing"
)

func TestPrintAllDependencies(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run main.go: %v\noutput:\n%s", err, string(out))
	}

	want := `=== All Dependencies ===
- github.com/Yang-33/line-bot-sdk-go-521-lib v0.1.0
========================
Status: 200 OK
Response Body: <!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <m`

	got := string(out)
	if got != want {
		t.Errorf("unexpected output.\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}