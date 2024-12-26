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

	want := `level=info msg="Hello from logrus!"
=== All Dependencies ===
- github.com/Yang-33/line-bot-sdk-go-521-lib v0.1.0
- github.com/sirupsen/logrus v1.9.0
- golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8
========================
Status: 200 OK
`

	got := string(out)
	if got != want {
		t.Errorf("unexpected output.\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}
