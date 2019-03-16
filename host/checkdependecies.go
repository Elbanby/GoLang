package main

import (
	"fmt"
	"os/exec"
)

func main() {
	err := checkDep("apt")
	if err != nil {
		fmt.Println(err)
	}
}

func checkDep(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("Couldn't find %s in PATH: %s", name, err)
	}
	return nil
}
