package run

import (
	"fmt"
	"os/exec"
)

func RunMacOS(path string, args ...string) error {
	cmd := exec.Command(`open`, append([]string{`-a`, path, `--args`}, args...)...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		if xerr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("%s\n", xerr.Stderr)
		}
		return err
	}
	return nil
}
