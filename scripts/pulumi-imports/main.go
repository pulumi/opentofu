package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var buf bytes.Buffer
	cmd := exec.Command("git", "ls-files", "**/*.go")
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range strings.Split(buf.String(), "\n") {
		f = strings.TrimSpace(f)
		if f == "" {
			continue
		}

		if strings.Contains(f, "pulumi") {
			continue
		}

		origBytes, err := os.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}

		search := []byte("github.com/opentofu/opentofu")
		repl := []byte("github.com/pulumi/opentofu")

		newBytes := bytes.ReplaceAll(origBytes, search, repl)

		search = []byte("internal/")
		repl = []byte("")

		newBytes = bytes.ReplaceAll(newBytes, search, repl)

		if err := os.WriteFile(f, newBytes, 0700); err != nil {
			log.Fatal(err)
		}
	}

	internals, err := os.ReadDir("internal")
	if err != nil {
		log.Fatal(err)
	}

	for _, subdir := range internals {
		old := filepath.Join("internal", subdir.Name())
		new := subdir.Name()
		if err := os.Rename(old, new); err != nil {
			log.Fatal(err)
		}
	}

}
