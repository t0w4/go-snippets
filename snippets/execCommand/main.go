package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("At least one argument is required.")
		os.Exit(1)
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	// cmd.Run(): cmd.Start() + cmd.Wait()
	//err := cmd.Run()
	// cmd.Output(): Run() + Stdout
	// cmd.CombinedOutput(): Run() + Stdout + Stderr
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("exec command error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("--------Output Start--------")
	fmt.Print(string(output))
	fmt.Println("--------Output END  --------")

	state := cmd.ProcessState
	fmt.Printf("%s\n", state.String())
	fmt.Printf(" Pid: %d\n", state.Pid())
	fmt.Printf(" Exited: %v\n", state.Exited())
	fmt.Printf(" Success: %v\n", state.Success())
	fmt.Printf(" System: %v\n", state.SystemTime())
	fmt.Printf(" User: %v\n", state.UserTime())
}
