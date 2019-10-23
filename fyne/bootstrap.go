package main

import "fmt"
import "os/exec"
import "strings"

func goGetFyne() {
	exec.Command("go", "get", "-u", "fyne.io/fyne").Run()
}

func goEnvPath() string {
	path, _ := exec.Command("go", "env", "GOPATH").CombinedOutput()
	return strings.TrimSpace(string(path))
}

func fyneExamplesPath() string {
	return fmt.Sprintf("%s/src/github.com/fyne-io/fyne/cmd/fyne_demo", goEnvPath())
}

func main() {
	fmt.Println("[INFO ] Updating Fyne APIs")
	goGetFyne()

	fmt.Println("[INFO ] Running Fyne examples")
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = fyneExamplesPath()
	cmd.Run()
}
