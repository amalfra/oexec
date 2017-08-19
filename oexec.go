package oexec

import (
	"os/exec"
	"strings"
)

// Series will execute specified cmds in series and return slice of Outputs
func Series(cmds ...string) []*Output {
	outputs := make([]*Output, len(cmds))
	for index, cmd := range cmds {
		cmdName, cmdArgs := processCmdStr(cmd)
		outStruct := new(Output)
		outStruct.Stdout, outStruct.Stderr = execute(cmdName, cmdArgs...)
		outputs[index] = outStruct
	}
	return outputs
}

// processCmdStr will split full command string to command and arguments slice
func processCmdStr(cmd string) (cmdName string, cmdArgs []string) {
	cmdParts := strings.Split(cmd, " ")
	return cmdParts[0], cmdParts[1:]
}

// execute will run the specified command with arguments
// and return output, error
func execute(cmd string, args ...string) ([]byte, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
