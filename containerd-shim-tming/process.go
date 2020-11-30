package main

import (
	"fmt"
	"os/exec"
	"os"
)

type process struct {
	id 		string
	bundle 		string
	runtime 	string
}

func newProcess(id, bundle, runtimeName string) (*process, error) {
	p := &process{
		id: 		id,
		bundle: 	bundle,
		runtime: 	runtimeName,
	}
	return p, nil
}


func (p *process) Close() error {
	fmt.Println("process close")
	return nil
}

func (p *process) start() error {
	args := make([]string, 0)
	args = append(args, "start", "--bundle", p.bundle)
	args = append(args, "-d")

	cmd := exec.Command(p.runtime, args...)
	cmd.Dir = p.bundle
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("error err: %v\n", err)
		return err
	}
	if err := cmd.Wait(); err != nil {
		fmt.Printf("error err: %v\n", err)
		return err
	}
	return nil
}