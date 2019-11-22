package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	//预先检查ls程序是否存在
	binary,lookErr :=exec.LookPath("ls")
	if lookErr !=nil{
		panic(lookErr)
	}

	args := []string{"ls","-a","-l","-h"}

	env :=os.Environ()

	execErr := syscall.Exec(binary,args,env)
	if execErr!=nil{
		panic(execErr)
	}
}