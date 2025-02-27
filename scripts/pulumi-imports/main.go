package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
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

		if err := os.WriteFile(f, newBytes, 0700); err != nil {
			log.Fatal(err)
		}
	}

}
