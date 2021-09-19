package oexec

import (
	"os/exec"
	"strings"
	"sync"
)

// Series will execute specified cmds in series and return slice of Outputs
func Series(cmds ...string) []*Output {
	outputs := make([]*Output, len(cmds))
	for index, cmd := range cmds {
		outputs[index] = run(cmd)
	}
	return outputs
}

// Parallel will execute specified cmds in parallel and return slice of
// Outputs when all of them have completed
func Parallel(cmds ...string) []*Output {
	var wg sync.WaitGroup
	outputs := make([]*Output, len(cmds))
	for index, cmd := range cmds {
		wg.Add(1)
		go runParallel(&wg, cmd, index, outputs)
	}
	wg.Wait()
	return outputs
}

// runParallel is a helper go routine to run cmd in async
// and put result into array
func runParallel(wg *sync.WaitGroup, cmd string, index int, outputs []*Output) {
	defer wg.Done()
	outputs[index] = run(cmd)
}

// run will process the specified cmd, execute it and return the Output struct
func run(cmd string) *Output {
	cmdName, cmdArgs := processCmdStr(cmd)
	outStruct := new(Output)
	outStruct.Stdout, outStruct.Stderr = execute(cmdName, cmdArgs...)
	return outStruct
}

// removeEmptyStr will remove empty strings from the slice
func removeEmptyStr(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" && str != " " {
			r = append(r, str)
		}
	}
	return r
}

// processCmdStr will split full command string to command and arguments slice
func processCmdStr(cmd string) (cmdName string, cmdArgs []string) {
	cmdParts := removeEmptyStr(strings.Split(cmd, " "))
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
