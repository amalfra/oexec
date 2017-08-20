package oexec

import (
	"strings"
	"sync"
	"testing"
	"time"
)

func TestExecuteWithOutArguments(t *testing.T) {
	cmd := "echo"
	tmpOutput, returnedError := execute(cmd)
	returnedOutput := strings.TrimSpace(string(tmpOutput))
	if returnedOutput != "" || returnedError != nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestExecuteWithArguments(t *testing.T) {
	outputStr := "cmd executed"
	cmd := "echo"
	tmpOutput, returnedError := execute(cmd, outputStr)
	returnedOutput := strings.TrimSpace(string(tmpOutput))
	if returnedOutput != outputStr || returnedError != nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestExecuteWithInvalidCmd(t *testing.T) {
	cmd := "dfgffd"
	returnedOutput, returnedError := execute(cmd)
	if returnedOutput != nil || returnedError == nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestProcessCmdStr(t *testing.T) {
	cmd := "ls -l"
	cmdName, args := processCmdStr(cmd)
	if cmdName != "ls" || len(args) > 1 || args[0] != "-l" {
		t.Fatalf("Incorrect output generated")
	}
}

func TestRun(t *testing.T) {
	cmd := "echo 'test'"
	out := run(cmd)
	if strings.TrimSpace(string(out.Stdout)) != "'test'" || out.Stderr != nil {
		t.Fatalf("Incorrect output generated")
	}
}
func TestRunInvalidCmd(t *testing.T) {
	cmd := "slkmfdlkds"
	out := run(cmd)
	if out.Stdout != nil || out.Stderr == nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestRunParallel(t *testing.T) {
	var wg sync.WaitGroup
	outputs := make([]*Output, 4)
	wg.Add(1)
	go runParallel(&wg, "echo test", 3, outputs)
	wg.Wait()
	if strings.TrimSpace(string(outputs[3].Stdout)) != "test" || outputs[3].Stderr != nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestSeries(t *testing.T) {
	startTime := time.Now()
	outs := Series("sleep 2s", "sleep 3s", "sdfds", "sleep 2s")
	elapsedTime := time.Since(startTime)
	if elapsedTime < 7 || outs[2].Stderr == nil {
		t.Fatalf("Incorrect output generated")
	}
}

func TestParallel(t *testing.T) {
	startTime := time.Now()
	outs := Parallel("sleep 2s", "sleep 3s", "sdfds", "sleep 2s")
	elapsedTime := time.Since(startTime)
	if elapsedTime.Seconds() > 3.1 || outs[2].Stderr == nil {
		t.Fatalf("Incorrect output generated")
	}
}
