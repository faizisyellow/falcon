package tools

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallGoTool(pkg, bin string) error {
	if HasCommand(bin) {
		return nil
	}

	fmt.Printf("Installing %s...\n", bin)
	cmd := exec.Command("go", "install", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func HasCommand(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}
