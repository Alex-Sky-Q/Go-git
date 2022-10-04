package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func killService(pidFile string) error {
	f, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	var pid int
	_, err = fmt.Fscan(rd, &pid)
	if err != nil {
		return fmt.Errorf("%s. There is no PID in the file", err)
	}
	fmt.Printf("Service with PID %d was killed\n", pid)
	return nil
}

func main() {
	err := killService("service_pid.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
