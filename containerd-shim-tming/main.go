package main

import (
	"flag"
	"os"
	"os/signal"
	"fmt"
)

func writeMessage(f *os.File, level string, err error) {
	fmt.Fprintf(f, `{"level": "%s","msg": "%s"}`, level, err)
}

func main() {
	flag.Parse()
	//cwd, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	if err := start(); err != nil {
		fmt.Printf("start() error : %v\n", err)
		os.Exit(1)
	}
}

func start() error {
	signals := make(chan os.Signal, 2048)
	signal.Notify(signals)

	p, err := newProcess(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	if err != nil {
		return err
	}
	defer func() {
		if err := p.Close(); err != nil {
			fmt.Printf("close error %v\n", err)
			return err
		}
	}()

	if err := p.start(); err != nil {
		fmt.Printf("p.start() error : %v\n", err)
		return err
	}
	return nil
}
