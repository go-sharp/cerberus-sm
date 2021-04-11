// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

var versionRe = regexp.MustCompile(`v([0-9]+)[.]([0-9]+)[.]([0-9]+)`)

func main() {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	version := string(output)

	tags := versionRe.FindStringSubmatch(version)
	if len(tags) < 4 {
		log.Fatalln("Didn't find a valid tag (ex. v2.1.0):", version)
	}

	build(version)
	fmt.Println("Successfully build binaries for tag:", version)
}

func build(version string) {
	fmt.Println("Building cerberus-sm binaries...")
	c := createCommand(version)
	if err := c.Run(); err != nil {
		log.Fatalln(err)
	}
}

func createCommand(version string) *exec.Cmd {
	cmd := exec.Command("wails", "build", "-f", "-ldflags", "-X main.version="+version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
