package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	testwait()
}

func testwait() {
	cmd := exec.Command("bash", "-c", "~/go/src/golang-test/exec/exec.sh")
	var stdout, stderr []byte
	var errStdout, errStderr error
	//stdoutIn, _ := cmd.StdoutPipe()
	//stderrIn, _ := cmd.StderrPipe()
	cmd.Start()
	//go func() {
	//	stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	//}()
	//go func() {
	//	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	//}()
	err := cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			os.Stdout.Write(d)
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
	// never reached
	panic(true)
	return nil, nil
}
