package cls

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//clears terminal's screen
func CLS() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear") //Linux example, its tested
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") //Windows example, its tested
	default:
		fmt.Println("CLS for ", runtime.GOOS, " not implemented")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
