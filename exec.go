package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

func Exec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	var (
		stdOut bytes.Buffer
		stdErr bytes.Buffer
	)
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if err := cmd.Run(); err != nil {
		return err
	}
	if stdErr.Len() > 0 {
		return errors.New(stdErr.String())
	}
	if stdOut.Len() > 0 {
		log.Print(stdOut.String())
	}
	return nil
}
