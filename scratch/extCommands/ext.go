/*
Thanks nate.
https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		cmdOut []byte
		err    error
	)
	// -F 0.95 -L 0.03 -M 30 dealer.pdf
	cmdName := "python"
	cmdArgs := []string{"./pdf2text.py", "-F", "0.95", "-L", "0.03", "-M", "30", "dealer.pdf"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running command: ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	fmt.Println(sha)
}
